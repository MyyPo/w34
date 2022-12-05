package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	hasher := NewHasher()

	psw := "Sill3pASsworD"
	var pswHash string

	t.Run("Hash the password, make sure that it has been turned into hash", func(t *testing.T) {

		got, err := hasher.Hash(psw)
		if err != nil {
			t.Errorf("error while trying to hash a password: %q", err)
		}
		assert.NotEqual(t, psw, got)
		pswHash = got
	})
	t.Run("Compare the password and VALID hash for it", func(t *testing.T) {
		err := hasher.Compare(psw, pswHash)
		if err != nil {
			t.Errorf("the valid password wasn't equal to its hash: %q", err)
		}
	})
	t.Run("Try INVALID password for hash", func(t *testing.T) {
		invalidPsw := "SillyPassword"
		err := hasher.Compare(invalidPsw, pswHash)
		if err == nil {
			t.Errorf("hasher didn't raise the error for invalid password")
		}
	})
}
