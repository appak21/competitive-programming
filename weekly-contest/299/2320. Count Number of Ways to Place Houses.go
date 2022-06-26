package contest

func countHousePlacements(n int) int {
	f := make([]int, n+1)
	m := 1000000007
	f[0], f[1] = 1, 2
	for i := 2; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2]) % m
		//fmt.Println(f[i], f[i-1], f[i-2])
	}
	return f[n] * f[n] % m
}
