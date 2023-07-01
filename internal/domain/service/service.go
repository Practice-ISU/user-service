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

func (sv *userService) VerifyUser(token string) (*entity.User, error) {
	return sv.storage.GetUserByToken(token)
}

func (sv *userService) RegisterUser(username, password string) (*entity.User, error) {
	result, _ := sv.storage.GetUserByUserName(username)

	if result != nil {
		return nil, fmt.Errorf("this user is already exists")
	}

	tokenBytes := sha256.Sum256([]byte(username))
	token := hex.EncodeToString(tokenBytes[:])
	passwordBytes := sha256.Sum256([]byte(password))
	passwordHash := hex.EncodeToString(passwordBytes[:])

	dto := &dto.UserAddDTO{
		UserName: username,
		Password: passwordHash,
		Token:    token,
	}
	return sv.storage.AddUser(dto)
}

func (sv *userService) LoginUser(username, password string) (*entity.User, error) {
	passwordBytes := sha256.Sum256([]byte(password))
	passwordHash := hex.EncodeToString(passwordBytes[:])
	return sv.storage.LoginUser(&dto.UserLoginDTO{UserName: username, Password: passwordHash})

}

func (sv *userService) GetUserById(id int64) (*entity.User, error) {
	return sv.storage.GetUserById(id)
}

func (sv *userService) GetUserByToken(token string) (*entity.User, error) {
	return sv.storage.GetUserByToken(token)
}
