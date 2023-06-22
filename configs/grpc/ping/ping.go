package ping

import "os"

type Config struct {
	Port string
}

func GetConfig() *Config {
	conf := &Config{
		Port: "8010",
	}
	port := os.Getenv("PORT_PING")
	if port != "" {
		conf.Port = port
	}
	return conf
}
