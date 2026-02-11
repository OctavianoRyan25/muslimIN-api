package mapper

import (
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/request"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/response"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
)

func ToUserResponse(user *domain.User) *response.UserResponse {
	return &response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserDomain(user *request.RegisterRequest) *domain.User {
	return &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToLoginUserDomain(user *request.LoginRequest) *domain.User {
	return &domain.User{
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToRegisterResponse() *response.RegisterResponse {
	return &response.RegisterResponse{
		Message: "User registered successfully",
	}
}

func ToLoginResponse(token string) *response.LoginResponse {
	return &response.LoginResponse{
		Token: token,
	}
}
