//easy, 0 ms
package contest

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
#hex
res := 0x3f3f3f3f // 3f3f3f3f = 1061109567
– what is 0x  in programming ?
– The prefix 0x is used in code to indicate that the number is being written in hex.
Sometimes you need to set a large number. We can use math.MaxInt, however, we get overflow if we add the large number by 1.
Programmers like to use the hex number 3f3f3f3f as INF.
Reasons for that:
0x3f3f3f3f is large enough in most cases
0x3f3f3f3f + 0x3f3f3f3f will not overflow
*/
