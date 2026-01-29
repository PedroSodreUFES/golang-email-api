package models

import (
	"context"
	"main/internal/users/DTO/requests"
	"main/internal/users/DTO/responses"
	"mime/multipart"
)

type UserService interface{
	AuthenticateUser(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse,error)
	DeleteUser(ctx context.Context, id int32) (error)
	GetMe(ctx context.Context, id int32) (*responses.MeResponse, error)
	SignUp(ctx context.Context, request *requests.CreateUserRequest) (*responses.SignUpResponse, error)
	UpdateUserPhoto(ctx context.Context, id int32, file *multipart.FileHeader) error
	DeleteUserPhoto(ctx context.Context, id int32) error
}