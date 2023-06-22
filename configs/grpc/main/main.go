package mainconfig

import "os"

type MainConfig struct {
	IpAddr string
}

func GetMainConfig() *MainConfig {
	conf := &MainConfig{
		IpAddr: "192.168.207.152",
	}
	ip := os.Getenv("SERVICE_IP")
	if ip != "" {
		conf.IpAddr = ip
	}
	return conf
}

