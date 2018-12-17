package p

type Node struct {
	X, Y              int
	Size, Used, Avail int
}

func P1(nn []Node) (n int) {
	for i1, n1 := range nn {
		if n1.Used == 0 {
			continue
		}
		for i2, n2 := range nn {
			if i1 == i2 {
				continue
			}
			if n2.Avail < n1.Used {
				continue
			}
			n++
		}
	}
	return
}
