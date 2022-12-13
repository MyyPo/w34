package validators

import (
	"fmt"
	"net/mail"
	"regexp"

	pwv "github.com/wagslane/go-password-validator"
)

type AuthValidator struct {
	minEntropy  float64
	usernameRgx *regexp.Regexp
}

func NewAuthValidator(minEntropy float64, usernameRgxString string) (*AuthValidator, error) {
	if usernameRgxString == "" {
		usernameRgxString = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
	}

	usernameRgx, err := regexp.Compile(usernameRgxString)
	if err != nil {
		return nil, err
	}

	return &AuthValidator{
		minEntropy:  minEntropy,
		usernameRgx: usernameRgx,
	}, nil
}

func (v AuthValidator) ValidateSignUpCredentials(
	username string,
	email string,
	password string,
) error {
	if err := v.ValidateUsername(username); err != nil {
		return err
	}
	if err := v.ValidateEmail(email); err != nil {
		return err
	}
	if err := v.ValidatePassword(password); err != nil {
		return err
	}

	return nil
}

func (v AuthValidator) ValidateUsername(username string) error {
	if len(username) > 20 {
		return fmt.Errorf("username must be shorter than 20 symbols")
	}
	if len(username) < 3 {
		return fmt.Errorf("username must be longer than 2 symbols")
	}
	// try to match username with our regex
	if !v.usernameRgx.MatchString(username) {
		return fmt.Errorf("username contains invalid characters")
	}
	return nil

}

func (v AuthValidator) ValidatePassword(password string) error {
	if err := pwv.Validate(password, v.minEntropy); err != nil {
		return err
	}
	return nil
}

func (v AuthValidator) ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
