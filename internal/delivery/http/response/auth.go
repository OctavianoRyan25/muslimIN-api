package response

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
