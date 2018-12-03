package p

import (
	"fmt"
	"strings"
)

type Claim struct {
	ID            string
	Left, Top     int
	Width, Height int
}

func (c *Claim) Parse(s string) (err error) {
	p := strings.Split(s, " ")
	if len(p) != 4 {
		err = fmt.Errorf("Invalid array size %v", len(p))
	}
	c.ID = p[0]
	if p[1] != "@" {
		err = fmt.Errorf("Expected '@', unexpected %v", p[1])
	}
	i := strings.Index(p[2], ",")
	if i == -1 {
		err = fmt.Errorf("Expected ',', unexpected %v", p[2])
	}
	j := strings.Index(p[2], ":")
	if j == -1 {
		err = fmt.Errorf("Expected ':', unexpected %v", p[2])
	}
	if _, err = fmt.Sscanf(p[2][:i], "%d", &c.Left); err != nil {
		return
	}
	if _, err = fmt.Sscanf(p[2][i+1:j], "%d", &c.Top); err != nil {
		return
	}
	k := strings.Index(p[3], "x")
	if k == -1 {
		err = fmt.Errorf("Expected 'x', unexpected %v", p[3])
	}
	if _, err = fmt.Sscanf(p[3][:k], "%d", &c.Width); err != nil {
		return
	}
	if _, err = fmt.Sscanf(p[3][k+1:], "%d", &c.Height); err != nil {
		return
	}
	return
}

func P1(cc ...Claim) (n int) {
	m1 := map[int]map[int]int{}
	for _, c := range cc {
		for x := c.Left; x < c.Left+c.Width; x++ {
			m2 := m1[x]
			if m2 == nil {
				m2 = map[int]int{}
				m1[x] = m2
			}
			for y := c.Top; y < c.Top+c.Height; y++ {
				m2[y]++
			}
		}
	}
	for _, m2 := range m1 {
		for _, c := range m2 {
			if c >= 2 {
				n++
			}
		}
	}
	return
}

func P2(cc ...Claim) (ids []string) {
	mS, m1 := map[string]struct{}{}, map[int]map[int][]string{}
	for _, c := range cc {
		mS[c.ID] = struct{}{}
		for x := c.Left; x < c.Left+c.Width; x++ {
			m2 := m1[x]
			if m2 == nil {
				m2 = map[int][]string{}
				m1[x] = m2
			}
			for y := c.Top; y < c.Top+c.Height; y++ {
				m2[y] = append(m2[y], c.ID)
			}
		}
	}
	for _, m2 := range m1 {
		for _, ss := range m2 {
			if len(ss) >= 2 {
				for _, s := range ss {
					delete(mS, s)
				}
			}
		}
	}
	for s := range mS {
		ids = append(ids, s)
	}
	return
}
