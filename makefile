include .env
export

profiles:
	go run services/users/cmd/main.go

.PHONY: profiles