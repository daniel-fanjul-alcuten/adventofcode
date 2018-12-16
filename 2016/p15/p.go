package p

type Disc struct {
	Time      int
	Position  int
	Positions int
}

func Check(input []Disc, time int) bool {
	for _, d := range input {
		time++
		pos := (d.Position + time - d.Time) % d.Positions
		if pos != 0 {
			return false
		}
	}
	return true
}

func P1(input []Disc) int {
	time := 0
	for {
		if Check(input, time) {
			return time
		}
		time++
	}
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func P3(input []Disc) int {
	time, lcm := 0, 1
	for i, d := range input {
		pos := (d.Position + i + time - d.Time) % d.Positions
		m := map[int]struct{}{}
		m[pos] = struct{}{}
		for pos != d.Positions-1 {
			if time+lcm < time {
				panic("overflow")
			}
			time += lcm
			pos = (pos + lcm) % d.Positions
			if _, ok := m[pos]; ok {
				return -1
			}
		}
		lcm = LCM(lcm, d.Positions)
	}
	return time
}
