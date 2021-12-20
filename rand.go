package log4shell

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randBool() bool {
	return rand.Int63()%2 == 0 // #nosec
}

func randString(n int) string {
	str := make([]rune, n)
	for i := 0; i < n; i++ {
		s := ' ' + 1 + rand.Intn(90) // #nosec
		switch {
		case s >= '0' && s <= '9':
		case s >= 'A' && s <= 'Z':
		case s >= 'a' && s <= 'z':
		case isValidSymbol(s):
		default:
			i--
			continue
		}
		str[i] = rune(s)
	}
	return string(str)
}

func isValidSymbol(s int) bool {
	switch s {
	case '(', ')':
	case '*', '.':
	case '$', '_':
	case '[', ']':
	case '@', '=':
	default:
		return false
	}
	return true
}

func randSecret() string {
	const n = 8

	str := make([]rune, n)
	for i := 0; i < n; i++ {
		s := ' ' + 1 + rand.Intn(90) // #nosec
		switch {
		case s >= '0' && s <= '9':
		case s >= 'A' && s <= 'Z':
		case s >= 'a' && s <= 'z':
		default:
			i--
			continue
		}
		str[i] = rune(s)
	}
	return string(str)
}
