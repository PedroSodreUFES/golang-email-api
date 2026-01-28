package jsonutils

import (
	"encoding/json"
	"errors"
	"main/internal/validator"
	"net/http"
)

var (
	ErrFailedToDecodeJson = errors.New("failed to decode json")
	ErrFailedToValidateJson = errors.New("failed to validate json")
)


func DecodeValidJson[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var data T
	
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, ErrFailedToDecodeJson
	}

	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, ErrFailedToValidateJson
	}

	return data, nil, nil
}