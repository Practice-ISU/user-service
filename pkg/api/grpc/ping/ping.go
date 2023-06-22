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
	conn, err := grpc.Dial(cnfReg.DiscoveryAddr, grpc.WithInsecure())
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
	fmt.Println("Pinged in " + s.lastCallTime.Format("2006-01-02 HH15:04:05.000"))
	return &ping.PingResponse{
		Timestamp: time.Now().Format("2006-01-02 HH15:04:05.000"),
		Success:   true,
	}, nil
}

func (s *DiscoveryPingServer) SendRegistrationRequest() {
	for {
		result, err := s.register.Registration(context.TODO(), &registration.ServiceRequest{
			Timestamp:   time.Now().Format("2006-01-02 HH15:04:05.000"),
			ServiceName: s.serviceName,
			Channel:     s.ip + ":" + s.servicePort,
			ChannelPing: s.ip + ":" + s.pingPort,
		})

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("tried")

		if result.Success {
			fmt.Printf("registered in %s", result.Timestamp)
			fmt.Println(result.Success, result.Timestamp)
			break
		}
		time.Sleep(time.Minute)
	}
}

func (s *DiscoveryPingServer) StartTimeout(f func()) {

	time.AfterFunc(10 * time.Second, func() {
		if time.Since(s.lastCallTime) > 10 * time.Second {
			fmt.Println("Timer is over!!!")
		} else {
			fmt.Println("Time is not over!!!")
		}
	})

}
