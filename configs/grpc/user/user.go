package usesr

import "os"

type Config struct {
	Port        string
	ServiceName string
}

func GetConfig() *Config {
	conf := &Config{
		Port:        "8000",
		ServiceName: "UserGrpcService",
	}

	port := os.Getenv("PORT_USERS")
	serviceName := os.Getenv("SERVICE_NAME")

	if port != "" {
		conf.Port = port
	}
	if serviceName != "" {
		conf.ServiceName = serviceName
	}

	return conf
}

func (cnf *Config) GetPort() string {
	return cnf.Port
}
func (cnf *Config) GetServiceName() string {
	return cnf.ServiceName
}
