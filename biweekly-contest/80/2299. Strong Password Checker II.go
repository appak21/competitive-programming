package contest

import "strings"

func strongPasswordCheckerII(password string) bool {
	l := len(password)
	var lower, upper, digit, spec bool
	if l < 8 {
		return false
	}
	for i := 0; i < l; i++ {
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		if strings.Contains("!@#$%^&*()-+", string(password[i])) {
			spec = true
		} else if password[i] >= 97 && password[i] <= 122 {
			lower = true
		} else if password[i] >= 65 && password[i] <= 90 {
			upper = true
		} else if password[i] >= 48 && password[i] <= 57 {
			digit = true
		}
	}
	return lower && upper && digit && spec
}
