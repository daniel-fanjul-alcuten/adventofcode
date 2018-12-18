package p

import (
	"encoding/json"
	"fmt"
)

func P1(input string) (output int) {
	var v interface{}
	if err := json.Unmarshal([]byte(input), &v); err != nil {
		panic(err)
	}
	var sum func(v interface{}) int
	sum = func(v interface{}) (s int) {
		switch w := v.(type) {
		case string:
		case float64:
			s += int(w)
		case []interface{}:
			for _, v := range w {
				s += sum(v)
			}
		case map[string]interface{}:
			for _, v := range w {
				s += sum(v)
			}
		default:
			panic(fmt.Sprintf("%T", w))
		}
		return
	}
	output = sum(v)
	return
}

func P2(input string) (output int) {
	var v interface{}
	if err := json.Unmarshal([]byte(input), &v); err != nil {
		panic(err)
	}
	var sum func(v interface{}) int
	sum = func(v interface{}) (s int) {
		switch w := v.(type) {
		case string:
		case float64:
			s += int(w)
		case []interface{}:
			for _, v := range w {
				s += sum(v)
			}
		case map[string]interface{}:
			for _, v := range w {
				if v == "red" {
					return
				}
			}
			for _, v := range w {
				s += sum(v)
			}
		default:
			panic(fmt.Sprintf("%T", w))
		}
		return
	}
	output = sum(v)
	return
}
