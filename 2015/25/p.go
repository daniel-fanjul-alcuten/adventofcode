package p

func P1(x, y int) (n int) {
	type P struct {
		x, y int
	}
	p, c := P{1, 1}, 20151125
	nextP := func() P {
		if p.y == 1 {
			return P{1, p.x + 1}
		}
		return P{p.x + 1, p.y - 1}
	}
	nextC := func() int {
		return c * 252533 % 33554393
	}
	for p != (P{x, y}) {
		p, c = nextP(), nextC()
	}
	return c
}
