package models

import (
	"context"
	"main/internal/users/DTO/requests"
	"main/internal/users/DTO/responses"
)

type UserService interface{
	AuthenticateUser(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse,error)
	DeleteUser(ctx context.Context, id int32) (error)
	GetMe(ctx context.Context, token string)
	SignUp(ctx context.Context, request *requests.CreateUserRequest) (*responses.SignUpResponse, error)
	UpdateUserPhoto(ctx context.Context, id int32) (error)
}