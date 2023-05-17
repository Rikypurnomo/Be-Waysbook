package authdto

type RegisterResponse struct {
	Email string `json:"email"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
	Photo string `json:"photo"`
}
