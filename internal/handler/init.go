package handler

import (
	"github.com/isaaczzzz/grpc-module/gen/api/echo"
	"github.com/isaaczzzz/grpc-module/internal/service"
	"google.golang.org/grpc"
)

func InitHandlers(s *grpc.Server, services *service.Services) {
	// 注册 Echo 服务
	echo.RegisterEchoServiceServer(s, newEchoHandler(services.EchoService))
}
