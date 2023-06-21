package main

import (
	"log"
	"net"
	"time"
	ping_conf "user-service/configs/grpc/ping"
	reg_conf "user-service/configs/grpc/registration"
	psql_conf "user-service/configs/postgre"
	server "user-service/internal/adapters/api/grpc"
	psql_stor "user-service/internal/adapters/db/postgre"
	"user-service/internal/domain/service"
	ping_grpc "user-service/pkg/grpc/discovery/ping"
	user_grpc "user-service/pkg/grpc/user"

	"user-service/pkg/api/grpc/ping"

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
	user_grpc.RegisterUserGrpcServiceServer(s, userServer)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	pingConf := ping_conf.GetDefaultConfig()
	sp, err := ping.NewDiscoveryPingServer(
		pingConf,
		reg_conf.GetDefaultConfig(),
	)
	if err != nil {
		log.Fatal(err)
	}

	discoveryServer := grpc.NewServer()
	ping_grpc.RegisterDiscoveryPingServer(discoveryServer, sp)

	l2, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err = discoveryServer.Serve(l2)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for {
			sp.StartTimeout(sp.SendRegistrationRequest)
			time.Sleep(1 * time.Second)
		}
	}()

	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}

}
