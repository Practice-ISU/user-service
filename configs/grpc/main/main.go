package mainconfig

import "os"

type MainConfig struct {
	IpAddr string
}

func GetMainConfig() *MainConfig {
	conf := &MainConfig{
		IpAddr: "127.0.0.1",
	}
	ip := os.Getenv("SERVICE_IP")
	if ip != "" {
		conf.IpAddr = ip
	}
	return conf
}

