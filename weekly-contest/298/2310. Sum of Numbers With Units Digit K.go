package contest

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 10; i++ {
		if i*k%10 == num%10 && i*k <= num {
			return i
		}
	}
	return -1
}
