package p

func P1(trunk bool, weights ...int) (n int64) {
	weight := 0
	for _, w := range weights {
		weight += w
	}
	if trunk {
		weight = weight / 4
	} else {
		weight = weight / 3
	}
	bestN, bestP := -1, int64(-1)
	var dist2 func(w int, w2 []int) bool
	dist2 = func(w int, w2 []int) bool {
		if len(w2) > 0 {
			w0 := w2[0]
			if w+w0 == weight {
				return true
			} else if w+w0 < weight {
				if dist2(w+w0, w2[1:]) {
					return true
				}
			}
			if dist2(w, w2[1:]) {
				return true
			}
		}
		return false
	}
	var dist1 func(w, n int, p int64, w1, w2 []int)
	dist1 = func(w, n int, p int64, w1, w2 []int) {
		if len(w1) > 0 {
			w0 := w1[0]
			if w+w0 == weight {
				if dist2(0, append(w2, w1[1:]...)) {
					n, p := n+1, p*int64(w0)
					if bestN < 0 || n < bestN || n == bestN && p < bestP {
						bestN, bestP = n, p
					}
				}
			} else if w+w0 < weight {
				if bestN < 0 || n+1 <= bestN {
					dist1(w+w0, n+1, p*int64(w0), w1[1:], w2)
				}
			}
			dist1(w, n, p, w1[1:], append(w2, w0))
		}
	}
	dist1(0, 0, 1, weights, make([]int, 0, len(weights)))
	n = bestP
	return
}
