package p

type Item struct {
	N       string
	C, D, A int
}

type Character struct {
	HP   int
	D, A int
}

func P1(p, e Character, w, a, r []Item, p2 bool) (gold int) {
	gold = -1
	c := 0
	for _, k := range w {
		c := c + k.C
		p := p
		p.D += k.D
		p.A += k.A
		for _, k := range a {
			c := c + k.C
			p := p
			p.D += k.D
			p.A += k.A
			for i, k := range r {
				c := c + k.C
				p := p
				p.D += k.D
				p.A += k.A
				for j, k := range r {
					if i == j {
						continue
					}
					c := c + k.C
					p := p
					p.D += k.D
					p.A += k.A
					win := func() bool {
						p, e := p, e
						for {
							d := p.D - e.A
							if d < 1 {
								d = 1
							}
							e.HP -= d
							if e.HP <= 0 {
								return true
							}
							d = e.D - p.A
							if d < 1 {
								d = 1
							}
							p.HP -= d
							if p.HP <= 0 {
								return false
							}
						}
					}()
					if p2 {
						if !win {
							if gold == -1 || c > gold {
								gold = c
							}
						}
					} else {
						if win {
							if gold == -1 || c < gold {
								gold = c
							}
						}
					}
				}
			}
		}
	}
	return
}
