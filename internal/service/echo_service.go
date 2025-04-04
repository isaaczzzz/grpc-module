package service

import (
	"context"
	"github.com/isaaczzzz/grpc-module/gen/api/echo"
)

// IEchoService 定义 Echo 服务接口
type IEchoService interface {
	Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error)
}

// EchoService 实现 Echo 服务接口
type EchoService struct{}

// newEchoService 创建一个新的 Echo 服务实例
func newEchoService() IEchoService {
	return &EchoService{}
}

// Echo 实现 Echo 服务的 Echo 方法
func (s *EchoService) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: req.Message}, nil
}
