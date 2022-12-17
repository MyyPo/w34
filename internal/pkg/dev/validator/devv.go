package dev_validator

import (
	"fmt"
	"regexp"
)

type DevValidator struct {
	nameRgx *regexp.Regexp
}

func NewDevValidator(nameRgxString string) (*DevValidator, error) {
	if nameRgxString == "" {
		nameRgxString = "^[a-zA-Z0-9 ]+(?:-[a-zA-Z0-9]+)*$"
	}

	nameRgx, err := regexp.Compile(nameRgxString)
	if err != nil {
		return nil, err
	}

	return &DevValidator{
		nameRgx: nameRgx,
	}, nil
}

func (v DevValidator) ValidateName(projectName string) error {
	if len(projectName) > 50 {
		return fmt.Errorf("name must be shorter than 20 symbols")
	}
	if len(projectName) == 0 {
		return fmt.Errorf("name can't be empty")
	}
	if !v.nameRgx.MatchString(projectName) {
		return fmt.Errorf("name containts invalid characters")
	}

	return nil
}
