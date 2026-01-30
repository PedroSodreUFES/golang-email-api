package service

import (
	"context"
	"errors"
	"main/internal/emails/DTO/requests"
	"main/internal/emails/DTO/responses"
	"main/internal/emails/models"
	"main/internal/exceptions"
	userModels "main/internal/users/models"
)

type EmailService struct {
	emailsRepository models.EmailsRepository
	usersRepository  userModels.UserRepository
}

func NewEmailService(emailsRepository models.EmailsRepository, usersRepository userModels.UserRepository) models.EmailService {
	return &EmailService{
		emailsRepository: emailsRepository,
		usersRepository:  usersRepository,
	}
}

func (e *EmailService) DeleteEmail(ctx context.Context, emailId int32, userId int32) error {
	email, err := e.emailsRepository.GetEmailById(ctx, emailId)
	if err != nil {
		return err
	}

	if email.IDSender != userId || email.Wasseen == true {
		return exceptions.ErrNotAllowed
	}

	err = e.emailsRepository.DeleteEmailById(ctx, emailId)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmailService) SendEmail(ctx context.Context, request *requests.SendEmailRequest, senderID int32) (*models.Email, error) {
	userReceiver, err := e.usersRepository.FindUserByEmail(ctx, request.EmailReceiver)
	if err != nil {
		if errors.Is(err, exceptions.ErrInvalidCredentials) {
			return nil, exceptions.ErrUserNotFound
		}
		return nil, err
	}

	email, err := e.emailsRepository.CreateEmail(ctx, userReceiver.ID, senderID, request.Title, request.Content)
	if err != nil {
		return nil, err
	}

	return email, nil
}

func (e *EmailService) GetMyReceivedEmails(ctx context.Context, id int32) ([]responses.MyReceivedEmailsResponse, error) {
	myReceivedEmails, err := e.emailsRepository.GetMyReceivedEmails(ctx, id)
	if err != nil {
		return nil, err
	}

	return myReceivedEmails, nil
}

func (e *EmailService) GetMySentEmails(ctx context.Context, id int32) ([]responses.MySentEmailsResponse, error) {
	mySentEmails, err := e.emailsRepository.GetMySentEmails(ctx, id)
	if err != nil {
		return nil, err
	}

	return mySentEmails, nil
}

func (e *EmailService) GetEmailById(ctx context.Context, userId int32, emailId int32) (*models.Email, error) {
	email, err := e.emailsRepository.GetEmailById(ctx, emailId)
	if err != nil {
		return nil, err
	}

	if email.IDReceiver != userId && email.IDSender != userId {
		return nil, exceptions.ErrNotAllowed
	}

	if email.Wasseen == false && email.IDReceiver == userId{
		email, err = e.emailsRepository.UpdateEmailById(ctx, emailId)
		if err != nil {
			return nil, err
		}
	}

	return email, nil
}
