package p

import (
	"fmt"
)

func Intcode(input int, code *[]int) (outputs []int, err error) {
	p := 0
	for {
		v1 := func() int {
			switch (*code)[p] / 100 % 10 {
			case 0:
				return (*code)[(*code)[p+1]]
			case 1:
				return (*code)[p+1]
			default:
				panic((*code)[p])
			}
		}
		v2 := func() int {
			switch (*code)[p] / 1000 % 10 {
			case 0:
				return (*code)[(*code)[p+2]]
			case 1:
				return (*code)[p+2]
			default:
				panic((*code)[p])
			}
		}
		switch (*code)[p] % 100 {
		case 1:
			(*code)[(*code)[p+3]] = v1() + v2()
			p += 4
		case 2:
			(*code)[(*code)[p+3]] = v1() * v2()
			p += 4
		case 3:
			(*code)[(*code)[p+1]] = input
			p += 2
		case 4:
			outputs = append(outputs, v1())
			p += 2
		case 5:
			if v1() != 0 {
				p = v2()
				break
			}
			p += 3
		case 6:
			if v1() == 0 {
				p = v2()
				break
			}
			p += 3
		case 7:
			if v1() < v2() {
				(*code)[(*code)[p+3]] = 1
			} else {
				(*code)[(*code)[p+3]] = 0
			}
			p += 4
		case 8:
			if v1() == v2() {
				(*code)[(*code)[p+3]] = 1
			} else {
				(*code)[(*code)[p+3]] = 0
			}
			p += 4
		case 99:
			return
		default:
			err = fmt.Errorf("Unexpected opcode %v", (*code)[p])
			return
		}
	}
}
