package handlers

import (
	"github.com/isaaczzzz/grpc-module/gen/api/echo"
	"github.com/isaaczzzz/grpc-module/internal/services"
	"google.golang.org/grpc"
)

func InitHandlers(s *grpc.Server, services *services.Services) {
	// 注册 Echo 服务
	echoHandler := NewEchoHandler(services.EchoService)
	echo.RegisterEchoServiceServer(s, echoHandler)
}
