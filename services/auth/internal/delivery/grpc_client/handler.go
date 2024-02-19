package grpc_client

import (
	"log"

	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/config"
	"google.golang.org/grpc"
)

func NewGrpcClient(cfg *config.AuthConfig) (*ServerClient, []error) {
	var (
		client ServerClient
		err    error
		errs   []error
	)

	client.UsersClient, err = grpc.Dial(cfg.RPC_USERS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect grpc auth: %v", err)

		errs = append(errs, err)
	}
	log.Println("rpc users started on", cfg.RPC_USERS)

	return &client, errs
}
