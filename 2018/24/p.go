package p

import (
	"sort"
	"strconv"
	"strings"
)

type Group struct {
	tid, gid int
	u        int
	a, hp, n int
	at       string
	wt, it   map[string]struct{}
}

func (g Group) p() int {
	return g.u * g.a
}

type Attack [2]*Group

func (a Attack) Damage() int {
	p := a[0].p()
	if _, ok := a[1].wt[a[0].at]; ok {
		p *= 2
	}
	if _, ok := a[1].it[a[0].at]; ok {
		p = 0
	}
	return p
}

func (a Attack) KilledUnits() int {
	du := a.Damage() / a[1].hp
	if du > a[1].u {
		du = a[1].u
	}
	return du
}

type bySelecting []Attack

func (b *bySelecting) Len() int {
	return len(*b)
}

func (b *bySelecting) Less(i, j int) bool {
	ai, aj := (*b)[i], (*b)[j]
	if p := ai[0].p() - aj[0].p(); p > 0 {
		return true
	} else if p < 0 {
		return false
	}
	if n := ai[0].n - aj[0].n; n > 0 {
		return true
	} else if n < 0 {
		return false
	}
	if d := ai.Damage() - aj.Damage(); d > 0 {
		return true
	} else if d < 0 {
		return false
	}
	if p := ai[1].p() - aj[1].p(); p > 0 {
		return true
	} else if p < 0 {
		return false
	}
	if n := ai[1].n - aj[1].n; n > 0 {
		return true
	} else if n < 0 {
		return false
	}
	return false
}

func (b *bySelecting) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

type byAttacking []Attack

func (b *byAttacking) Len() int {
	return len(*b)
}

func (b *byAttacking) Less(i, j int) bool {
	ai, aj := (*b)[i], (*b)[j]
	if n := ai[0].n - aj[0].n; n > 0 {
		return true
	} else if n < 0 {
		return false
	}
	return false
}

func (b *byAttacking) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

type Army []*Group

func atoi(s string) (n int) {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}

func P1(imm, inf []string, boost int) (side, units int) {
	parse := func(s string) (g *Group) {
		g = &Group{}
		spl1 := strings.Split(s, " ")
		g.u = atoi(spl1[0])
		g.hp = atoi(spl1[4])
		i := strings.Index(s, "(")
		j := strings.Index(s, ")")
		if i > -1 {
			for _, s2 := range strings.Split(s[i+1:j], "; ") {
				vul := ""
				for _, s3 := range strings.Split(s2, ", ") {
					spl4 := strings.Split(s3, " ")
					if vul == "" {
						vul = spl4[0]
						if vul == "weak" {
							if g.wt == nil {
								g.wt = make(map[string]struct{})
							}
							g.wt[spl4[2]] = struct{}{}
						} else if vul == "immune" {
							if g.it == nil {
								g.it = make(map[string]struct{})
							}
							g.it[spl4[2]] = struct{}{}
						} else {
							panic(vul)
						}
					} else {
						if vul == "weak" {
							if g.wt == nil {
								g.wt = make(map[string]struct{})
							}
							g.wt[spl4[0]] = struct{}{}
						} else if vul == "immune" {
							if g.it == nil {
								g.it = make(map[string]struct{})
							}
							g.it[spl4[0]] = struct{}{}
						} else {
							panic(vul)
						}
					}
				}
			}
			spl2 := strings.Split(s[j+2:], " ")
			g.a = atoi(spl2[5])
			g.at = spl2[6]
			g.n = atoi(spl2[10])
		} else {
			g.a = atoi(spl1[12])
			g.at = spl1[13]
			g.n = atoi(spl1[17])
		}
		return
	}
	r1 := make(Army, len(imm))
	for i, s := range imm {
		r1[i] = parse(s)
		r1[i].tid, r1[i].gid = 0, i
		r1[i].a += boost
	}
	r2 := make(Army, len(inf))
	for i, s := range inf {
		r2[i] = parse(s)
		r2[i].tid, r2[i].gid = 1, i
	}
	for len(r1) > 0 && len(r2) > 0 {
		bs, ba := bySelecting{}, byAttacking{}
		for _, g1 := range r1 {
			for _, g2 := range r2 {
				if a := (Attack{g1, g2}); a.Damage() > 0 {
					bs = append(bs, a)
				}
				if a := (Attack{g2, g1}); a.Damage() > 0 {
					bs = append(bs, a)
				}
			}
		}
		if len(bs) == 0 {
			side = -1
			units = -1
			return
		}
		sort.Sort(&bs)
		m0, m1 := map[*Group]struct{}{}, map[*Group]struct{}{}
		for len(bs) > 0 {
			a := bs[0]
			bs = bs[1:]
			if _, ok := m0[a[0]]; ok {
				continue
			}
			if _, ok := m1[a[1]]; ok {
				continue
			}
			ba = append(ba, a)
			m0[a[0]] = struct{}{}
			m1[a[1]] = struct{}{}
		}
		sort.Sort(&ba)
		tku := 0
		for _, a := range ba {
			g1, g2 := a[0], a[1]
			if g1.u <= 0 {
				continue
			}
			if g2.u <= 0 {
				continue
			}
			ku := a.KilledUnits()
			g2.u -= ku
			tku += ku
			if g2.u <= 0 {
				clean := func(r *Army) {
					for i, g := range *r {
						if g == g2 {
							*r = append((*r)[:i], (*r)[i+1:]...)
							break
						}
					}
				}
				clean(&r1)
				clean(&r2)
			}
		}
		if tku == 0 {
			side = -1
			units = -1
			return
		}
	}
	for _, g := range r1 {
		side = 0
		units += g.u
	}
	for _, g := range r2 {
		side = 1
		units += g.u
	}
	return
}

func P2(imm, inf []string) (boost, units int) {
	for boost = 0; ; boost++ {
		var side int
		side, units = P1(imm, inf, boost)
		if side != 0 {
			continue
		}
		return
	}
}
