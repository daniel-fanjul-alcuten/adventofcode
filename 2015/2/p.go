package p

func P1(input [][3]int) (output int) {
	for _, in := range input {
		l, w, h := in[0], in[1], in[2]
		output += 2*l*w + 2*w*h + 2*h*l
		min := l * w
		if w*h < min {
			min = w * h
		}
		if h*l < min {
			min = h * l
		}
		output += min
	}
	return
}

func P2(input [][3]int) (output int) {
	for _, in := range input {
		l, w, h := in[0], in[1], in[2]
		p := 2*l + 2*w
		if 2*l+2*h < p {
			p = 2*l + 2*h
		}
		if 2*w+2*h < p {
			p = 2*w + 2*h
		}
		output += l*w*h + p
	}
	return
}
