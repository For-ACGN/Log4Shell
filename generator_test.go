package log4shell

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestGenerateExecuteClass(t *testing.T) {
	template, err := os.ReadFile("testdata/template/Exec.class")
	require.NoError(t, err)
	spew.Dump(template)

	t.Run("common", func(t *testing.T) {
		class, err := GenerateExecuteClass(template, "whoami", "Test")
		require.NoError(t, err)
		spew.Dump(class)
	})

	t.Run("default class", func(t *testing.T) {
		class, err := GenerateExecuteClass(template, "${cmd}", "")
		require.NoError(t, err)
		spew.Dump(class)
		require.Equal(t, template, class)
	})

	t.Run("compare with Calc", func(t *testing.T) {
		class, err := GenerateExecuteClass(template, "calc", "Calc")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/Calc.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("compare with Notepad", func(t *testing.T) {
		class, err := GenerateExecuteClass(template, "notepad", "Notepad")
		require.NoError(t, err)
		spew.Dump(class)

		expected, err := os.ReadFile("testdata/template/Notepad.class")
		require.NoError(t, err)
		require.Equal(t, expected, class)
	})

	t.Run("empty command", func(t *testing.T) {
		class, err := GenerateExecuteClass(template, "", "Test")
		require.EqualError(t, err, "empty command")
		require.Zero(t, class)
	})
}
