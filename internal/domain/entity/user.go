package entity

type User struct {
	Id        int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Tocken  string `json:"tocken"`
}
