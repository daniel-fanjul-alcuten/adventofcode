package p

func P1(min uint, input [][2]uint) (output uint) {
	output = min
	for {
		old := output
		for _, x := range input {
			if x[0] <= output && output <= x[1] {
				output = x[1] + 1
			}
		}
		if output == old {
			break
		}
	}
	return
}

func P2(min, max uint, input [][2]uint) (output uint) {
	rr, rr2 := [][2]uint{{min, max}}, [][2]uint{}
	for _, x := range input {
		for _, r := range rr {
			rr2 = append(rr2, split(r[0], r[1], x[0], x[1])...)
		}
		rr, rr2 = rr2, rr[:0]
	}
	for _, r := range rr {
		output += r[1] - r[0] + 1
	}
	return
}

func split(rB, rE, xB, xE uint) (s [][2]uint) {
	if xE < rB {
		//      rB-rE
		// xB-xE
		s = append(s, [2]uint{rB, rE})
	} else if xE < rE {
		if xB <= rB {
			//   rB--rE
			// xB--xE
		} else {
			// rB-----rE
			//   xB-xE
			s = append(s, [2]uint{rB, xB - 1})
		}
		s = append(s, [2]uint{xE + 1, rE})
	} else {
		if xB <= rB {
			//   rB-rE
			// xB-----xE
		} else if xB <= rE {
			// rB--rE
			//   xB--xE
			s = append(s, [2]uint{rB, xB - 1})
		} else {
			// rB-rE
			//      xB-xE
			s = append(s, [2]uint{rB, rE})
		}
	}
	return
}
