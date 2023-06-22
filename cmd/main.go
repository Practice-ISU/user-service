package main

import (
	"fmt"
	"log"
	"net"
	"time"
	mainconfig "user-service/configs/grpc/main"
	ping_conf "user-service/configs/grpc/ping"
	reg_conf "user-service/configs/grpc/registration"
	user_conf "user-service/configs/grpc/user"
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
	mainConfig := mainconfig.GetMainConfig()
	userConfig := user_conf.GetConfig()
	pingConf := ping_conf.GetConfig()
	regCong := reg_conf.GetConfig()

	userStorage, err := psql_stor.NewUserStorage(psql_conf.GetDefaultConfig())
	if err != nil {
		panic(err)
	}
	userSevice := service.NewUserService(userStorage)
	userServer := server.NewUserServer(userSevice)

	s := grpc.NewServer()
	user_grpc.RegisterUserGrpcServiceServer(s, userServer)

	userListener, err := net.Listen("tcp", mainConfig.IpAddr+":"+userConfig.Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IP " + userConfig.ServiceName + " = " + userListener.Addr().String())

	pingServer, err := ping.NewDiscoveryPingServer(
		pingConf,
		regCong,
		mainConfig,
		userConfig,
	)
	if err != nil {
		log.Fatal(err)
	}

	ping_grpc.RegisterDiscoveryPingServer(s, pingServer)

	pingListener, err := net.Listen("tcp", mainConfig.IpAddr+":"+pingConf.Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("IP PingGrpcServer = " + pingListener.Addr().String())

	go func() {
		err = s.Serve(pingListener)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		pingServer.SendRegistrationRequest()
	}()

	go func() {
		for {
			pingServer.StartTimeout(pingServer.SendRegistrationRequest)
			time.Sleep(12 * time.Second)
		}
	}()

	err = s.Serve(userListener)
	if err != nil {
		log.Fatal(err)
	}
}
