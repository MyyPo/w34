package validators

import (
	"net/mail"
)

type AuthValidator struct {
}

func NewAuthValidator() *AuthValidator {
	return &AuthValidator{}
}

func (v AuthValidator) ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
