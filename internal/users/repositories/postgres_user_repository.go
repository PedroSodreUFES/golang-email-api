package repositories

import (
	"context"
	"main/internal/store/pgstore"
	"main/internal/users/DTO/requests"
	"main/internal/users/models"

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

// ChangeProfilePicture implements [models.UserRepository].
func (p *PostgresqlUserRepository) ChangeProfilePicture(ctx context.Context, id int32, profile_picture_link string) error {
	panic("unimplemented")
}

func (p *PostgresqlUserRepository) CreateUser(ctx context.Context, request *requests.CreateUserRequest) (*models.User, error) {
	user, err := p.queries.CreateUser(ctx, pgstore.CreateUserParams{
		FullName: request.FullName,
		Email: request.Email,
		PasswordHash: request.Password,
	})

	if err != nil {
		return nil, err
	}

	return &models.User{
		ID: user.ID,
		FullName: user.FullName,
		PasswordHash: "",
		Email: user.Email,
		ProfilePicture: user.ProfilePicture.String,
	}, nil
}

// DeleteUserById implements [models.UserRepository].
func (p *PostgresqlUserRepository) DeleteUserById(ctx context.Context, id int32) error {
	panic("unimplemented")
}

// FindUserByEmail implements [models.UserRepository].
func (p *PostgresqlUserRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	panic("unimplemented")
}

// FindUserById implements [models.UserRepository].
func (p *PostgresqlUserRepository) FindUserById(ctx context.Context, id int32) (*models.User, error) {
	panic("unimplemented")
}
