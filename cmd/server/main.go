package main

import (
	"fmt"
	"github.com/isaaczzzz/grpc-module/internal/handler"
	"github.com/isaaczzzz/grpc-module/internal/service"
	"github.com/isaaczzzz/grpc-module/pkg/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 创建 gRPC 服务器
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GRPC.Address, cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// 创建服务实例
	services := service.InitServices()

	// 注册 handler 服务
	handler.InitHandlers(s, services)

	log.Printf("Server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
