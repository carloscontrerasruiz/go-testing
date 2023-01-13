package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccess(t *testing.T) {
	assert := require.New(t)

	result := Add(20, 2)

	expect := 22

	assert.Equal(expect, result)
}
