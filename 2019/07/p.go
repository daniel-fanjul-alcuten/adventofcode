package p

import (
	"math"
)

func Intcode(inputs, code []int) (outputs []int) {
	p, s := 0, append(make([]int, 0, len(code)), code...)
	for {
		v1 := func() int {
			switch s[p] / 100 % 10 {
			case 0:
				return s[s[p+1]]
			case 1:
				return s[p+1]
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
			default:
				panic(s[p] / 1000 % 10)
			}
		}
		a1 := func(v int) {
			s[s[p+1]] = v
		}
		a3 := func(v int) {
			s[s[p+3]] = v
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
		case 99:
			return
		default:
			panic(s[p] % 100)
		}
	}
}

func P1(code []int) (n int) {
	n = math.MinInt64
	inputs := []int{0, 1, 2, 3, 4}
	for _, i0 := range inputs {
		outputs := Intcode([]int{i0, 0}, code)
		output := outputs[0]
		for _, i1 := range inputs {
			if i1 == i0 {
				continue
			}
			outputs = Intcode([]int{i1, output}, code)
			output := outputs[0]
			for _, i2 := range inputs {
				if i2 == i0 {
					continue
				}
				if i2 == i1 {
					continue
				}
				outputs = Intcode([]int{i2, output}, code)
				output := outputs[0]
				for _, i3 := range inputs {
					if i3 == i0 {
						continue
					}
					if i3 == i1 {
						continue
					}
					if i3 == i2 {
						continue
					}
					outputs = Intcode([]int{i3, output}, code)
					output := outputs[0]
					for _, i4 := range inputs {
						if i4 == i0 {
							continue
						}
						if i4 == i1 {
							continue
						}
						if i4 == i2 {
							continue
						}
						if i4 == i3 {
							continue
						}
						outputs = Intcode([]int{i4, output}, code)
						if output := outputs[0]; output > n {
							n = output
						}
					}
				}
			}
		}
	}
	return
}

func P2(code []int) (n int) {
	n = math.MinInt64
	inputs := []int{5, 6, 7, 8, 9}
	for _, i0 := range inputs {
		inputs0 := []int{i0, 0}
		outputs0 := Intcode(inputs0, code)
		for _, i1 := range inputs {
			if i1 == i0 {
				continue
			}
			inputs1 := []int{i1, outputs0[len(outputs0)-1]}
			outputs1 := Intcode(inputs1, code)
			for _, i2 := range inputs {
				if i2 == i0 {
					continue
				}
				if i2 == i1 {
					continue
				}
				inputs2 := []int{i2, outputs1[len(outputs1)-1]}
				outputs2 := Intcode(inputs2, code)
				for _, i3 := range inputs {
					if i3 == i0 {
						continue
					}
					if i3 == i1 {
						continue
					}
					if i3 == i2 {
						continue
					}
					inputs3 := []int{i3, outputs2[len(outputs2)-1]}
					outputs3 := Intcode(inputs3, code)
					for _, i4 := range inputs {
						if i4 == i0 {
							continue
						}
						if i4 == i1 {
							continue
						}
						if i4 == i2 {
							continue
						}
						if i4 == i3 {
							continue
						}
						inputs4 := []int{i4, outputs3[len(outputs3)-1]}
						outputs4 := Intcode(inputs4, code)
						if output := outputs4[len(outputs4)-1]; output > n {
							n = output
						}
						inputs0 := inputs0
						inputs1 := inputs1
						inputs2 := inputs2
						inputs3 := inputs3
						outputs0 := outputs0
						outputs1 := outputs1
						outputs2 := outputs2
						outputs3 := outputs3
						for {
							inputs0 = append(inputs0, outputs4[len(outputs4)-1])
							outputs0 = Intcode(inputs0, code)
							if len(outputs0) < len(inputs0)-1 {
								break
							}
							inputs1 = append(inputs1, outputs0[len(outputs0)-1])
							outputs1 = Intcode(inputs1, code)
							if len(outputs1) < len(inputs1)-1 {
								break
							}
							inputs2 = append(inputs2, outputs1[len(outputs1)-1])
							outputs2 = Intcode(inputs2, code)
							if len(outputs2) < len(inputs2)-1 {
								break
							}
							inputs3 = append(inputs3, outputs2[len(outputs2)-1])
							outputs3 = Intcode(inputs3, code)
							if len(outputs3) < len(inputs3)-1 {
								break
							}
							inputs4 = append(inputs4, outputs3[len(outputs3)-1])
							outputs4 = Intcode(inputs4, code)
							if output := outputs4[len(outputs4)-1]; output > n {
								n = output
							}
							if len(outputs4) < len(inputs4)-1 {
								break
							}
						}
					}
				}
			}
		}
	}
	return
}
