package p

func P1(input string) (output string) {
	p := []byte(input)
	for {
		for i := len(p) - 1; i >= 0; i-- {
			p[i]++
			if p[i] == 'i' || p[i] == 'o' || p[i] == 'l' {
				p[i]++
			}
			if p[i] > 'z' {
				p[i] = 'a'
				continue
			}
			break
		}
		{
			found := 0
			for i := 0; i < len(p)-2; i++ {
				if p[i]+1 != p[i+1] {
					continue
				}
				if p[i+1]+1 != p[i+2] {
					continue
				}
				found = 1
				break
			}
			if found == 0 {
				continue
			}
		}
		{
			found := 0
			for i := 0; i < len(p); i++ {
				if p[i] == 'i' || p[i] == 'o' || p[i] == 'l' {
					found = 1
					break
				}
			}
			if found == 1 {
				continue
			}
		}
		{
			found := 0
			for i := 0; i < len(p)-1; i++ {
				if p[i] == p[i+1] {
					found++
					i++
				}
				if found > 1 {
					break
				}
			}
			if found < 2 {
				continue
			}
		}
		return string(p)
	}
}
