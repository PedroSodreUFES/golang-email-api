package requests

import (
	"context"
	"main/internal/validator"
)

type SendEmailRequest struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	EmailReceiver string `json:"email_receiver"`
}

func (req SendEmailRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotNull(&req.Title), "title", "cannot be null")
	eval.CheckField(validator.NotBlank(req.Title), "title", "cannot be blank")

	eval.CheckField(validator.NotNull(&req.Content), "content", "cannot be null")
	eval.CheckField(validator.NotBlank(req.Content), "content", "cannot be blank")

	eval.CheckField(validator.NotNull(&req.EmailReceiver), "email_receiver", "cannot be null")
	eval.CheckField(validator.NotBlank(req.EmailReceiver), "email_receiver", "cannot be blank")
	eval.CheckField(validator.Matches(req.EmailReceiver, validator.EmailRX), "email_receiver", "must be a valid email")

	return eval
}
