package handlers

import (
	"context"
	"github.com/isaaczzzz/grpc-module/gen/api/echo"
	"github.com/isaaczzzz/grpc-module/internal/services"
)

// EchoHandler 实现 EchoService 接口
type EchoHandler struct {
	echoService services.EchoService
}

// NewEchoHandler 创建一个新的 EchoHandler 实例
func NewEchoHandler(echoService services.EchoService) *EchoHandler {
	return &EchoHandler{
		echoService: echoService,
	}
}

// Echo 实现 EchoService 的 Echo 方法
func (h *EchoHandler) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return h.echoService.Echo(ctx, req)
}
