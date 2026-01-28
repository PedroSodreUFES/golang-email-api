package repositories

import (
	"main/internal/store/pgstore"
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