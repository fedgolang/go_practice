package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDevisable(t *testing.T) {
	lowNum := 0
	highNum := 100

	expectedRes := "7, 14, 21, 28, 42, 49, 56, 63, 77, 84, 91, 98"

	actualRes := fmt.Sprint(devisable(lowNum, highNum))

	assert.Equal(t, expectedRes, actualRes)
}
