package handler

import (
	"context"
	"github.com/isaaczzzz/grpc-module/gen/api/echo"
	"github.com/isaaczzzz/grpc-module/internal/service"
)

// EchoHandler 实现 IEchoService 接口
type EchoHandler struct {
	echoService service.IEchoService
}

// newEchoHandler 创建一个新的 EchoHandler 实例
func newEchoHandler(echoService service.IEchoService) *EchoHandler {
	return &EchoHandler{
		echoService: echoService,
	}
}

// Echo 实现 IEchoService 的 Echo 方法
func (h *EchoHandler) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return h.echoService.Echo(ctx, req)
}
