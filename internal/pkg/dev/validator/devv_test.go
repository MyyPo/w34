package dev_validator

import (
	"testing"
)

var devV, _ = NewDevValidator("")

func TestValidateProjectName(t *testing.T) {
	t.Parallel()
	t.Run("Valid project name", func(t *testing.T) {
		projectName := "Hello"
		err := devV.ValidateName(projectName)
		if err != nil {
			t.Errorf("valid project name seen as invalid: %v", err)
		}
	})
	t.Run("Invalid project name passed", func(t *testing.T) {
		invalidProjectName := "''///o"
		err := devV.ValidateName(invalidProjectName)
		if err == nil {
			t.Errorf("invalid project name validated")
		}
	})
}
