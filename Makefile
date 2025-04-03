protoc-gen:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. api/proto/echo.proto

run-server:
	go run .\cmd\server\main.go