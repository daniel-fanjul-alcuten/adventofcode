package p

import (
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	LowTarget  string
	LowID      int
	HighTarget string
	HighID     int
}

func P12(input []string, val1, val2 int, sto1, sto2, sto3 int) (p1 int, p2 int) {
	botToVals := map[int][]int{}
	stoToVals := map[int][]int{}
	rules := map[int]Rule{}
	for _, in := range input {
		split := strings.Split(in, " ")
		if split[0] == "value" {
			value, bot := atoi(split[1]), atoi(split[5])
			botToVals[bot] = append(botToVals[bot], value)
		} else if split[0] == "bot" {
			bot := atoi(split[1])
			lowTarget, lowID := split[5], atoi(split[6])
			highTarget, highID := split[10], atoi(split[11])
			rules[bot] = Rule{lowTarget, lowID, highTarget, highID}
		}
	}
	for len(rules) > 0 {
		for b, r := range rules {
			values := botToVals[b]
			if len(values) > 1 {
				sort.Ints(values)
				l, h := values[0], values[1]
				if r.LowTarget == "bot" {
					botToVals[r.LowID] = append(botToVals[r.LowID], l)
				} else {
					stoToVals[r.LowID] = append(stoToVals[r.LowID], l)
				}
				if r.HighTarget == "bot" {
					botToVals[r.HighID] = append(botToVals[r.HighID], h)
				} else {
					stoToVals[r.HighID] = append(stoToVals[r.HighID], h)
				}
				delete(rules, b)
			}
		}
	}
	for b, vv := range botToVals {
		ok1, ok2 := false, false
		for _, v := range vv {
			if v == val1 {
				ok1 = true
			}
			if v == val2 {
				ok2 = true
			}
		}
		if ok1 && ok2 {
			p1 = b
			break
		}
	}
	p2 = stoToVals[sto1][0] * stoToVals[sto2][0] * stoToVals[sto3][0]
	return
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
