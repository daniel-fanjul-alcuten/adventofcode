package p

import (
	"sort"
	"strings"
)

func P1(pp [][]int) (s string) {
	tt := map[int]struct{}{}
	dd, rr := map[int][]int{}, map[int][]int{}
	for _, p := range pp {
		dd[p[0]] = append(dd[p[0]], p[1])
		rr[p[1]] = append(rr[p[1]], p[0])
		tt[p[0]] = struct{}{}
		tt[p[1]] = struct{}{}
	}
	for r := range rr {
		delete(tt, r)
	}
	for len(tt) > 0 {
		ttt := make([]int, 0, len(tt))
		for t := range tt {
			ttt = append(ttt, t)
		}
		sort.Ints(ttt)
		n := ttt[0]
		delete(tt, n)
		s += string(n)
		for _, d := range dd[n] {
			ok := true
			for _, r := range rr[d] {
				ok = ok && strings.Index(s, string(r)) != -1
			}
			if ok {
				tt[d] = struct{}{}
			}
		}
	}
	return
}

func P2(pp [][]int, nw, tw int) (s string, dur int) {
	tt := map[int]struct{}{}
	dd, rr := map[int][]int{}, map[int][]int{}
	for _, p := range pp {
		dd[p[0]] = append(dd[p[0]], p[1])
		rr[p[1]] = append(rr[p[1]], p[0])
		tt[p[0]] = struct{}{}
		tt[p[1]] = struct{}{}
	}
	for r := range rr {
		delete(tt, r)
	}
	wwt, wwn := []int{}, []int{}
	for len(tt) > 0 || len(wwt) > 0 {
		for len(tt) > 0 && len(wwt) < nw {
			ttt := make([]int, 0, len(tt))
			for t := range tt {
				ttt = append(ttt, t)
			}
			sort.Ints(ttt)
			n := ttt[0]
			delete(tt, n)
			wwt = append(wwt, dur+tw+n-'A'+1)
			wwn = append(wwn, n)
		}
		if len(wwt) > 0 {
			t := 0
			for i := 1; i < len(wwt); i++ {
				if wwt[i] < wwt[t] {
					t = i
				}
			}
			dur = wwt[t]
			wwt = append(wwt[:t], wwt[t+1:]...)
			n := wwn[t]
			s, wwn = s+string(n), append(wwn[:t], wwn[t+1:]...)
			for _, d := range dd[n] {
				ok := true
				for _, r := range rr[d] {
					ok = ok && strings.Index(s, string(r)) != -1
				}
				if ok {
					tt[d] = struct{}{}
				}
			}
		}
	}
	return
}
