package dtos

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTKey struct {
	Email  string `json:"email"`
	Key    string `json:"key"`
}
