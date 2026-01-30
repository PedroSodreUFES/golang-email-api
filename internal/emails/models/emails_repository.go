package models

import (
	"context"
	"main/internal/emails/DTO/responses"
)

type EmailsRepository interface {
	DeleteEmailById(ctx context.Context, id int32) (error)
	GetEmailById(ctx context.Context, id int32) (*Email, error)
	CreateEmail(ctx context.Context, receiverId, senderId int32, title, content string) (*Email, error)
	GetMyReceivedEmails(ctx context.Context, id int32) ([]responses.MyReceivedEmailsResponse, error)
	GetMySentEmails(ctx context.Context, id int32) ([]responses.MySentEmailsResponse, error)
	UpdateEmailById(ctx context.Context, id int32) (*Email, error)
}