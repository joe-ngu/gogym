package types

type LoginResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
