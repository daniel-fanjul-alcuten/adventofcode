package p

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type P struct {
	X, Y int
}

type State struct {
	P    P
	Path string
}

func P1(start, end P, input string) (path string) {
	s := State{start, ""}
	h, hsum, states := md5.New(), []byte{}, []State{s}
	for len(states) > 0 {
		s, states = states[0], states[1:]
		if s.P == end {
			if path == "" || len(s.Path) < len(path) {
				path = s.Path
			}
			continue
		}
		if len(path) > 0 && len(s.Path) >= len(path)-1 {
			continue
		}
		h.Reset()
		if _, err := fmt.Fprintf(h, "%v%v", input, s.Path); err != nil {
			panic(err)
		}
		hsum = h.Sum(hsum[:0])
		h := hex.EncodeToString(hsum)
		good := func(b byte) bool {
			return b >= 'b' && b <= 'f'
		}
		p := s.P
		if p.Y > start.Y && good(h[0]) {
			s := State{P{p.X, p.Y - 1}, s.Path + "U"}
			states = append(states, s)
		}
		if p.Y < end.Y && good(h[1]) {
			s := State{P{p.X, p.Y + 1}, s.Path + "D"}
			states = append(states, s)
		}
		if p.X > start.X && good(h[2]) {
			s := State{P{p.X - 1, p.Y}, s.Path + "L"}
			states = append(states, s)
		}
		if p.X < end.X && good(h[3]) {
			s := State{P{p.X + 1, p.Y}, s.Path + "R"}
			states = append(states, s)
		}
	}
	return
}

func P2(start, end P, input string) (path string) {
	s := State{start, ""}
	h, hsum, states := md5.New(), []byte{}, []State{s}
	for len(states) > 0 {
		s, states = states[0], states[1:]
		if s.P == end {
			if path == "" || len(s.Path) > len(path) {
				path = s.Path
			}
			continue
		}
		h.Reset()
		if _, err := fmt.Fprintf(h, "%v%v", input, s.Path); err != nil {
			panic(err)
		}
		hsum = h.Sum(hsum[:0])
		h := hex.EncodeToString(hsum)
		good := func(b byte) bool {
			return b >= 'b' && b <= 'f'
		}
		p := s.P
		if p.Y > start.Y && good(h[0]) {
			s := State{P{p.X, p.Y - 1}, s.Path + "U"}
			states = append(states, s)
		}
		if p.Y < end.Y && good(h[1]) {
			s := State{P{p.X, p.Y + 1}, s.Path + "D"}
			states = append(states, s)
		}
		if p.X > start.X && good(h[2]) {
			s := State{P{p.X - 1, p.Y}, s.Path + "L"}
			states = append(states, s)
		}
		if p.X < end.X && good(h[3]) {
			s := State{P{p.X + 1, p.Y}, s.Path + "R"}
			states = append(states, s)
		}
	}
	return
}
