package dev_validator

import (
	"fmt"
	"regexp"
)

type DevValidator struct {
	projectNameRgx *regexp.Regexp
}

func NewDevValidator(projectNameRgxString string) (*DevValidator, error) {
	if projectNameRgxString == "" {
		projectNameRgxString = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
	}

	projectNameRgx, err := regexp.Compile(projectNameRgxString)
	if err != nil {
		return nil, err
	}

	return &DevValidator{
		projectNameRgx: projectNameRgx,
	}, nil
}

func (v DevValidator) ValidateProjectName(projectName string) error {
	if len(projectName) > 20 {
		return fmt.Errorf("project name must be shorter than 20 symbols")
	}
	if len(projectName) < 3 {
		return fmt.Errorf("project name must be longer than 2 symbols")
	}
	if !v.projectNameRgx.MatchString(projectName) {
		return fmt.Errorf("project name containts invalid characters")
	}

	return nil
}
