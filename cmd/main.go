package main

import (
	"log"
	"net"
	psql_conf "user-service/configs/postgre"
	server "user-service/internal/adapters/api/grpc"
	psql_stor "user-service/internal/adapters/db/postgre"
	"user-service/internal/domain/service"
	pkg_grpc "user-service/pkg/grpc"

	"google.golang.org/grpc"
)

func main() {
	userStorage, err := psql_stor.NewUserStorage(psql_conf.GetDefaultConfig())
	if err != nil {
		panic(err)
	}
	userSevice := service.NewUserService(userStorage)
	userServer := server.NewUserServer(userSevice)

	s := grpc.NewServer()
	pkg_grpc.RegisterUserGrpcServiceServer(s, userServer)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}

}
