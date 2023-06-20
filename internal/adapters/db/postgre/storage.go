package postgre

import (
	"database/sql"
	"fmt"
	psql_conf "user-service/configs/postgre"
	"user-service/internal/domain/dto"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/storage"
	psql "user-service/pkg/clients/postgre"
)

type rowScanner interface {
	Scan(dest ...any) error
}

func readUserFromRow(rs rowScanner) (*entity.User, error) {
	user := &entity.User{}
	err := rs.Scan(&user.Id, &user.Username, &user.Tocken)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type userStorage struct {
	db *sql.DB
}

func NewUserStorage(conf *psql_conf.Config) (storage.UserStorage, error) {
	db, err := psql.GetDb(conf)
	if err != nil {
		return nil, err
	}
	return &userStorage{
		db: db,
	}, nil
}

func (st *userStorage) getUserWithWhereCase(wherecase string) (*entity.User, error) {
	stmt := `SELECT id, username, tocken FROM users`
	if wherecase != "" {
		stmt += " WHERE " + wherecase
	}

	result := st.db.QueryRow(stmt)
	return readUserFromRow(result)
}

func (st *userStorage) GetUserByUserName(username string) (*entity.User, error) {
	return st.getUserWithWhereCase(fmt.Sprintf(`username = '%s'`, username))
}

func (st *userStorage) GetUserById(id int64) (*entity.User, error) {
	return st.getUserWithWhereCase(fmt.Sprintf(`id = %d`, id))
}

func (st *userStorage) GetUserByTocken(tocken string) (*entity.User, error) {
	return st.getUserWithWhereCase(fmt.Sprintf(`tocken = '%s'`, tocken))
}

func (st *userStorage) AddUser(dto *dto.UserAddDTO) (*entity.User, error) {
	stmt := fmt.Sprintf(`INSERT INTO users (username, tocken, password) VALUES %s`, dto.ExtractInsertSQL())
	var id int64
	err := st.db.QueryRow(stmt + "RETURNING id;").Scan(&id)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Id:       id,
		Username: dto.UserName,
		Tocken:   dto.Tocken,
		Password: dto.Password,
	}, nil
}