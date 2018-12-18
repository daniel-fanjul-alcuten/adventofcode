package p

func P1(input string) (output int) {
	type P struct {
		X, Y int
	}
	p := P{}
	m := map[P]struct{}{p: struct{}{}}
	for i := 0; i < len(input); i++ {
		switch in := input[i]; in {
		case '^':
			p.Y--
		case 'v':
			p.Y++
		case '<':
			p.X--
		case '>':
			p.X++
		}
		m[p] = struct{}{}
	}
	output = len(m)
	return
}

func P2(input string) (output int) {
	type P struct {
		X, Y int
	}
	p, q := P{}, P{}
	m := map[P]struct{}{p: struct{}{}}
	for i := 0; i < len(input); i++ {
		switch in := input[i]; in {
		case '^':
			p.Y--
		case 'v':
			p.Y++
		case '<':
			p.X--
		case '>':
			p.X++
		}
		m[p] = struct{}{}
		p, q = q, p
	}
	output = len(m)
	return
}
