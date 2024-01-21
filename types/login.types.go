package types

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSuccess[T any] struct {
	Success bool   `json:"success"`
	Tasks   []T    `json:"tasks"`
	Token   string `json:"token"`
}
