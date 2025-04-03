package services

import (
	"context"
	"github.com/isaaczzzz/grpc-module/gen/api/echo"
)

// EchoService 定义 Echo 服务接口
type EchoService interface {
	Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error)
}

// EchoServiceImpl 实现 Echo 服务接口
type EchoServiceImpl struct{}

// NewEchoService 创建一个新的 Echo 服务实例
func NewEchoService() EchoService {
	return &EchoServiceImpl{}
}

// Echo 实现 Echo 服务的 Echo 方法
func (s *EchoServiceImpl) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: req.Message}, nil
}
