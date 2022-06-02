package contest

import "sort"

func minimumLines(stockPrices [][]int) int {
	l := len(stockPrices)
	lines := 1
	sort.Slice(stockPrices, func(i, j int) bool { //sort by x
		return stockPrices[i][0] < stockPrices[j][0]
	})
	for i := 1; i < l-1; i++ {
		x1, x2, x3 := stockPrices[i-1][0], stockPrices[i][0], stockPrices[i+1][0]
		y1, y2, y3 := stockPrices[i-1][1], stockPrices[i][1], stockPrices[i+1][1]
		a := (y3 - y2) * (x2 - x1)
		b := (y2 - y1) * (x3 - x2)
		if a != b {
			lines++
		}
	}
	return lines
}
