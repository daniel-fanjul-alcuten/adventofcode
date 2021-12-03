package p

func P1(input ...S) (f, d int) {
	for _, i := range input {
		switch i.d {
		case "forward":
			f += i.q
		case "down":
			d += i.q
		case "up":
			d -= i.q
		default:
			panic(i.d)
		}
	}
	return
}

func P2(input ...S) (f, d int) {
	aim := 0
	for _, i := range input {
		switch i.d {
		case "forward":
			f += i.q
			d += aim * i.q
		case "down":
			aim += i.q
		case "up":
			aim -= i.q
		default:
			panic(i.d)
		}
	}
	return
}
