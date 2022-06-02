package contest

import "strings"

func digitCount(num string) bool {
	var i byte
	for i = 0; i < byte(len(num)); i++ {
		tmp := int(num[i]) - 48
		if tmp > len(num) {
			return false
		}
		if strings.Count(num, string(i+48)) != tmp {
			return false
		}
	}
	return true
}
