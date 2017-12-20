package p15

type Generator struct {
	Prev    int32
	Factor  int32
	Divisor int32
}

func (g *Generator) Next() (next int32) {
	next = int32(int64(g.Prev) * int64(g.Factor) % 2147483647)
	g.Prev = next
	return
}

func (g *Generator) Next2() (next int32) {
	for {
		next = int32(int64(g.Prev) * int64(g.Factor) % 2147483647)
		g.Prev = next
		if next%g.Divisor == 0 {
			break
		}
	}
	return
}

func Compare(a, b int32) (n int) {
	if a&0xffff == b&0xffff {
		n += 1
	}
	return
}
