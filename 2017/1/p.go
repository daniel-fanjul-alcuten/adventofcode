package p1

func P1(s string) (r int) {
	bb := []byte(s)
	for i, b := range bb {
		var c byte
		if i < len(bb)-1 {
			c = bb[i+1]
		} else {
			c = bb[0]
		}
		if b == c {
			r += int(b - '0')
		}
	}
	return
}

func P2(s string) (r int) {
	bb := []byte(s)
	for i, b := range bb {
		j := (i + len(bb)/2) % len(bb)
		c := bb[j]
		if b == c {
			r += int(b - '0')
		}
	}
	return
}
