package p

import "fmt"

func Intcode(code *[]int) (err error) {
	p := 0
	for {
		switch (*code)[p] {
		case 1:
			(*code)[(*code)[p+3]] = (*code)[(*code)[p+1]] + (*code)[(*code)[p+2]]
		case 2:
			(*code)[(*code)[p+3]] = (*code)[(*code)[p+1]] * (*code)[(*code)[p+2]]
		case 99:
			return
		default:
			err = fmt.Errorf("Unexpected opcode %v", (*code)[p])
			return
		}
		p += 4
	}
}

func Search(code []int) (noun, verb int, err error) {
	s := make([]int, len(code))
	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			s = append(s[:0], code...)
			s[1], s[2] = noun, verb
			if err := Intcode(&s); err != nil {
				continue
			}
			if v := s[0]; v == 19690720 {
				return
			}
		}
	}
	err = fmt.Errorf("Not Found")
	return
}
