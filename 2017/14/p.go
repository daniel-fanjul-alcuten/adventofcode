package p14

import (
	"adventofcode/2017/kh"
	"math/bits"
	"strconv"
)

func Hash(key string) (hh [][]byte, u int) {
	for i := 0; i < 128; i++ {
		k := key + "-" + strconv.Itoa(i)
		h, _ := kh.KnotHash([]byte(k))
		hh = append(hh, h)
		for _, b := range h {
			u += bits.OnesCount8(b)
		}
	}
	return
}

type Position struct {
	X, Y int
}

func ToMap(hh [][]byte) (m map[Position]struct{}) {
	m = make(map[Position]struct{}, len(hh))
	y := 0
	for _, h := range hh {
		x := 0
		for _, b := range h {
			if b&0x80 == 0x80 {
				m[Position{8 * x, y}] = struct{}{}
			}
			if b&0x40 == 0x40 {
				m[Position{8*x + 1, y}] = struct{}{}
			}
			if b&0x20 == 0x20 {
				m[Position{8*x + 2, y}] = struct{}{}
			}
			if b&0x10 == 0x10 {
				m[Position{8*x + 3, y}] = struct{}{}
			}
			if b&0x08 == 0x08 {
				m[Position{8*x + 4, y}] = struct{}{}
			}
			if b&0x04 == 0x04 {
				m[Position{8*x + 5, y}] = struct{}{}
			}
			if b&0x02 == 0x02 {
				m[Position{8*x + 6, y}] = struct{}{}
			}
			if b&0x01 == 0x01 {
				m[Position{8*x + 7, y}] = struct{}{}
			}
			x++
		}
		y++
	}
	return
}

func Regions(m map[Position]struct{}) (n int) {
	pending, done := []Position{}, map[Position]struct{}{}
	for len(m) > 0 {
		for p := range m {
			n++
			pending = append(pending, p)
			for len(pending) > 0 {
				p, pending = pending[0], pending[1:]
				if _, ok := done[p]; !ok {
					if _, ok := m[p]; ok {
						delete(m, p)
						pending = append(pending,
							Position{p.X + 1, p.Y},
							Position{p.X - 1, p.Y},
							Position{p.X, p.Y + 1},
							Position{p.X, p.Y - 1},
						)
						done[p] = struct{}{}
					}
				}
			}
		}
	}
	return
}
