package p20

import (
	"math"
	"strconv"
	"strings"
)

func DistFromZeroDiff(a, b int) int {
	aa, bb := a, b
	if aa < 0 {
		aa = -aa
	}
	if bb < 0 {
		bb = -bb
	}
	return bb - aa
}

type Vector struct {
	X, Y, Z int
}

func (a Vector) DistFromZeroDiff(b Vector) (c Vector) {
	c.X = DistFromZeroDiff(a.X, b.X)
	c.Y = DistFromZeroDiff(a.Y, b.Y)
	c.Z = DistFromZeroDiff(a.Z, b.Z)
	return
}

func (v Vector) Dist() int {
	return v.X + v.Y + v.Z
}

type Particle struct {
	ID            int
	Pos, Vel, Acc Vector
}

func Parse1(s string) (t Particle) {
	pva := strings.Split(s, ", ")
	p := strings.Split(strings.Trim(pva[0], "p=<>"), ",")
	v := strings.Split(strings.Trim(pva[1], "v=<>"), ",")
	a := strings.Split(strings.Trim(pva[2], "a=<>"), ",")
	var err error
	if t.Pos.X, err = strconv.Atoi(p[0]); err != nil {
		panic(err)
	}
	if t.Pos.Y, err = strconv.Atoi(p[1]); err != nil {
		panic(err)
	}
	if t.Pos.Z, err = strconv.Atoi(p[2]); err != nil {
		panic(err)
	}
	if t.Vel.X, err = strconv.Atoi(v[0]); err != nil {
		panic(err)
	}
	if t.Vel.Y, err = strconv.Atoi(v[1]); err != nil {
		panic(err)
	}
	if t.Vel.Z, err = strconv.Atoi(v[2]); err != nil {
		panic(err)
	}
	if t.Acc.X, err = strconv.Atoi(a[0]); err != nil {
		panic(err)
	}
	if t.Acc.Y, err = strconv.Atoi(a[1]); err != nil {
		panic(err)
	}
	if t.Acc.Z, err = strconv.Atoi(a[2]); err != nil {
		panic(err)
	}
	return
}

func ParseN(ss []string) (tt []Particle) {
	for i, s := range ss {
		t := Parse1(s)
		t.ID = i
		tt = append(tt, t)
	}
	return
}

type ParticlesByCloseToZero []Particle

func (p *ParticlesByCloseToZero) Len() int {
	return len(*p)
}

func (p *ParticlesByCloseToZero) Less(i, j int) bool {
	ii, jj := (*p)[i], (*p)[j]
	if d := ii.Acc.DistFromZeroDiff(jj.Acc).Dist(); d > 0 {
		return true
	} else if d < 0 {
		return false
	}
	if d := ii.Vel.DistFromZeroDiff(jj.Vel).Dist(); d > 0 {
		return true
	} else if d < 0 {
		return false
	}
	if d := ii.Pos.DistFromZeroDiff(jj.Pos).Dist(); d > 0 {
		return true
	} else if d < 0 {
		return false
	}
	return false
}

func (p *ParticlesByCloseToZero) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func PositionAtTime(p, v, a int, t int) int {
	// Pos+t*Vel+t*(t+1)*Acc/2
	// Pos+(Vel+Acc/2)*t+(Acc/2)*t*t
	return p + v*t + t*(t+1)*a/2
}

func CollisionTimes(p1, v1, a1, p2, v2, a2 int) (t1, t2 float64) {
	// d(Pos)+d(Vel+Acc/2)*t+d(Acc/2)*t*t == 0
	// d(2*Pos)+d(2*Vel+Acc)*t+d(Acc)*t*t == 0
	a := 2 * (p2 - p1)
	b := 2*(v2-v1) + a2 - a1
	c := a2 - a1
	if c == 0 {
		if b == 0 {
			// d(2*Pos) == 0
			if a == 0 {
				t1 = 0
			} else {
				t1 = math.Inf(1)
			}
			t2 = t1
			return
		}
		// d(2*Pos)+d(2*Vel+Acc)*t == 0
		t1 = float64(-a) / float64(b)
		t2 = t1
		return
	}
	q := math.Sqrt(float64(b*b - 4*a*c))
	t1 = (float64(-b) + q) / float64(2*c)
	t2 = (float64(-b) - q) / float64(2*c)
	return
}

func CollisionSet(p1, v1, a1, p2, v2, a2 int) (tt map[int]struct{}) {
	t1, t2 := CollisionTimes(p1, v1, a1, p2, v2, a2)
	tt = map[int]struct{}{}
	for _, tf := range []float64{t1, t2} {
		if !math.IsNaN(tf) && !math.IsInf(tf, 0) {
			t := int(tf)
			if t >= 0 {
				c1 := PositionAtTime(p1, v1, a1, t)
				c2 := PositionAtTime(p2, v2, a2, t)
				if c1 == c2 {
					tt[t] = struct{}{}
				}
			}
		}
	}
	return
}

func (p Particle) Collides(q Particle) bool {
	t1 := CollisionSet(p.Pos.X, p.Vel.X, p.Acc.X, q.Pos.X, q.Vel.X, q.Acc.X)
	t2 := CollisionSet(p.Pos.Y, p.Vel.Y, p.Acc.Y, q.Pos.Y, q.Vel.Y, q.Acc.Y)
	t3 := CollisionSet(p.Pos.Z, p.Vel.Z, p.Acc.Z, q.Pos.Z, q.Vel.Z, q.Acc.Z)
	f := func(a, b map[int]struct{}) (c map[int]struct{}) {
		c = map[int]struct{}{}
		for t := range a {
			if _, ok := b[t]; ok {
				c[t] = struct{}{}
			}
		}
		return
	}
	return len(f(f(t1, t2), t3)) > 0
}
