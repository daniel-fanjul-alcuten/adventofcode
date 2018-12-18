package p8

import (
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	Variable     string
	Operation    string
	Value        int
	CondVariable string
	CondOperator string
	CondValue    int
}

func Parse1(s string) (k Instruction) {
	i := strings.Index(s, " ")
	if i < 0 {
		panic(s)
	}
	k.Variable, s = s[:i], s[i+1:]
	if len(s) < 3 {
		panic(s)
	}
	k.Operation, s = s[:3], s[4:]
	i = strings.Index(s, " ")
	if i < 0 {
		panic(s)
	}
	a, err := strconv.Atoi(s[:i])
	if err != nil {
		panic(s)
	}
	k.Value, s = a, s[i+1:]
	if len(s) < 3 {
		panic(s)
	}
	s = s[3:]
	i = strings.Index(s, " ")
	if i < 0 {
		panic(s)
	}
	k.CondVariable, s = s[:i], s[i+1:]
	i = strings.Index(s, " ")
	if i < 0 {
		panic(s)
	}
	k.CondOperator, s = s[:i], s[i+1:]
	a, err = strconv.Atoi(s)
	if err != nil {
		panic(s)
	}
	k.CondValue, s = a, s[len(s):]
	return
}

func ParseN(ss []string) (kk []Instruction) {
	for _, s := range ss {
		kk = append(kk, Parse1(s))
	}
	return
}

type Records map[string]int

func (rr Records) Apply1(k Instruction) {
	value1, value2 := rr[k.CondVariable], k.CondValue
	switch k.CondOperator {
	case "<":
		if value1 >= value2 {
			return
		}
	case ">":
		if value1 <= value2 {
			return
		}
	case "<=":
		if value1 > value2 {
			return
		}
	case ">=":
		if value1 < value2 {
			return
		}
	case "==":
		if value1 != value2 {
			return
		}
	case "!=":
		if value1 == value2 {
			return
		}
	default:
		panic(k.CondOperator)
	}
	switch k.Operation {
	case "inc":
		rr[k.Variable] += k.Value
	case "dec":
		rr[k.Variable] -= k.Value
	default:
		panic(k.CondOperator)
	}
}

func (rr Records) ApplyN(kk []Instruction) {
	for _, k := range kk {
		rr.Apply1(k)
	}
}

func (rr Records) Max() (max int) {
	max = math.MinInt32
	for _, v := range rr {
		if v > max {
			max = v
		}
	}
	return
}
