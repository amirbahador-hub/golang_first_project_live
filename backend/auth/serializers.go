package auth

type UserRequest struct {
	Name string `json:"name"`
	Role string  `json:"role"`
	Password string  `json:"password"`
}

type LoginRequest struct {
	Name string `json:"name"`
	Password string  `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type VerifyRequest struct {
	Token string `json:"token"`
}

type VerifyResponse struct {
	IsValid bool `json:"isValid"`
}