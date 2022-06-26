package contest

func checkXMatrix(grid [][]int) bool {
	for i := range grid {
		for j := range grid[0] {
			if (grid[i][j] != 0) != (i == j || i+j == len(grid)-1) {
				return false
			}
		}
	}
	return true
}

func checkXMatrix1(grid [][]int) bool {
	k := len(grid[0]) - 1
	for i := 0; i < len(grid); i++ {
		if grid[i][i] == 0 {
			return false
		}
		if grid[i][k-i] == 0 {
			return false
		}
		for j := 0; j < len(grid[0]); j++ {
			if i != j && j != k-i && grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
