//easy, 0 ms
package contest

import "strings"

func percentageLetter(s string, letter byte) int {
	return 100 * (len(strings.Split(s, string(letter))) - 1) / len(s)
}
