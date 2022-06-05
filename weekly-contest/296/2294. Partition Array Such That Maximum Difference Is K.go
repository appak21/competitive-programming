package contest

import "sort"

//Binary Search
func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	length := len(nums) - 1
	left, right := 0, length
	count, start := 0, 0
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid]-nums[start] > k {
			right = mid - 1
		} else if mid == length || nums[mid+1]-nums[start] > k {
			count++
			start = mid + 1
			left = mid + 1
			right = length
		} else {
			left = mid + 1
		}
	}
	return count
}

//easy solution, but it becomes slower when there are many numbers
func partitionArray1(nums []int, k int) int {
	sort.Ints(nums)
	count, start := 1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-nums[start] <= k {
			continue
		}
		count++
		start = i
	}
	return count
}
