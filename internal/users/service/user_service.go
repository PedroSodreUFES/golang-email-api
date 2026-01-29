package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"main/internal/auth"
	"main/internal/exceptions"
	imagestore "main/internal/store/image_store"
	"main/internal/users/DTO/requests"
	"main/internal/users/DTO/responses"
	"main/internal/users/models"
	"main/internal/utils"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository models.UserRepository
	jwtMaker       auth.JWTMaker
	imageStore     imagestore.ImageStore
}

func NewUserService(userRepository models.UserRepository, jwtMaker auth.JWTMaker, imageStore imagestore.ImageStore) models.UserService {
	return &UserService{
		userRepository: userRepository,
		jwtMaker:       jwtMaker,
		imageStore:     imageStore,
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

func (u *UserService) UpdateUserPhoto(ctx context.Context, id int32, file *multipart.FileHeader) error {
	user, err := u.userRepository.FindUserById(ctx, id)
	if err != nil {
		return err
	}
	old := strings.TrimSpace(user.ProfilePicture)
	
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	rs, ok := f.(io.ReadSeeker)
	if !ok {
		return fmt.Errorf("file is not seekable")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	key := fmt.Sprintf("users/%d/profile_%d%s", id, time.Now().UnixNano(), ext)

	contentType := "application/octet-stream"
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".webp":
		contentType = "image/webp"
	}

	publicURL, err := u.imageStore.Upload(ctx, key, rs, file.Size, contentType)
	if err != nil {
		_ = u.imageStore.Delete(ctx, key)
		return err
	}

	valueToStore := key
	if publicURL != "" {
		valueToStore = publicURL
	}

	if old != "" {
		fmt.Println(old)
		oldKey := utils.ExtractR2Key(old)
		if oldKey != "" && oldKey != key {
			_ = u.imageStore.Delete(ctx, oldKey)
		}
	}

	return u.userRepository.ChangeProfilePicture(ctx, id, valueToStore)
}

