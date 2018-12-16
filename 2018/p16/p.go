package p

type Registers [4]int

type Instruction [4]int

type Operation func(a, b, c int, r Registers) Registers

func addr(a, b, c int, r Registers) Registers {
	r[c] = r[a] + r[b]
	return r
}

func addi(a, b, c int, r Registers) Registers {
	r[c] = r[a] + b
	return r
}

func mulr(a, b, c int, r Registers) Registers {
	r[c] = r[a] * r[b]
	return r
}

func muli(a, b, c int, r Registers) Registers {
	r[c] = r[a] * b
	return r
}

func banr(a, b, c int, r Registers) Registers {
	r[c] = r[a] & r[b]
	return r
}

func bani(a, b, c int, r Registers) Registers {
	r[c] = r[a] & b
	return r
}

func borr(a, b, c int, r Registers) Registers {
	r[c] = r[a] | r[b]
	return r
}

func bori(a, b, c int, r Registers) Registers {
	r[c] = r[a] | b
	return r
}

func setr(a, b, c int, r Registers) Registers {
	r[c] = r[a]
	return r
}

func seti(a, b, c int, r Registers) Registers {
	r[c] = a
	return r
}

func gtir(a, b, c int, r Registers) Registers {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func gtri(a, b, c int, r Registers) Registers {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func gtrr(a, b, c int, r Registers) Registers {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqir(a, b, c int, r Registers) Registers {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqri(a, b, c int, r Registers) Registers {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqrr(a, b, c int, r Registers) Registers {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

var Operations = map[string]Operation{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

func P1(input [][4]int) (n int) {
	for i := 0; i < len(input)-2; i += 3 {
		bef := Registers(input[i])
		ins := Instruction(input[i+1])
		a, b, c := ins[1], ins[2], ins[3]
		aft := Registers(input[i+2])
		ni := 0
		for _, op := range Operations {
			if op(a, b, c, bef) == aft {
				ni++
			}
		}
		if ni >= 3 {
			n++
		}
	}
	return
}

func P2(input [][4]int, program []Instruction) (r0 int) {
	m := map[int]map[string]struct{}{}
	for i := 0; i < len(Operations); i++ {
		m[i] = map[string]struct{}{}
		for n := range Operations {
			m[i][n] = struct{}{}
		}
	}
	for i := 0; i < len(input)-2; i += 3 {
		bef := Registers(input[i])
		ins := Instruction(input[i+1])
		a, b, c := ins[1], ins[2], ins[3]
		aft := Registers(input[i+2])
		for n, op := range Operations {
			if op(a, b, c, bef) != aft {
				delete(m[ins[0]], n)
			}
		}
	}
	table := map[int]Operation{}
	for l := 0; l < len(Operations); l++ {
		for i, m2 := range m {
			if len(m2) == 1 {
				for n := range m2 {
					table[i] = Operations[n]
					delete(m, i)
					for _, m3 := range m {
						delete(m3, n)
					}
				}
			}
		}
	}
	if len(table) != len(Operations) {
		panic(len(table))
	}
	r := Registers{}
	for _, p := range program {
		r = table[p[0]](p[1], p[2], p[3], r)
	}
	r0 = r[0]
	return
}
