package p

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"sort"
	"strings"
)

func triplet(h string) (b string, ok bool) {
	for i := 0; i < len(h)-2; i++ {
		if h[i] != h[i+1] {
			continue
		}
		if h[i] != h[i+2] {
			continue
		}
		return h[i : i+1], true
	}
	return "", false
}

func fivelet(h string, b string) bool {
	return strings.Contains(h, strings.Repeat(b, 5))
}

func P1(input string, ks func(h hash.Hash, hh string) string) int {
	i, h, s, m, f := 0, md5.New(), []byte(nil), map[int]string{}, []int{}
	for {
		h.Reset()
		if _, err := fmt.Fprintf(h, "%v%d", input, i); err != nil {
			panic(err)
		}
		s = h.Sum(s[:0])
		hh := hex.EncodeToString(s)
		if ks != nil {
			hh = ks(h, hh)
		}
		for j, b := range m {
			if i > j+1000 {
				delete(m, j)
				continue
			}
			if fivelet(hh, b) {
				f = append(f, j)
				sort.Ints(f)
				delete(m, j)
				continue
			}
		}
		if len(f) < 64 {
			if b, ok := triplet(hh); ok {
				m[i] = b
			}
			i++
			continue
		}
		if len(m) > 0 || i < f[len(f)-1]+1000 {
			i++
			continue
		}
		break
	}
	return f[63]
}

func P2(h hash.Hash, hh string) string {
	s := []byte(nil)
	for i := 0; i < 2016; i++ {
		h.Reset()
		if _, err := h.Write([]byte(hh)); err != nil {
			panic(err)
		}
		s = h.Sum(s[:0])
		hh = hex.EncodeToString(s)
	}
	return hh
}
