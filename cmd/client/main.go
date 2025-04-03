package main

import (
	"context"
	"fmt"
	"github.com/isaaczzzz/grpc-module/gen/api/echo"
	"github.com/isaaczzzz/grpc-module/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 连接到 gRPC 服务器
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.GRPC.Address, cfg.GRPC.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connect: %v", err)
		}
	}(conn)

	// 创建 Echo 服务客户端
	client := echo.NewEchoServiceClient(conn)

	// 创建请求
	req := &echo.EchoRequest{
		Message: "Hello, Echo!",
	}

	// 发送请求
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Echo(ctx, req)
	if err != nil {
		log.Fatalf("Could not echo: %v", err)
	}

	log.Printf("Echo response: %s", resp.Message)
}
