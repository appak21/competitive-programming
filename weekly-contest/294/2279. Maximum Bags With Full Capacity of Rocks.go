package contest

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	moreRocks, sum := make([]int, len(capacity)), 0
	for i := 0; i < len(capacity); i++ {
		moreRocks[i] = capacity[i] - rocks[i]
		sum += moreRocks[i]
	}
	if sum < additionalRocks {
		return len(rocks)
	}
	sort.Ints(moreRocks)
	count := 0
	for i := 0; i < len(moreRocks) && moreRocks[i] <= additionalRocks; i++ {
		additionalRocks -= moreRocks[i]
		count++
	}
	return count
}
