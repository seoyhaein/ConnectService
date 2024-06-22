package v1rpc

import (
	"context"
	pb "github.com/seoyhaein/ConnectService/protos"
	"google.golang.org/grpc"
)

// FileTransferManSrv 인터페이스 구현
type FileTransferManSrv struct {
	pb.UnimplementedFileTransferServiceServer
}

// factory 메서드로 구현
func newFileTransferManSrv() pb.FileTransferServiceServer {
	return &FileTransferManSrv{}
}

// RegisterFileTransferManSrv 서비스 등록
func RegisterFileTransferManSrv(service *grpc.Server) {
	pb.RegisterFileTransferServiceServer(service, newFileTransferManSrv())
}

// UploadFile 형태만 잡음
func (f *FileTransferManSrv) UploadFile(ctx context.Context, fq *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	return nil, nil
}
