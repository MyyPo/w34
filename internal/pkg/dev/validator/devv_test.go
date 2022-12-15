package dev_validator

import (
	"testing"
)

var devV, _ = NewDevValidator("")

func TestValidateProjectName(t *testing.T) {
	t.Parallel()
	t.Run("Valid project name", func(t *testing.T) {
		projectName := "Hello"
		err := devV.ValidateProjectName(projectName)
		if err != nil {
			t.Errorf("valid project name seen as invalid: %v", err)
		}
	})
	t.Run("Invalid project name passed", func(t *testing.T) {
		invalidProjectName := "''///o"
		err := devV.ValidateProjectName(invalidProjectName)
		if err == nil {
			t.Errorf("invalid project name validated")
		}
	})
}
