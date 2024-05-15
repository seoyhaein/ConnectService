package v1rpc

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"

	pb "github.com/seoyhaein/ConnectService/protos"
	"google.golang.org/grpc"
)

// JobManSrv는 Job 관리 서비스를 구현합니다.
type JobManSrv struct {
	pb.UnimplementedLongLivedJobCallServer
	subscribers sync.Map
}

type sub struct {
	id       int64
	status   pb.JobsResponse_Status
	stream   pb.LongLivedJobCall_SubscribeServer
	finished chan<- bool
}

// RegisterJobsManSrv는 JobManSrv 서비스를 gRPC 서버에 등록합니다.
func RegisterJobsManSrv(service *grpc.Server) {
	pb.RegisterLongLivedJobCallServer(service, newJobsManSrv())
}

func newJobsManSrv() pb.LongLivedJobCallServer {
	return &JobManSrv{}
}

// Subscribe 메서드는 클라이언트의 Job 구독 요청을 처리합니다.
func (j *JobManSrv) Subscribe(in *pb.JobsRequest, s pb.LongLivedJobCall_SubscribeServer) error {
	fin := make(chan bool)
	j.subscribers.Store(in.JobReqId, sub{stream: s, finished: fin})
	ctx := s.Context()

	cmd, r, err := j.scriptRunner(ctx, in)
	if err != nil {
		log.Printf("Error creating script runner: %v", err)
		return err
	}

	go func(cmd *exec.Cmd) {
		if err := cmd.Start(); err != nil {
			log.Printf("Error starting Cmd: %v", err)
			return
		}
		if err := cmd.Wait(); err != nil {
			log.Printf("Error waiting for Cmd: %v", err)
		}
	}(cmd)

	go j.reply(r)

	for {
		select {
		case <-fin:
			log.Printf("Closing stream for client ID: %d", in.JobReqId)
			return nil
		case <-ctx.Done():
			log.Printf("Client ID %d has disconnected", in.JobReqId)
			return nil
		}
	}
}

// Unsubscribe 메서드는 클라이언트의 Job 구독 해제 요청을 처리합니다.
func (j *JobManSrv) Unsubscribe(ctx context.Context, req *pb.JobsRequest) (*pb.JobsResponse, error) {
	v, ok := j.subscribers.Load(req.JobReqId)
	if !ok {
		return nil, fmt.Errorf("failed to load subscriber key: %d", req.JobReqId)
	}
	sub, ok := v.(sub)
	if !ok {
		return nil, fmt.Errorf("failed to cast subscriber value: %T", v)
	}

	select {
	case sub.finished <- true:
		log.Printf("Unsubscribed client: %d", req.JobReqId)
	default:
	}

	j.subscribers.Delete(req.JobReqId)
	return &pb.JobsResponse{JobResId: req.JobReqId}, nil
}

// scriptRunner는 스크립트를 실행하고 stdout을 반환합니다.
func (j *JobManSrv) scriptRunner(ctx context.Context, in *pb.JobsRequest) (*exec.Cmd, io.Reader, error) {
	cmd := exec.CommandContext(ctx, "echo", in.InputMessage)
	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	return cmd, r, nil
}

// reply는 스크립트 실행 결과를 클라이언트에 전송합니다.
func (j *JobManSrv) reply(i io.Reader) {
	scan := bufio.NewScanner(i)
	unsubscribe := []int64{}

	for scan.Scan() {
		s := scan.Text()

		j.subscribers.Range(func(k, v interface{}) bool {
			id, ok := k.(int64)
			if !ok {
				log.Printf("Failed to cast subscriber key: %T", k)
				return false
			}
			sb, ok := v.(sub)
			if !ok {
				log.Printf("Failed to cast subscriber value: %T", v)
				return false
			}

			if err := sb.stream.Send(&pb.JobsResponse{JobResId: id, OutputMessage: s}); err != nil {
				log.Printf("Failed to send data to client: %v", err)
				select {
				case sb.finished <- true:
					log.Printf("Unsubscribed client: %d", id)
				default:
				}
				unsubscribe = append(unsubscribe, id)
			}
			return true
		})
	}

	if err := scan.Err(); err != nil {
		log.Printf("Error reading from stdout: %v", err)
	}

	for _, id := range unsubscribe {
		j.subscribers.Delete(id)
	}
}
