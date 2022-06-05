package contest

func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums) >> 1
	newNums := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			newNums[i] = min(nums[2*i], nums[2*i+1])
		} else {
			newNums[i] = max(nums[2*i], nums[2*i+1])
		}
	}
	return minMaxGame(newNums)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//without recursion
func minMaxGame2(nums []int) int {
rep:
	newNums := make([]int, len(nums)/2)
	n := len(nums)
	for i := 0; i < n/2; i += 2 {
		newNums[i] = min(nums[2*i], nums[2*i+1])
		if i != n/2-1 {
			newNums[i+1] = max(nums[2*(i+1)], nums[2*(i+1)+1])
		}
	}
	if len(newNums) > 0 {
		nums = newNums
		goto rep
	}
	return nums[len(nums)-1]
}
