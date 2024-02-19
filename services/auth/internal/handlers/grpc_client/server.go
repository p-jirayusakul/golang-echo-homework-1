package grpc_client

import "google.golang.org/grpc"

type ServerClient struct {
	UsersClient *grpc.ClientConn
}
