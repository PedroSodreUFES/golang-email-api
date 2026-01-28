package service

import (
	"context"
	"main/internal/users/DTO/requests"
	"main/internal/users/DTO/responses"
	"main/internal/users/models"
)

type UserService struct {
	userRepository models.UserRepository
}

func NewUserService(userRepository models.UserRepository) models.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// AuthenticateUser implements [models.UserService].
func (u *UserService) AuthenticateUser(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error) {
	panic("unimplemented")
}

// DeleteUser implements [models.UserService].
func (u *UserService) DeleteUser(ctx context.Context, id int32) error {
	panic("unimplemented")
}

// GetMe implements [models.UserService].
func (u *UserService) GetMe(ctx context.Context, token string) {
	panic("unimplemented")
}

// SignUp implements [models.UserService].
func (u *UserService) SignUp(ctx context.Context, request *requests.CreateUserRequest) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUserPhoto implements [models.UserService].
func (u *UserService) UpdateUserPhoto(ctx context.Context, id int32) error {
	panic("unimplemented")
}
