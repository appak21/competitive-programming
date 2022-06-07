package contest

import "math"

const mod, mx int = 1e9 + 7, 1e5

var f = [mx + 1]int{1, 1, 2, 4}
var g = f

func init() {
	for i := 4; i <= mx; i++ {
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
		g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	}
}

func countTexts(pressedKeys string) int {
	ans, cnt := 1, 0
	for i, s := range pressedKeys {
		cnt++
		if i == len(pressedKeys)-1 || pressedKeys[i+1] != byte(s) {
			if s != '7' && s != '9' {
				ans = ans * f[cnt] % mod
			} else {
				ans = ans * g[cnt] % mod
			}
			cnt = 0
		}
	}
	return ans
}

//DP
func countTexts1(pressedKeys string) int {
	mod := int(math.Pow(10, 9)) + 7
	dp := make([]int, len(pressedKeys)+1)
	dp[0] = 1
	for i := 1; i < len(pressedKeys)+1; i++ {
		dp[i] = dp[i-1] % mod
		if i > 1 && pressedKeys[i-1] == pressedKeys[i-2] {
			dp[i] = (dp[i] + dp[i-2]) % mod
			if i > 2 && pressedKeys[i-1] == pressedKeys[i-3] {
				dp[i] = (dp[i] + dp[i-3]) % mod
				if i > 3 && (pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9') && pressedKeys[i-1] == pressedKeys[i-4] {
					dp[i] = (dp[i] + dp[i-4]) % mod
				}
			}

		}
	}
	return dp[len(pressedKeys)]
}
