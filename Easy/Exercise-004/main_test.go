package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	inputStr := "34, 67, 55, 33, 12, 98"

	expectedRes := []int{34, 67, 55, 33, 12, 98}
	actualRes, err := concat(inputStr)

	require.NoError(t, err)

	assert.Equal(t, expectedRes, actualRes)
}
