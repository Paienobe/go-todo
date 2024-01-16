package types

type RegistrationParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
