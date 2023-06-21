package dto

import "fmt"

type UserAddDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Token   string `json:"token"`
}

func (dto *UserAddDTO) ExtractInsertSQL() string {
	return fmt.Sprintf(`('%s', '%s', '%s')`, dto.UserName, dto.Token, dto.Password)
}

type UserLoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (dto *UserLoginDTO) ExtractWhereSQL() string {
	return fmt.Sprintf(`username = '%s' AND password = '%s'`, dto.UserName, dto.Password)
}

type UserToken struct {
	Token string `json:"token"`
}

type UserVerifyAnnswer struct {
	Success  bool   `json:"success"`
	Id       int64  `json:"id"`
	UserName string `json:"username"`
}
