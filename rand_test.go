package log4shell

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandString(t *testing.T) {
	for i := 0; i < 10000; i++ {
		str := randString(64)
		require.False(t, strings.Contains(str, " "))
	}
}

func TestRandSecret(t *testing.T) {
	for i := 0; i < 10000; i++ {
		str := randSecret()
		require.False(t, strings.Contains(str, " "))
	}
}
