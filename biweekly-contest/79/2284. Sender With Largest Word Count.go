package contest

import (
	"sort"
	"strings"
)

//Optimized solution
func largestWordCount(messages []string, senders []string) string {
	msgCount := make(map[string]int)
	for i, sender := range senders {
		count := 1
		for _, ch := range messages[i] {
			if ch == ' ' {
				count++
			}
		}
		msgCount[sender] += count
	}
	var (
		max = 0
		ans string
	)
	for sender, words := range msgCount {
		switch {
		case max > words:
		case max < words:
			max = words
			fallthrough
		case ans < sender:
			ans = sender
		}
	}
	return ans
}

//First Submission
func largestWordCount2(messages []string, senders []string) string {
	m := make(map[string]int)
	for i, sender := range senders {
		sp := strings.Split(messages[i], " ")
		m[sender] += len(sp)
	}

	mapConv := converter(m)
	sort.Slice(mapConv, func(i, j int) bool {
		return mapConv[i].wordCount > mapConv[j].wordCount
	})
	name := mapConv[0].name
	for i := 0; i < len(mapConv)-1; i++ {
		if mapConv[i].wordCount == mapConv[i+1].wordCount {
			a := strings.Compare(name, mapConv[i+1].name)
			if a < 0 {
				name = mapConv[i+1].name
			}
		} else {
			break
		}
	}
	return name
}

type Msg struct {
	name      string
	wordCount int
}

func converter(m map[string]int) []Msg {
	res := make([]Msg, 0, len(m))
	for i, val := range m {
		res = append(res, Msg{name: i, wordCount: val})
	}
	return res
}
