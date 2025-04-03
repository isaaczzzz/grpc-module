package main

import (
	"fmt"
	"github.com/isaaczzzz/grpc-module/internal/handlers"
	"github.com/isaaczzzz/grpc-module/internal/services"
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

	//// 初始化数据库连接
	//dbConn, err := db.InitDB(cfg.DB)
	//if err != nil {
	//	log.Fatalf("Failed to initialize database: %v", err)
	//}
	//defer dbConn.Close()
	//
	//// 初始化 Redis 缓存
	//redisClient := cache.InitRedis(cfg.Redis)
	//defer redisClient.Close()

	// 创建 gRPC 服务器
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GRPC.Address, cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// 创建服务实例
	serviceImpl := services.InitServices()

	// 注册 handlers 服务
	handlers.InitHandlers(s, serviceImpl)

	log.Printf("Server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
