package dev_validator

import (
	"fmt"
	"regexp"
	"strconv"
)

type DevValidator struct {
	nameRgx    *regexp.Regexp
	optionsRgx *regexp.Regexp
}

func NewDevValidator(nameRgxString, optionsRgxString string) (*DevValidator, error) {
	if nameRgxString == "" {
		nameRgxString = "^[a-zA-Z0-9 ]+(?:-[a-zA-Z0-9]+)*$"
	}

	nameRgx, err := regexp.Compile(nameRgxString)
	if err != nil {
		return nil, err
	}

	if optionsRgxString == "" {
		optionsRgxString = "[TI]([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9])"
	}
	optionsRgx, err := regexp.Compile(optionsRgxString)
	if err != nil {
		return nil, err
	}

	return &DevValidator{
		nameRgx:    nameRgx,
		optionsRgx: optionsRgx,
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

func (v DevValidator) ValidateOptions(options map[string]string) error {
	for key := range options {
		if !validateInt(key) && !v.optionsRgx.MatchString(key) {
			return fmt.Errorf("invalid key: %v", key)
		}
	}

	return nil
}

func validateInt(key string) bool {
	intKey, err := strconv.ParseInt(key, 10, 32)
	if err != nil {
		return false
	}
	if intKey >= 100 || intKey < 0 {
		return false
	}
	return true
}
