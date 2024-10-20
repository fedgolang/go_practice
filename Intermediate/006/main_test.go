package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestExecFormula(t *testing.T) {
	d := 100
	expectedRes := 18

	actualRes := execFormula(d)

	assert.Equal(t, expectedRes, actualRes)
}
