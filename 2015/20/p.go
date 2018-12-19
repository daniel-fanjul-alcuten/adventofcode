package p

func P1(presents int, p2 bool) int {
	house, f, mm, primes := 1, []int{}, map[int]struct{}{}, []int{}
	limit := map[int]int{}
	for {
		h := house
		f = f[:0]
		for _, p := range primes {
			for h%p == 0 {
				f = append(f, p)
				h = h / p
			}
			if p*p > h {
				break
			}
		}
		if p := h; p != 1 {
			primes = append(primes, p)
			f = append(f, p)
			h = h / p
		}
		var mul func(i int, m int)
		mul = func(i int, m int) {
			if i >= len(f) {
				mm[m] = struct{}{}
				return
			}
			mul(i+1, m)
			mul(i+1, m*f[i])
		}
		mul(0, 1)
		pr := 0
		for m := range mm {
			if p2 {
				if limit[m] < 50 {
					pr += 11 * m
					limit[m]++
				}
			} else {
				pr += 10 * m
			}
			delete(mm, m)
		}
		if pr >= presents {
			return house
		}
		house++
	}
}
