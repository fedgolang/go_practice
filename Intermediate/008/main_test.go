package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestWordsSorting(t *testing.T) {
	inputWords := "without,hello,bag,world"
	expectedRes := "bag,hello,without,world"

	actualRes := wordsSorting(inputWords)

	assert.Equal(t, expectedRes, actualRes)
}
