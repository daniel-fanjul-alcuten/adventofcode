package p21

import (
	"strings"
)

type Buffer [][]byte

func NewBuffer(size int) (b Buffer) {
	b = make(Buffer, size)
	for i := range b {
		b[i] = make([]byte, size)
		for j := range b[i] {
			b[i][j] = '?'
		}
	}
	return
}

type Image struct {
	Buffer Buffer
	Y, X   int
	Size   int
}

func (p Image) XX(y, x int) []byte {
	y, x = p.Y+y, p.X+x
	return p.Buffer[y][x : x+1]
}

func (p Image) YY(y int) []byte {
	y, x := p.Y+y, p.X
	return p.Buffer[y][x : x+p.Size]
}

func (p Image) ZZ() (zz [][]byte) {
	for y := 0; y < p.Size; y++ {
		zz = append(zz, p.YY(y))
	}
	return
}

func (p Image) Parse(s string) {
	for y, s := range strings.Split(s, "/") {
		copy(p.YY(y), []byte(s))
	}
}

func (p Image) String() (s string) {
	for y, r := range p.ZZ() {
		if y > 0 {
			s += "/"
		}
		s += string(r)
	}
	return
}

func (p Image) Flip2x2() {
	a, b := p.XX(0, 0)[0], p.XX(0, 1)[0]
	c, d := p.XX(1, 0)[0], p.XX(1, 1)[0]
	p.XX(0, 0)[0], p.XX(0, 1)[0] = b, a
	p.XX(1, 0)[0], p.XX(1, 1)[0] = d, c
}

func (p Image) Rotate2x2() {
	a, b := p.XX(0, 0)[0], p.XX(0, 1)[0]
	c, d := p.XX(1, 0)[0], p.XX(1, 1)[0]
	p.XX(0, 0)[0], p.XX(0, 1)[0] = c, a
	p.XX(1, 0)[0], p.XX(1, 1)[0] = d, b
}

func (p Image) Flip3x3() {
	a, b := p.XX(0, 0)[0], p.XX(0, 2)[0]
	c, d := p.XX(1, 0)[0], p.XX(1, 2)[0]
	e, f := p.XX(2, 0)[0], p.XX(2, 2)[0]
	p.XX(0, 0)[0], p.XX(0, 2)[0] = b, a
	p.XX(1, 0)[0], p.XX(1, 2)[0] = d, c
	p.XX(2, 0)[0], p.XX(2, 2)[0] = f, e
}

func (p Image) Rotate3x3() {
	a, b, c := p.XX(0, 0)[0], p.XX(0, 1)[0], p.XX(0, 2)[0]
	d, e := p.XX(1, 0)[0], p.XX(1, 2)[0]
	f, g, h := p.XX(2, 0)[0], p.XX(2, 1)[0], p.XX(2, 2)[0]
	p.XX(0, 0)[0], p.XX(0, 1)[0], p.XX(0, 2)[0] = f, d, a
	p.XX(1, 0)[0], p.XX(1, 2)[0] = g, b
	p.XX(2, 0)[0], p.XX(2, 1)[0], p.XX(2, 2)[0] = h, e, c
}

func (p Image) CopyTo(q Image) {
	for y := 0; y < p.Size; y++ {
		copy(q.YY(y), p.YY(y))
	}
	return
}

func (p Image) Clone() (q Image) {
	q = Image{NewBuffer(p.Size), 0, 0, p.Size}
	p.CopyTo(q)
	return
}

func (p Image) Count(f byte) (n int) {
	for _, r := range p.ZZ() {
		for _, c := range r {
			if c == f {
				n++
			}
		}
	}
	return
}

type Pattern struct {
	Source, Target Image
}

func (p *Pattern) Parse(s string) {
	switch len(s) {
	case 20:
		p.Source = Image{NewBuffer(2), 0, 0, 2}
		p.Target = Image{NewBuffer(3), 0, 0, 3}
	case 34:
		p.Source = Image{NewBuffer(3), 0, 0, 3}
		p.Target = Image{NewBuffer(4), 0, 0, 4}
	default:
		panic(len(s))
	}
	ss := strings.Split(s, " => ")
	p.Source.Parse(ss[0])
	p.Target.Parse(ss[1])
}

func Parse(ss []string) (qq map[string]Image) {
	qq = map[string]Image{}
	for _, s := range ss {
		p := Pattern{}
		p.Parse(s)
		switch p.Source.Size {
		case 2:
			for i := 0; i < 2; i++ {
				p.Source.Flip2x2()
				qq[p.Source.String()] = p.Target
				for j := 0; j < 4; j++ {
					p.Source.Rotate2x2()
					qq[p.Source.String()] = p.Target
				}
			}
		case 3:
			for i := 0; i < 2; i++ {
				p.Source.Flip3x3()
				qq[p.Source.String()] = p.Target
				for j := 0; j < 4; j++ {
					p.Source.Rotate3x3()
					qq[p.Source.String()] = p.Target
				}
			}
		default:
			panic(p.Source.Size)
		}
	}
	return
}

func (p Image) Enhance(qq map[string]Image) (r Image) {
	sSize, tSize := 0, 0
	if p.Size%2 == 0 {
		sSize, tSize = 2, 3
	} else {
		sSize, tSize = 3, 4
	}
	{
		s := p.Size * tSize / sSize
		r = Image{NewBuffer(s), 0, 0, s}
	}
	y := 0
	for y*sSize < p.Size {
		x := 0
		for x*sSize < p.Size {
			k := Image{p.Buffer, y * sSize, x * sSize, sSize}.String()
			q, ok := qq[k]
			if !ok {
				panic(k)
			}
			q.CopyTo(Image{r.Buffer, y * tSize, x * tSize, tSize})
			x++
		}
		y++
	}
	return
}
