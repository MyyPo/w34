package dev_validator

import (
	"testing"
)

var devV, _ = NewDevValidator("", "")

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

func TestValidateOptions(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"invalid options key": testInvalidKeys,
		"valid options keys":  testValidKeys,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testInvalidKeys(t *testing.T) {
	invalidOptions := []map[string]string{
		{
			"-1": "ADD 1",
			"15": "ADD 1",
		},
		{
			"TI": "ADD 1",
			"1":  "ADD 1",
		},
	}

	for _, oMap := range invalidOptions {
		err := devV.ValidateOptions(oMap)
		if err == nil {
			t.Errorf("expected error validating invalid options key: %v", oMap)
		}
	}
}
func testValidKeys(t *testing.T) {
	validOptions := []map[string]string{
		{
			"0":    "ADD 1",
			"1":    "ADD 1",
			"2":    "ADD 1",
			"T409": "ADD 1",
			"I555": "ADD 1",
		},
	}

	for _, oMap := range validOptions {
		err := devV.ValidateOptions(oMap)
		if err != nil {
			t.Errorf("got error validating valid keys: %v", err)
		}
	}
}
