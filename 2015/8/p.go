package p

func P1(input []string) (n1, n2 int) {
	for _, in := range input {
		n1 += len(in)
		inside := false
		for i := 0; i < len(in); i++ {
			b := in[i]
			if !inside {
				if b != '"' {
					panic(b)
				}
				inside = true
				continue
			}
			if b == '"' {
				inside = false
				continue
			}
			if b == '\\' {
				i++
				b = in[i]
				if b == '\\' {
					n2++
					continue
				}
				if b == '"' {
					n2++
					continue
				}
				if b == 'x' {
					i += 2
					n2++
					continue
				}
			}
			n2++
			continue
		}
	}
	return
}

func P2(input []string) (n1, n2 int) {
	for _, in := range input {
		n1 += 2
		for i := 0; i < len(in); i++ {
			b := in[i]
			if b == '\\' {
				n1 += 2
				continue
			}
			if b == '"' {
				n1 += 2
				continue
			}
			n1++
			continue
		}
		n2 += len(in)
	}

	return
}
