package contest

func countAsterisks(s string) int {
	count, between := 0, false
	for _, v := range s {
		if v == '|' {
			between = !between
		} else if v == '*' && !between {
			count++
		}
	}
	return count
}

func countAsterisks1(s string) int {
	count, between := 0, 0
	for _, v := range s {
		if v == '|' {
			between ^= 1
		} else if v == '*' && between == 0 {
			count++
		}
	}
	return count
}
