package validators

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	authV := NewAuthValidator()

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
