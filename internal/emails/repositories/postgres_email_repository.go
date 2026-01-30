package repositories

import (
	"context"
	"errors"
	"main/internal/emails/DTO/responses"
	"main/internal/emails/models"
	"main/internal/exceptions"
	"main/internal/store/pgstore"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresqlEmailRepository struct {
	queries *pgstore.Queries
}

func NewPostgreEmailRepository(pool *pgxpool.Pool) models.EmailsRepository {
	return &PostgresqlEmailRepository{
		queries: pgstore.New(pool),
	}
}

func (p *PostgresqlEmailRepository) DeleteEmailById(ctx context.Context, id int32) error {
	rows, err := p.queries.DeleteEmail(ctx, id)
	if err != nil {
		return err
	}

	if rows == 0 {
		return exceptions.ErrEmailNotFound
	}

	return nil
}

func (p *PostgresqlEmailRepository) GetEmailById(ctx context.Context, id int32) (*models.Email, error) {
	email, err := p.queries.GetEmailById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, exceptions.ErrEmailNotFound
		}
		return nil, err
	}

	return &models.Email{
		ID:         email.ID,
		Title:      email.Title,
		Content:    email.Content.String,
		Wasseen:    email.Wasseen.Bool,
		CreatedAt:  email.CreatedAt.Time,
		IDReceiver: email.IDReceiver,
		IDSender:   email.IDSender,
	}, nil
}

func (p *PostgresqlEmailRepository) CreateEmail(ctx context.Context, receiverId int32, senderId int32, title string, content string) (*models.Email, error) {
	email, err := p.queries.CreateEmail(ctx, pgstore.CreateEmailParams{
		Title:      title,
		Content:    pgtype.Text{String: content, Valid: true},
		IDReceiver: receiverId,
		IDSender:   senderId,
	})
	if err != nil {
		return nil, err
	}

	return &models.Email{
		ID:         email.ID,
		Title:      email.Title,
		Content:    email.Content.String,
		Wasseen:    email.Wasseen.Bool,
		CreatedAt:  email.CreatedAt.Time,
		IDReceiver: email.IDReceiver,
		IDSender:   email.IDSender,
	}, nil
}

func (p *PostgresqlEmailRepository) GetMyReceivedEmails(ctx context.Context, id int32) ([]responses.MyReceivedEmailsResponse, error) {
	result, err := p.queries.GetMyReceivedEmails(ctx, id)
	if err != nil {
		return nil, err
	}

	response := make([]responses.MyReceivedEmailsResponse, 0, len(result))
	for _, row := range result {
		content := ""
		if row.Content.Valid {
			content = row.Content.String
		}

		profilePicture := ""
		if row.ProfilePicture.Valid {
			profilePicture = row.ProfilePicture.String
		}

		response = append(response, responses.MyReceivedEmailsResponse{
			ID:             row.ID,
			Title:          row.Title,
			Content:        content,
			Wasseen:        row.Wasseen.Bool,
			CreatedAt:      row.CreatedAt.Time,
			IDReceiver:     row.IDReceiver,
			IDSender:       row.IDSender,
			FullName:       row.FullName,
			Email:          row.Email,
			ProfilePicture: profilePicture,
		})
	}

	return response, nil
}

func (p *PostgresqlEmailRepository) GetMySentEmails(ctx context.Context, id int32) ([]responses.MySentEmailsResponse, error) {
	result, err := p.queries.GetMySentEmails(ctx, id)
	if err != nil {
		return nil, err
	}

	response := make([]responses.MySentEmailsResponse, 0, len(result))
	for _, row := range result {
		content := ""
		if row.Content.Valid {
			content = row.Content.String
		}

		profilePicture := ""
		if row.ProfilePicture.Valid {
			profilePicture = row.ProfilePicture.String
		}

		response = append(response, responses.MySentEmailsResponse{
			ID:             row.ID,
			Title:          row.Title,
			Content:        content,
			Wasseen:        row.Wasseen.Bool,
			CreatedAt:      row.CreatedAt.Time,
			IDReceiver:     row.IDReceiver,
			IDSender:       row.IDSender,
			FullName:       row.FullName,
			Email:          row.Email,
			ProfilePicture: profilePicture,
		})
	}

	return response, nil
}

func (p *PostgresqlEmailRepository) UpdateEmailById(ctx context.Context, id int32) (*models.Email, error) {
	email, err := p.queries.UpdateAndGetEmailByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Email{
		ID:         email.ID,
		Title:      email.Title,
		Content:    email.Content.String,
		Wasseen:    email.Wasseen.Bool,
		CreatedAt:  email.CreatedAt.Time,
		IDReceiver: email.IDReceiver,
		IDSender:   email.IDSender,
	}, nil	
}