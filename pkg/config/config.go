package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// Config 定义配置结构体
type Config struct {
	GRPC struct {
		Address string
		Port    string
	}
	DB struct {
		Driver   string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Redis struct {
		Address  string
		Password string
		DB       int
	}
}

// LoadConfig 从 .env 文件加载配置
func LoadConfig() (*Config, error) {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	cfg := &Config{}

	// 加载 gRPC 配置
	cfg.GRPC.Address = os.Getenv("GRPC_ADDRESS")
	if cfg.GRPC.Address == "" {
		cfg.GRPC.Address = "localhost"
	}
	cfg.GRPC.Port = os.Getenv("GRPC_PORT")
	if cfg.GRPC.Port == "" {
		cfg.GRPC.Port = "50051"
	}

	// 加载数据库配置
	cfg.DB.Driver = os.Getenv("DB_DRIVER")
	if cfg.DB.Driver == "" {
		cfg.DB.Driver = "mysql"
	}
	cfg.DB.Host = os.Getenv("DB_HOST")
	if cfg.DB.Host == "" {
		cfg.DB.Host = "localhost"
	}
	cfg.DB.Port = os.Getenv("DB_PORT")
	if cfg.DB.Port == "" {
		cfg.DB.Port = "3306"
	}
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Name = os.Getenv("DB_NAME")

	// 加载 Redis 配置
	cfg.Redis.Address = os.Getenv("REDIS_ADDRESS")
	if cfg.Redis.Address == "" {
		cfg.Redis.Address = "localhost:6379"
	}
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")
	redisDBStr := os.Getenv("REDIS_DB")
	if redisDBStr == "" {
		cfg.Redis.DB = 0
	} else {
		redisDB, e := strconv.Atoi(redisDBStr)
		if e != nil {
			return nil, fmt.Errorf("failed to convert REDIS_DB to int: %w", err)
		}
		cfg.Redis.DB = redisDB
	}

	return cfg, nil
}
