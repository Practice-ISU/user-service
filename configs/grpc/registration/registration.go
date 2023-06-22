package registration

import "os"

type Config struct {
	ServiceName   string
	DiscoveryAddr string
}

func GetConfig() *Config {
	conf := &Config{
		ServiceName:   "UserGrpcService",
		DiscoveryAddr: "158.160.26.1:80",
	}
	serviceName := os.Getenv("SERVICE_NAME")
	discoveryAddr := os.Getenv("DISCOVERY_ADDR")

	if serviceName != "" {
		conf.ServiceName = serviceName
	}
	if discoveryAddr != "" {
		conf.DiscoveryAddr = discoveryAddr
	}

	return conf


}
