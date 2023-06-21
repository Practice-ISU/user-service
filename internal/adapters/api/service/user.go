package service

import (
	"user-service/internal/domain/entity"
)

type UserService interface {
	RegisterUser(username string, password string) (*entity.User, error)
	VerifyUser(tocken string) (*entity.User, error)
	LoginUser(username, password string) (*entity.User, error)
	
	GetUserById(id int64) (*entity.User, error)
	GetUserByToken(tocken string) (*entity.User, error)
}
