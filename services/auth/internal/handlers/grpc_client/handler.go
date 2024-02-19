package grpc_client

import (
	"log"

	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	"google.golang.org/grpc"
)

func NewGrpcClient() (*ServerClient, []error) {
	var (
		client ServerClient
		err    error
		errs   []error
	)

	client.UsersClient, err = grpc.Dial(configs.Config.RPC_USERS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect grpc auth: %v", err)

		errs = append(errs, err)
	}
	log.Println("rpc users started on", configs.Config.RPC_USERS)

	return &client, errs
}
