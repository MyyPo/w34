package dev_validator

import (
	"fmt"
	"regexp"
	"strconv"
)

type DevValidator struct {
	nameRgx       *regexp.Regexp
	optionsKeyRgx *regexp.Regexp
	optionsValRgx *regexp.Regexp
}

func NewDevValidator(nameRgxString, optionsKeyRgxString, optionsValRgxString string) (*DevValidator, error) {
	if nameRgxString == "" {
		nameRgxString = "^[a-zA-Z0-9 ]+(?:-[a-zA-Z0-9]+)*$"
	}

	nameRgx, err := regexp.Compile(nameRgxString)
	if err != nil {
		return nil, err
	}

	if optionsKeyRgxString == "" {
		// Default rgx making sure invalid option keys can't be passed
		optionsKeyRgxString = "^[TI]([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9])$"
	}
	optionsKeyRgx, err := regexp.Compile(optionsKeyRgxString)
	if err != nil {
		return nil, err
	}

	if optionsValRgxString == "" {
		// Prevents users from passing invalid commands to scene options
		optionsValRgxString = "^(AD|NE) ([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9])$"
	}
	optionsValRgx, err := regexp.Compile(optionsValRgxString)
	if err != nil {
		return nil, err
	}

	return &DevValidator{
		nameRgx:       nameRgx,
		optionsKeyRgx: optionsKeyRgx,
		optionsValRgx: optionsValRgx,
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
	for key, val := range options {
		if !v.validateKey(key) {
			return fmt.Errorf("invalid key: %v", key)
		}

		if !v.optionsValRgx.MatchString(val) {
			return fmt.Errorf("invalid value: %v", val)
		}
	}

	return nil
}

func (v DevValidator) validateKey(key string) bool {
	intKey, err := strconv.ParseInt(key, 10, 32)
	// If it isn't a valid number, try to match it with rgx
	if err != nil {
		return v.optionsKeyRgx.MatchString(key)
	}
	// Otherwise we assue that it is an unconditionally available option
	if intKey >= 100 || intKey < 0 {
		return false
	}

	return true
}
