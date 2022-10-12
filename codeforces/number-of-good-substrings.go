package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T int
	for fmt.Fscan(in, &T); T > 0; T-- {
		var s, t string
		fmt.Fscan(in, &s, &t)
		fmt.Fprintln(out, GoodSubstr(s, t))
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
//a - first element, d - difference
func SumOfArithmeticPr(a, n, d int) int {
	return n * (2*a + (n-1)*d) / 2
}
