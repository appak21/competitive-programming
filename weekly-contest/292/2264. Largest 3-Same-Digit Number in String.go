package contest

func largestGoodInteger(num string) string {
	digits := ""
	for i := 1; i < len(num)-1; i++ {
		if num[i-1] == num[i] && num[i] == num[i+1] {
			digits = max(digits, num[i-1:i+2])
		}
	}
	return digits
}
func max(a, b string) string {
	if a > b {
		return a
	}
	return b
}
