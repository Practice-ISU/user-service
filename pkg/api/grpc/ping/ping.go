package ping

import (
	"context"
	"fmt"
	"time"
	ping_conf "user-service/configs/grpc/ping"
	reg_conf "user-service/configs/grpc/registration"
	"user-service/pkg/grpc/discovery/ping"
	"user-service/pkg/grpc/discovery/registration"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type DiscoveryPingServer struct {
	ping.UnimplementedDiscoveryPingServer
	port         string
	register     registration.ServiceRegistrationClient
	lastCallTime time.Time
}

func NewDiscoveryPingServer(cnfPing *ping_conf.Config, cnfReg *reg_conf.Config) (*DiscoveryPingServer, error) {
	conn, err := grpc.Dial(cnfReg.Port, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return nil, err
	}
	register := registration.NewServiceRegistrationClient(conn)
	return &DiscoveryPingServer{
		register: register,
		port:     cnfPing.Port,
	}, nil
}

func (s *DiscoveryPingServer) Ping(context.Context, *ping.PingRequest) (*ping.PingResponse, error) {
	s.lastCallTime = time.Now()
	return &ping.PingResponse{
		Timestamp: time.Now().Format("2006-01-02 HH15:04:05.000"),
		Success:   true,
	}, nil
}

func (s *DiscoveryPingServer) SendRegistrationRequest() {
	for {
		result, err := s.register.Registration(context.TODO(), &registration.ServiceRequest{
			Timestamp: time.Now().Format("2006-01-02 HH15:04:05.000"),
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		if result.Success {
			break
		}
		time.Sleep(time.Minute)
	}
}

func (s *DiscoveryPingServer) StartTimeout(f func()) {
	timer := time.AfterFunc(10*time.Second, func() {
		if time.Since(s.lastCallTime) >= 10*time.Second {
			f()
		}
	})
	defer timer.Stop()
}
