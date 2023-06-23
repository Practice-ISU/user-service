package postgre

import "os"

type Config struct {
	Host     string
	DBname   string
	Port     string
	Username string
	Password string
	SSL      string
}

func GetConfig() *Config {
	cnf := &Config{
		Host:     "localhost",
		Port:     "5432",
		DBname:   "users_db",
		Username: "postgres",
		Password: "qwer",
		SSL:      "disable",
	}
	host := os.Getenv("POSTGRES_HOST")
	if host != "" {
		cnf.Host = host
	}
	dbName := os.Getenv("POSTGRES_DBNAME")
	if dbName != "" {
		cnf.DBname = dbName
	}
	port := os.Getenv("POSTGRES_PORT")
	if port != "" {
		cnf.Port = port
	}
	user := os.Getenv("POSTGRES_USER")
	if user != "" {
		cnf.Username = user
	}
	pass := os.Getenv("POSTGRES_PASSWORD")
	if pass != "" {
		cnf.Password = pass
	}
	sslMode := os.Getenv("POSTGRES_SSL_MODE")
	if sslMode != "" {
		cnf.SSL = sslMode
	}

	return cnf
}

func (cnf *Config) GetUserName() string {
	return cnf.Username
}

func (cnf *Config) GetPassword() string {
	return cnf.Password
}

func (cnf *Config) GetHost() string {
	return cnf.Host
}

func (cnf *Config) GetPort() string {
	return cnf.Port
}

func (cnf *Config) GetDbName() string {
	return cnf.DBname
}

func (cnf *Config) GetSSL() string {
	return cnf.SSL
}
