package models

import (
	"context"
	"main/internal/users/DTO/requests"
)

type UserRepository interface{
	CreateUser(ctx context.Context, request *requests.CreateUserRequest) (*User, error)
	DeleteUserById(ctx context.Context, id int32) (error)
	FindUserById(ctx context.Context,id int32) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	ChangeProfilePicture(ctx context.Context, id int32, profile_picture_link string) (error)
}