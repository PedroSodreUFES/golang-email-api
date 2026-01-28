package requests

import (
	"context"
	"main/internal/validator"
)

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req CreateUserRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotNull(&req.FullName), "full_name", "cannot be null")
	eval.CheckField(validator.NotBlank(req.FullName), "full_name", "cannot be blank")

	eval.CheckField(validator.NotNull(&req.Password), "password", "cannot be null")
	eval.CheckField(validator.NotBlank(req.Password), "password", "cannot be blank")
	eval.CheckField(validator.MinChars(req.Password, 8), "password", "must have at least 8 characters")

	eval.CheckField(validator.NotNull(&req.Email), "email", "cannot be null")
	eval.CheckField(validator.NotBlank(req.Email), "email", "cannot be blank")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "must be a valid email")

	return eval
}
