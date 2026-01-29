package service

import (
	"context"
	"errors"
	"main/internal/auth"
	"main/internal/exceptions"
	"main/internal/users/DTO/requests"
	"main/internal/users/DTO/responses"
	"main/internal/users/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository models.UserRepository
	jwtMaker       auth.JWTMaker
}

func NewUserService(userRepository models.UserRepository, jwtMaker auth.JWTMaker) models.UserService {
	return &UserService{
		userRepository: userRepository,
		jwtMaker:       jwtMaker,
	}
}

func (u *UserService) AuthenticateUser(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error) {
	user, err := u.userRepository.FindUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, exceptions.ErrInvalidCredentials
		}
		return nil, err
	}

	token, err := u.jwtMaker.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &responses.LoginResponse{
		Token: token,
	}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, id int32) error {
	err := u.userRepository.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetMe(ctx context.Context, id int32) (*responses.MeResponse, error) {
	user, err := u.userRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &responses.MeResponse{
		ID:             user.ID,
		FullName:       user.FullName,
		Email:          user.Email,
		ProfilePicture: &user.ProfilePicture,
	}, nil
}

func (u *UserService) SignUp(ctx context.Context, request *requests.CreateUserRequest) (*responses.SignUpResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	if err != nil {
		return nil, err
	}
	request.Password = string(passwordHash)

	user, err := u.userRepository.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return &responses.SignUpResponse{
		ID:             user.ID,
		FullName:       user.FullName,
		Email:          user.Email,
		ProfilePicture: nil,
	}, nil
}

func (u *UserService) UpdateUserPhoto(ctx context.Context, id int32, new_photo string) error {
	err := u.userRepository.ChangeProfilePicture(ctx, id, new_photo)
	if err != nil {
		return err
	}

	return nil
}
