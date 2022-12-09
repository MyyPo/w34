package auth

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h Hasher) HashSecret(secret string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(secret), 12)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (h Hasher) CompareWithSecret(secret, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err
}
