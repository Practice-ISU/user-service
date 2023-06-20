package storage

import (
	"user-service/internal/domain/dto"
	"user-service/internal/domain/entity"
)

type UserStorage interface {
	GetUserByUserName(username string) (*entity.User, error)
	GetUserById(id int64) (*entity.User, error)
	GetUserByTocken(tocken string) (*entity.User, error)

	AddUser(dto *dto.UserAddDTO) (*entity.User, error)
}
