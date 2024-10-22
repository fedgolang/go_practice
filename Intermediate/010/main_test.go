package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDistinctSortLine(t *testing.T) {
	inputLine := "hello world and practice makes perfect and hello world again"
	expectedRes := "again makes perfect practice"
	actualRes := distinctSortLine(inputLine)

	assert.Equal(t, expectedRes, actualRes)
}
