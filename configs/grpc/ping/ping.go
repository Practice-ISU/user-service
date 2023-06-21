package ping

type Config struct {
	Port string
}

func GetDefaultConfig() *Config {
	return &Config{
		Port: "8010",
	}
}
