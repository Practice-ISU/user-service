package postgre

import (
	"database/sql"
	"fmt"
	psql_cof "user-service/configs/postgre"

	_ "github.com/lib/pq"
)

func GetDb(cnf *psql_cof.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.DBname, cnf.SSL)
	return sql.Open("postgres", psqlInfo)
}
