include .env
export

auth:
	go run services/auth/cmd/main.go

users:
	go run services/users/cmd/main.go

proto:
	protoc --proto_path=proto --go_out=proto/_generated --go_opt=paths=source_relative \
	--go-grpc_out=proto/_generated --go-grpc_opt=paths=source_relative \
	proto/users/users.proto

.PHONY: users auth proto