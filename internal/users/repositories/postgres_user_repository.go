package repositories

import (
	"context"
	"errors"
	"main/internal/exceptions"
	"main/internal/store/pgstore"
	"main/internal/users/DTO/requests"
	"main/internal/users/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresqlUserRepository struct {
	queries *pgstore.Queries
}

func NewPostgreUserRepository(pool *pgxpool.Pool) models.UserRepository {
	return &PostgresqlUserRepository{
		queries: pgstore.New(pool),
	}
}

func (p *PostgresqlUserRepository) ChangeProfilePicture(ctx context.Context, id int32, profile_picture_link string) error {
	rows, err := p.queries.UpdateUserProfilePhoto(ctx, pgstore.UpdateUserProfilePhotoParams{
		ID:             id,
		ProfilePicture: pgtype.Text{String: profile_picture_link, Valid: true},
	})
	if err != nil {
		return err
	}

	if rows == 0 {
		return exceptions.ErrUserNotFound
	}

	return nil
}

func (p *PostgresqlUserRepository) CreateUser(ctx context.Context, request *requests.CreateUserRequest) (*models.User, error) {
	user, err := p.queries.CreateUser(ctx, pgstore.CreateUserParams{
		FullName:     request.FullName,
		Email:        request.Email,
		PasswordHash: request.Password,
	})
	if err != nil {
		// Se erro for de unique field
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, exceptions.ErrEmailShouldBeUnique
			}
		}
		// outro erro
		return nil, err
	}

	return &models.User{
		ID:             user.ID,
		FullName:       user.FullName,
		PasswordHash:   "",
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture.String,
	}, nil
}

func (p *PostgresqlUserRepository) DeleteUserById(ctx context.Context, id int32) error {
	rows, err := p.queries.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}

	if rows == 0 {
		return exceptions.ErrUserNotFound
	}

	return nil
}

func (p *PostgresqlUserRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := p.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, exceptions.ErrInvalidCredentials
		}
		return nil, err
	}

	return &models.User{
		ID:             user.ID,
		FullName:       user.FullName,
		PasswordHash:   user.PasswordHash,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture.String,
	}, nil
}

func (p *PostgresqlUserRepository) FindUserById(ctx context.Context, id int32) (*models.User, error) {
	user, err := p.queries.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, exceptions.ErrUserNotFound
		}
		return nil, err
	}

	return &models.User{
		ID:             user.ID,
		FullName:       user.FullName,
		PasswordHash:   user.PasswordHash,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture.String,
	}, nil
}
