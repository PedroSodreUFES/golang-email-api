package service

import (
	"context"
	"errors"
	"main/internal/exceptions"
	"main/internal/users/DTO/requests"
	"main/internal/users/DTO/responses"
	"main/internal/users/models"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserService) SignUp(ctx context.Context, request *requests.CreateUserRequest) (*responses.SignUpResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	if err != nil {
		return nil, err
	}
	request.Password = string(passwordHash)

	user, err := u.userRepository.CreateUser(ctx, request)
	if err != nil {
		// Se erro for de unique field
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, exceptions.ErrEmailShouldBeUnique
			}
		}
		// se error não for de unique field
		return nil, err
	}

	return &responses.SignUpResponse{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		ProfilePicture: nil,
	}, nil
}

// UpdateUserPhoto implements [models.UserService].
func (u *UserService) UpdateUserPhoto(ctx context.Context, id int32) error {
	panic("unimplemented")
}
