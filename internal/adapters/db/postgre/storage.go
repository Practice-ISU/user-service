package postgre

import (
	"database/sql"
	"fmt"
	"user-service/internal/domain/dto"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/storage"
)

type DbClient interface {
	GetDb() (*sql.DB, error)
}

type rowScanner interface {
	Scan(dest ...any) error
}

func readUserFromRow(rs rowScanner) (*entity.User, error) {
	user := &entity.User{}
	err := rs.Scan(&user.Id, &user.Username, &user.Token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type userStorage struct {
	dbClient DbClient
}

func NewUserStorage(dbClient DbClient) storage.UserStorage {
	return &userStorage{
		dbClient: dbClient,
	}
}

func (st *userStorage) getUserWithWhereCase(wherecase string) (*entity.User, error) {
	db, err := st.dbClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt := `SELECT id, username, token FROM users`
	if wherecase != "" {
		stmt += " WHERE " + wherecase
	}
	result := db.QueryRow(stmt)
	return readUserFromRow(result)
}

func (st *userStorage) GetUserByUserName(username string) (*entity.User, error) {
	return st.getUserWithWhereCase(fmt.Sprintf(`username = '%s'`, username))
}

func (st *userStorage) GetUserById(id int64) (*entity.User, error) {
	return st.getUserWithWhereCase(fmt.Sprintf(`id = %d`, id))
}

func (st *userStorage) GetUserByToken(token string) (*entity.User, error) {
	return st.getUserWithWhereCase(fmt.Sprintf(`token = '%s'`, token))
}

func (st *userStorage) AddUser(dto *dto.UserAddDTO) (*entity.User, error) {
	db, err := st.dbClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt := fmt.Sprintf(`INSERT INTO users (username, token, password) VALUES %s`, dto.ExtractInsertSQL())
	var id int64
	err = db.QueryRow(stmt + "RETURNING id;").Scan(&id)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Id:       id,
		Username: dto.UserName,
		Token:    dto.Token,
		Password: dto.Password,
	}, nil
}

func (st *userStorage) LoginUser(dto *dto.UserLoginDTO) (*entity.User, error) {
	return st.getUserWithWhereCase(dto.ExtractWhereSQL())
}
