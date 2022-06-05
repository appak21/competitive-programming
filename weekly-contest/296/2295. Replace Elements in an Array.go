package contest

func arrayChange(nums []int, operations [][]int) []int {
	m := make(map[int]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for _, op := range operations {
		nums[m[op[0]]] = op[1]
		m[op[1]] = m[op[0]]
		delete(m, op[0])
	}
	return nums
}
