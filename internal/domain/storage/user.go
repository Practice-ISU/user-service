package storage

import (
	"user-service/internal/domain/dto"
	"user-service/internal/domain/entity"
)

type UserStorage interface {
	GetUserByUserName(username string) (*entity.User, error)
	GetUserById(id int64) (*entity.User, error)
	GetUserByToken(token string) (*entity.User, error)

	LoginUser(dto *dto.UserLoginDTO) (*entity.User, error)

	AddUser(dto *dto.UserAddDTO) (*entity.User, error)
}
