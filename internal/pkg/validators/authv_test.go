package validators

import (
	"testing"

	authv1 "github.com/MyyPo/w34.Go/gen/go/auth/v1"
)

const usernameRegex = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"

var authV, _ = NewAuthValidator(60, usernameRegex)

func TestValidateCredentials(t *testing.T) {
	t.Parallel()
	t.Run("Correct credentials", func(t *testing.T) {
		req := authv1.SignUpRequest{
			Username: "Hey",
			Email:    "val@g.com",
			Password: "dqkrm23rmm9QM",
		}

		err := authV.ValidateCredentials(&req)
		if err != nil {
			t.Errorf("Correct credentials were seen as invalid %q", err)
		}
	})
	t.Run("Incorrect credentials", func(t *testing.T) {
		invalidEmail := "infv@"

		req := authv1.SignUpRequest{
			Username: "Hey",
			Email:    invalidEmail,
			Password: "dkJ2C3PRU093dm",
		}

		err := authV.ValidateCredentials(&req)
		if err == nil {
			t.Errorf("Incorrect email validated")
		}

	})
}

func TestValidateUsername(t *testing.T) {
	t.Parallel()
	t.Run("Too long username", func(t *testing.T) {
		longName := "WWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWW"
		err := authV.ValidateUsername(longName)
		if err == nil {
			t.Errorf("Too long username was validated as normal")
		}
	})
	t.Run("Too short username", func(t *testing.T) {
		shortName := "ww"
		err := authV.ValidateUsername(shortName)
		if err == nil {
			t.Errorf("Too short username was validated as normal")
		}
	})
	t.Run("Username with invalid symbols", func(t *testing.T) {
		invalidName := "#@$!C0)"
		err := authV.ValidateUsername(invalidName)
		if err == nil {
			t.Errorf("Username with invalid symbols validated as normal")
		}
	})
}

func TestValidatePassword(t *testing.T) {
	t.Parallel()
	t.Run("Pass safe password", func(t *testing.T) {
		safePsw := "QWDLKQqcmw3;r3uEQWAs"
		err := authV.ValidatePassword(safePsw)
		if err != nil {
			t.Errorf("Fail to validate safe password, %q", err)
		}

	})
	t.Run("Put unsafe password", func(t *testing.T) {
		unsafePsw := "PASSWORD"
		err := authV.ValidatePassword(unsafePsw)
		if err == nil {
			t.Errorf("Unsafe password passed as safe")
		}
	})
}

func TestValidateEmail(t *testing.T) {
	t.Parallel()
	t.Run("Pass valid email", func(t *testing.T) {
		validEmail := "magic@g.com"
		err := authV.ValidateEmail(validEmail)
		if err != nil {
			t.Errorf("Valid email failed to validate: %q", err)
		}
	})
	t.Run("Pass incorrect email", func(t *testing.T) {
		invalidEmail := "magic@"
		err := authV.ValidateEmail(invalidEmail)
		if err == nil {
			t.Errorf("Invalid email got verified")
		}
	})
}
