package ping

import (
	"context"
	"fmt"
	"time"

	// "net"
	main_conf "user-service/configs/grpc/main"
	ping_conf "user-service/configs/grpc/ping"
	reg_conf "user-service/configs/grpc/registration"

	"user-service/pkg/grpc/discovery/ping"
	"user-service/pkg/grpc/discovery/registration"

	"google.golang.org/grpc"
)

type ServiceConfig interface {
	GetPort() string
	GetServiceName() string
}

type DiscoveryPingServer struct {
	ping.UnimplementedDiscoveryPingServer
	serviceName  string
	pingPort     string
	servicePort  string
	ip           string
	register     registration.ServiceRegistrationClient
	lastCallTime time.Time
}

func NewDiscoveryPingServer(cnfPing *ping_conf.Config, cnfReg *reg_conf.Config, cnfMain *main_conf.MainConfig, cnfService ServiceConfig) (*DiscoveryPingServer, error) {
	fmt.Println("Discovery chanel = " + cnfReg.DiscoveryAddr)
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(cnfReg.DiscoveryAddr, options...)
	if err != nil {
		return nil, err
	}
	register := registration.NewServiceRegistrationClient(conn)
	serviceName := cnfService.GetServiceName()
	servicePort := cnfService.GetPort()

	return &DiscoveryPingServer{
		register:    register,
		serviceName: serviceName,
		pingPort:    cnfPing.Port,
		servicePort: servicePort,
		ip:          cnfMain.IpAddr,
	}, nil
}

func (s *DiscoveryPingServer) Ping(context.Context, *ping.PingRequest) (*ping.PingResponse, error) {
	s.lastCallTime = time.Now()
	// fmt.Println("Pinged in", s.lastCallTime.String())
	return &ping.PingResponse{
		Timestamp: time.Now().Format("2006-01-02 15:04:05.000"),
		Success:   true,
	}, nil
}

func (s *DiscoveryPingServer) SendRegistrationRequest() {
	for {
		data := &registration.ServiceRequest{
			Timestamp:   time.Now().Format("2006-01-02 15:04:05.000"),
			ServiceName: s.serviceName,
			Channel:     "http://" + s.ip + ":" + s.servicePort,
			ChannelPing: "http://" + s.ip + ":" + s.pingPort,
		}
		result, err := s.register.Registration(context.TODO(), data)

		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if result.Success {
			fmt.Println("registered in", result.Timestamp, "success -", result.Success)
			break
		}
		time.Sleep(time.Minute)
	}
}

func (s *DiscoveryPingServer) StartTimeout(f func()) {
	time.AfterFunc(6 * time.Minute, func() {
		if time.Since(s.lastCallTime) > 6 * time.Minute {
			fmt.Println("Lost ping from discovery-service!")
			f()
		}
	})

}
