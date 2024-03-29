## Reference
Github ต้นฉบับ
https://github.com/febrihidayan/go-architecture-monorepo.git

Document Postman สำหรับทดสอบ
https://documenter.getpostman.com/view/6440353/2sA2r9WNkV

## Features
 - Register
 - Login
 - Auth
 - CRUD with GORM user_profiles, user_address

## Tech stack
 - docker compose
 - mono repo
 - clean architecture
 - golang echo
 - go viper
 - go validator
 - gorm
 - HTTP,GRPC
 

## localhost run
**run service users**:
``` bash
go run services/users/cmd/main.go
```

**run service auth**:
``` bash
go run services/auth/cmd/main.go
```

**database env**:
```bash
DATABASE_USER=postgres
DATABASE_HOST=localhost
DATABASE_PASSWORD=1234
DATABASE_PORT=5432
DATABASE_NAME=homework1
```

**users env**:
```bash
HTTP_PORT=":3002"
RPC_PORT=":9082"
JWT_SECRET=UucwjDH7AY40XLDyWpBUalCB151WgAfF
```

**auth env**:
```bash
HTTP_PORT=":3001"
RPC_USERS=":9082"
JWT_SECRET=UucwjDH7AY40XLDyWpBUalCB151WgAfF
SECRET_KEY=L1K0zInpkIYzVXqUQdvnOc7FtbKOvpsJ
```

**docker**:
```bash
docker-compose build
```
```bash
docker-compose up
```