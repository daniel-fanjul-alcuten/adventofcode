package p

func Intcode(inputs, code []int) (outputs []int) {
	p, r, s := 0, 0, make(map[int]int, len(code))
	for i, c := range code {
		s[i] = c
	}
	for {
		v1 := func() int {
			switch s[p] / 100 % 10 {
			case 0:
				return s[s[p+1]]
			case 1:
				return s[p+1]
			case 2:
				return s[r+s[p+1]]
			default:
				panic(s[p] / 100 % 10)
			}
		}
		v2 := func() int {
			switch s[p] / 1000 % 10 {
			case 0:
				return s[s[p+2]]
			case 1:
				return s[p+2]
			case 2:
				return s[r+s[p+2]]
			default:
				panic(s[p] / 1000 % 10)
			}
		}
		a1 := func(v int) {
			switch s[p] / 100 % 10 {
			case 0:
				s[s[p+1]] = v
			case 2:
				s[r+s[p+1]] = v
			default:
				panic(s[p] / 100 % 10)
			}
		}
		a3 := func(v int) {
			switch s[p] / 10000 % 10 {
			case 0:
				s[s[p+3]] = v
			case 2:
				s[r+s[p+3]] = v
			default:
				panic(s[p] / 1000 % 10)
			}
		}
		switch s[p] % 100 {
		case 1:
			a3(v1() + v2())
			p += 4
		case 2:
			a3(v1() * v2())
			p += 4
		case 3:
			if len(inputs) == 0 {
				return
			}
			a1(inputs[0])
			inputs = inputs[1:]
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
				a3(1)
			} else {
				a3(0)
			}
			p += 4
		case 8:
			if v1() == v2() {
				a3(1)
			} else {
				a3(0)
			}
			p += 4
		case 9:
			r += v1()
			p += 2
		case 99:
			return
		default:
			panic(s[p] % 100)
		}
	}
}
