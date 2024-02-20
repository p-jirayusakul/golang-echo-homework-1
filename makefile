proto:
	protoc --proto_path=proto --go_out=proto/_generated --go_opt=paths=source_relative \
	--go-grpc_out=proto/_generated --go-grpc_opt=paths=source_relative \
	proto/users/users.proto

.PHONY: proto