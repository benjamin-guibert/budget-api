package tests

import (
	"testing"

	"gopkg.in/go-playground/validator.v9"
	"gotest.tools/assert"
)

func ValidateModel(t *testing.T, expected bool, model interface{}) {
	err := validator.New().Struct(model)

	assert.Assert(t, (err == nil) == expected, "validation invalid (expected %t) %v", expected, err)
}
