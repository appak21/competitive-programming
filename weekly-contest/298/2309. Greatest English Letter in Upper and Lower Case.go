package contest

import "strings"

func greatestLetter(s string) string {
	for c := 'Z'; c >= 'A'; c-- {
		if strings.Contains(s, string(c)) && strings.Contains(s, string(c+32)) {
			return string(c)
		}
	}
	return ""
}
