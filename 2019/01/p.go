package p

func Fuel(mass ...int) (fuel int) {
	for _, m := range mass {
		fuel += m/3 - 2
	}
	return
}

func Fuel2(mass ...int) (fuel int) {
	for _, m := range mass {
		for {
			f := m/3 - 2
			if f <= 0 {
				break
			}
			fuel, m = fuel+f, f
		}
	}
	return
}
