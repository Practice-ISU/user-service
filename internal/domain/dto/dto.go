package dto

import "fmt"

type UserAddDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Tocken   string `json:"tocken"`
}

func (dto *UserAddDTO) ExtractInsertSQL() string {
	return fmt.Sprintf(`('%s', '%s', '%s')`, dto.UserName, dto.Password, dto.Tocken)
}

type UserToken struct {
	Tocken string `json:"tocken"`
}

type UserVerifyAnnswer struct {
	Success  bool   `json:"success"`
	Id       int64  `json:"id"`
	UserName string `json:"username"`
}
