package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFactorial(t *testing.T) {
	numToFactorial := 5

	expectedRes := 120
	actualRes := factorial(numToFactorial)

	assert.Equal(t, expectedRes, actualRes)
}
