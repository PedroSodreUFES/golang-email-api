package requests

import (
	"context"
	"main/internal/validator"
)

type ChangeMyPhotoRequest struct {
	ProfilePicture string `json:"profile_picture"`
}

func (req ChangeMyPhotoRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotNull(&req.ProfilePicture), "profile_picture", "cannot be null")
	eval.CheckField(validator.NotBlank(req.ProfilePicture), "profile_picture", "cannot be blank")

	return eval
}
