package p

func P1(input ...int) (r int) {
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			r++
		}
	}
	return
}

func P2(input ...int) (r int) {
	for i := 3; i < len(input); i++ {
		if input[i] > input[i-3] {
			r++
		}
	}
	return
}
