package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"user-service/internal/domain/dto"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/storage"
)

type userService struct {
	storage storage.UserStorage
}

func NewUserService(storage storage.UserStorage) *userService {
	return &userService{
		storage: storage,
	}
}

func (sv *userService) VerifyUser(tocken string) (*entity.User, error) {
	return sv.storage.GetUserByTocken(tocken)
}

func (sv *userService) AddUser(username, password string) (*entity.User, error) {
	result, _ := sv.storage.GetUserByUserName(username)

	if result != nil {
		return nil, fmt.Errorf("this user is already exists")
	}

	tockenBytes := sha256.Sum256([]byte(username))
	token := hex.EncodeToString(tockenBytes[:])
	passwordBytes := sha256.Sum256([]byte(password))
	passwordHash := hex.EncodeToString(passwordBytes[:])

	dto := &dto.UserAddDTO{
		UserName: username,
		Password: passwordHash,
		Tocken:   token,
	}
	return sv.storage.AddUser(dto)
}
