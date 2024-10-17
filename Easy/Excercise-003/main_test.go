package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestMapSqr(t *testing.T) {
	inputNum := 8

	expectedRes := map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25, 6: 36, 7: 49, 8: 64}
	actualRes, err := mapSqr(inputNum)

	require.NoError(t, err)

	assert.Equal(t, expectedRes, actualRes)
}
