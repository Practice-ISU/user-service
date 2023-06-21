package registration

type Config struct {
	Port          string
	ServiceName   string
	DiscoveryAddr string
}

func GetDefaultConfig() *Config {
	return &Config{
		Port:          "8020",
		ServiceName:   "user-service",
		DiscoveryAddr: "localhost:5000",
	}
}
