package p

import (
	"strconv"
	"strings"
)

func P1(input []string, seconds int) (output int) {
	for _, in := range input {
		sp := strings.Split(in, " ")
		s, err := strconv.Atoi(sp[3])
		if err != nil {
			panic(err)
		}
		t, err := strconv.Atoi(sp[6])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(sp[13])
		if err != nil {
			panic(err)
		}
		d, sec := 0, 0
		for sec < seconds {
			for i := 0; i < t && sec < seconds; i++ {
				d += s
				sec++
			}
			for i := 0; i < r && sec < seconds; i++ {
				sec++
			}
		}
		if output == 0 || d > output {
			output = d
		}
	}
	return
}

func P2(input []string, seconds int) (output int) {
	type R struct {
		N       string
		S, T, R int
	}
	reindeers := []R{}
	for _, in := range input {
		sp := strings.Split(in, " ")
		s, err := strconv.Atoi(sp[3])
		if err != nil {
			panic(err)
		}
		t, err := strconv.Atoi(sp[6])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(sp[13])
		if err != nil {
			panic(err)
		}
		reindeers = append(reindeers, R{sp[0], s, t, r})
	}
	type S struct {
		R    bool
		T    int
		D, P int
	}
	states := map[string]*S{}
	for _, r := range reindeers {
		states[r.N] = &S{}
	}
	for sec := 0; sec < seconds; sec++ {
		for _, r := range reindeers {
			s := states[r.N]
			if !s.R {
				s.D += r.S
			}
			s.T++
			if !s.R && s.T >= r.T || s.R && s.T >= r.R {
				s.R = !s.R
				s.T = 0
			}
		}
		head := []*S{}
		for _, r := range reindeers {
			s := states[r.N]
			if len(head) == 0 || s.D > head[0].D {
				head = append(head[:0], s)
			} else if s.D == head[0].D {
				head = append(head, s)
			}
		}
		for _, s := range head {
			s.P++
		}
	}
	for _, s := range states {
		if output == 0 || s.P > output {
			output = s.P
		}
	}
	return
}
