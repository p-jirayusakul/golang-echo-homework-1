include .env
export

auth:
	go run services/auth/cmd/main.go

users:
	go run services/users/cmd/main.go

.PHONY: users auth