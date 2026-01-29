package requests

import (
	"context"
	"main/internal/validator"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator
	
	eval.CheckField(validator.NotNull(&req.Email), "email", "cannot be null")
	eval.CheckField(validator.NotBlank(req.Email), "email", "cannot be blank")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "must be a valid email")
	
	eval.CheckField(validator.NotNull(&req.Password), "password", "cannot be null")
	eval.CheckField(validator.NotBlank(req.Password), "password", "cannot be blank")

	return eval
}
