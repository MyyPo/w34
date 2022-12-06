package validators

import (
	"fmt"
	"net/mail"
	"regexp"

	pwv "github.com/wagslane/go-password-validator"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
)

type AuthValidator struct {
	minEntropy  float64
	usernameRgx *regexp.Regexp
}

func NewAuthValidator(minEntropy float64, usernameRgxString string) (*AuthValidator, error) {
	usernameRgx, err := regexp.Compile(usernameRgxString)
	if err != nil {
		return nil, err
	}

	return &AuthValidator{
		minEntropy:  minEntropy,
		usernameRgx: usernameRgx,
	}, nil
}

func (v AuthValidator) ValidateCredentials(req *authv1.SignUpRequest) error {
	if err := v.ValidateEmail(req.GetEmail()); err != nil {
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
