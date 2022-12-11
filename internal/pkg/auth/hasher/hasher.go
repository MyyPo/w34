package hasher

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

func (h Hasher) CompareWithSecret(hash, secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
}
