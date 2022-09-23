package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T int
	for fmt.Fscan(in, &T); T > 0; T-- {
		var k, j int
		fmt.Fscan(in, &k, &j)
		if j == 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		length := Length(k) - 1
		var l, dest int
		for l = 0; l <= 26; l++ {
			dest += Length(l - 1)
			letter := LetterNum(k, length, j+dest)
			if letter != l+1 {
				break
			}
		}
		fmt.Fprintln(out, Length(l)-1)
	}
}

func Length(k int) int {
	return int(math.Pow(2, float64(k)))
}

func LetterNum(k, n, i int) int {
	letter := k
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)>>1
		if mid == i {
			return letter
		}
		if mid > i {
			r = mid - 1
		} else {
			l = mid + 1
		}
		letter--
	}
	return letter
}

//abacaba d abacaba e abac[A]ba d abacaba F abacaba d [A]bacaba e abacaba d abacaba GGG abacaba d abacaba e abacaba d abacaba F abacaba d abacaba e abacaba d abacaba

/*
10
25 20
25 19
25 18
25 17
25 16
25 15
25 14
25 13
25 12
25 11
3
0
1
0
15
0
1
0
3
0
*/
