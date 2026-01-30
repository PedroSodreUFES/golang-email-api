package models

import (
	"context"
	"main/internal/emails/DTO/requests"
	"main/internal/emails/DTO/responses"
)

type EmailService interface {
	DeleteEmail(ctx context.Context, emailId, userId int32) error
	SendEmail(ctx context.Context, request *requests.SendEmailRequest, senderID int32) (*Email, error)
	GetMyReceivedEmails(ctx context.Context, id int32) ([]responses.MyReceivedEmailsResponse, error)
	GetMySentEmails(ctx context.Context, id int32) ([]responses.MySentEmailsResponse, error)
	GetEmailById(ctx context.Context, userId, emailId int32) (*Email, error)
}