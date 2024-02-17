include .env
export

profiles:
	go run services/users/cmd/main.go

auth:
	go run services/auth/cmd/main.go

.PHONY: profiles auth