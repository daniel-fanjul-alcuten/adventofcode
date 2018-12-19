package p

type Instruction struct {
	n string
	p [3]int
}

type Program struct {
	b int
	i []Instruction
}

type Registers [6]int

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

func P1(input Program, r0 int) int {
	ip, r := 0, Registers{}
	r[0] = r0
	for ip >= 0 && ip < len(input.i) {
		i := input.i[ip]
		op, a, b, c := Operations[i.n], i.p[0], i.p[1], i.p[2]
		r[input.b] = ip
		r = op(a, b, c, r)
		ip = r[input.b] + 1
	}
	return r[0]
}
