package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestGetInteger(t *testing.T) {
	require.Equal(t, 1, getInteger(1))
}
