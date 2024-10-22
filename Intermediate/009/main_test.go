package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLinesToUpper(t *testing.T) {
	inputLine := `some
	lines
	to
	upper
	`
	expectedRes := `SOME
	LINES
	TO
	UPPER
	`

	actualRes := linesToUpper(inputLine)

	assert.Equal(t, expectedRes, actualRes)
}
