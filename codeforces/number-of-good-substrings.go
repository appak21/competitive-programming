package main

import (
	"fmt"
	"strings"
)

func main() {
	var l int
	fmt.Scanf("%d\n", &l)
	for i := 0; i < l; i++ {
		var s, t string
		fmt.Scanf("%s\n", &s)
		fmt.Scanf("%s\n", &t)
		fmt.Println(GoodSubstr(s, t))
	}
}

func GoodSubstr(s, t string) int {
	n := len(s)
	idx, bad := 0, 0
	all := SumOfArithmeticPr(1, n, 1) //number of all substrings including bad and good substrings
	for idx != -1 {
		idx = strings.Index(s, t)
		bad += (n - len(t) + 1 - idx) * (idx + 1) //number of bad substrings
		s = s[idx+1:]
		n = n - idx - 1
	}
	return all - bad //all-bad=good
}

//SumOfArithmeticPr finds S(n) (sum of first n terms) of arithmetic progression
//if a - first element, d - difference
func SumOfArithmeticPr(a, n, d int) int {
	return n * (2*a + (n-1)*d) / 2
}
