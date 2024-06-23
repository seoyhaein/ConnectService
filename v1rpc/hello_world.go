package v1rpc

import (
	"context"
	pb "github.com/seoyhaein/ConnectService/protos"
	"google.golang.org/grpc"
)

// HelloWorldManSrv 인터페이스 구현
type HelloWorldManSrv struct {
	pb.UnimplementedGreeterServer
}

// newHelloWorldManSrv factory 메서드로 구현
func newHelloWorldManSrv() pb.GreeterServer {
	return &HelloWorldManSrv{}
}

// RegisterHelloWorldManSrv 서비스 등록
func RegisterHelloWorldManSrv(service *grpc.Server) {
	pb.RegisterGreeterServer(service, newHelloWorldManSrv())
}

func (h *HelloWorldManSrv) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
