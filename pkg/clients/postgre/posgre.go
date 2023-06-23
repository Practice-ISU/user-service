package postgre

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config interface {
	GetUserName() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetDbName() string
	GetSSL() string
}

type PsqlClient struct {
	cnf Config
}

func NewPsqlCliennt(cnf Config) *PsqlClient {
	return &PsqlClient{
		cnf: cnf,
	}
}

func (cl *PsqlClient) GetDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cl.cnf.GetUserName(),
		cl.cnf.GetPassword(),
		cl.cnf.GetHost(),
		cl.cnf.GetPort(),
		cl.cnf.GetDbName(),
		cl.cnf.GetSSL(),
	)
	return sql.Open("postgres", psqlInfo)
}
