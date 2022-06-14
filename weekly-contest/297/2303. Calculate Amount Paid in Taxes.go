package contest

func calculateTax(brackets [][]int, income int) float64 {
	var total float64
	for i := 0; i < len(brackets); i++ {
		if i == 0 {
			tmp := min(income, brackets[i][0])
			total += float64(tmp*brackets[i][1]) / 100
			income -= tmp
		} else {
			tmp := min(brackets[i][0]-brackets[i-1][0], income)
			total += float64(tmp*brackets[i][1]) / 100
			income -= tmp
		}
	}
	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
