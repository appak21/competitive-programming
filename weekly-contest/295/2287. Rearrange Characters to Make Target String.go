package contest

import "strings"

func rearrangeCharacters(s string, target string) int {
	sch, tch := make([]int, 26), make([]int, 26)
	for _, ch := range s {
		sch[ch-'a']++
	}
	for _, ch := range target {
		tch[ch-'a']++
	}
	res := 0x3f3f3f3f
	for i := 0; i < 26; i++ {
		if tch[i] > 0 {
			res = min(res, sch[i]/tch[i])
		}
	}
	return res
}

func rearrangeCharacters2(s string, target string) int {
	res := len(s)
	for _, ch := range target {
		count := strings.Count(s, string(ch))
		tcount := strings.Count(target, string(ch))
		if tcount > count {
			return 0
		}
		res = min(res, count/tcount)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
