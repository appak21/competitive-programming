package contest

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	sp := strings.Split(sentence, " ")
	for i, word := range sp {
		n := getPrice(word)
		if n < 0 {
			continue
		}
		sp[i] = fmt.Sprintf("$%.2f", float64(n*(100-discount))/float64(100))
	}
	return strings.Join(sp, " ")
}

func getPrice(s string) int {
	if !strings.HasPrefix(s, "$") || len(s) <= 1 {
		return -1
	}
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return -1
	}
	return n
}
