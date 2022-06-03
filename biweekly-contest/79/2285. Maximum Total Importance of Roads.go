package contest

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	h := make([]int, n)
	for _, r := range roads {
		h[r[0]]++
		h[r[1]]++
	}
	sort.Ints(h)
	var ans int64
	for i := 0; i < n; i++ {
		ans += int64(h[i] * (i + 1))
	}
	return ans
}
