package postgre

type Config struct {
	Host     string
	DBname   string
	Port     uint
	Username string
	Password string
	SSL      string
}

func GetDefaultConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     5432,
		DBname:   "users_db",
		Username: "postgres",
		Password: "qwer",
		SSL:      "disable",
	}
}
