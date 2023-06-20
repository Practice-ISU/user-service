package service

import (
	"user-service/internal/domain/entity"
)

type UserService interface {
	AddUser(username string, password string) (*entity.User, error)
	VerifyUser(tocken string) (*entity.User, error)
}
