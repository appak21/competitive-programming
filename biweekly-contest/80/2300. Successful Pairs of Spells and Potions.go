package contest

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	res := make([]int, n)
	sort.Ints(potions)
	for i := 0; i < n; i++ {
		div := (int(success) + spells[i] - 1) / spells[i]
		res[i] = m - sort.Search(len(potions), func(i int) bool {
			return potions[i] >= div
		})
	}
	return res
}

//the same, but without sort.Search
func successfulPairs1(spells []int, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	res := make([]int, n)
	sort.Ints(potions)
	for i := 0; i < n; i++ {
		div := success / int64(spells[i])
		if success%int64(spells[i]) != 0 {
			div++
		}
		idx := binary(potions, int(div))
		res[i] = m - idx
	}
	return res
}

func binary(a []int, t int) int {
	l, r := 0, len(a)
	for l < r {
		mid := l + (r-l)>>1
		if a[mid] < t {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r
}
