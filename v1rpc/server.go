package v1rpc

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

const (
	maxRequestBytes   = 1.5 * 1024 * 1024
	grpcOverheadBytes = 512 * 1024
	maxStreams        = 1<<32 - 1 // math.MaxUint32와 동일
	maxSendBytes      = 1<<31 - 1 // math.MaxInt32와 동일
)

var (
	address = ":50052"
)

func init() {
	// TODO: Prometheus 적용 예정
}

func Server() error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(int(maxRequestBytes + grpcOverheadBytes)),
		grpc.MaxSendMsgSize(maxSendBytes),
		grpc.MaxConcurrentStreams(maxStreams),
	}
	grpcServer := grpc.NewServer(opts...)
	// 서비스 등록
	RegisterJobsManSrv(grpcServer)
	RegisterHelloWorldManSrv(grpcServer)
	// TODO 향후 수정한다. Reflection 서비스 등록
	reflection.Register(grpcServer)

	log.Printf("gRPC server started, address: %s", address)

	err = grpcServer.Serve(lis)
	if err != nil {
		if !strings.Contains(err.Error(), "use of closed network connection") {
			log.Printf("gRPC server returned with error: %v", err)
		} else {
			log.Printf("gRPC server is shut down")
		}
	}
	return err
}
