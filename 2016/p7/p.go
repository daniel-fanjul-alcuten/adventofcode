package p

import (
	"strings"
)

func P1(input []string) (n int) {
	for _, in := range input {
		if abbaN(in) {
			n++
		}
	}
	return
}

func abbaN(s string) bool {
	out, in := false, true
	for len(s) > 0 {
		i := strings.Index(s, "[")
		if i == -1 {
			out = out || abba1(s)
			break
		}
		j := strings.Index(s, "]")
		out = out || abba1(s[:i])
		in = in && !abba1(s[i+1:j])
		s = s[j+1:]
	}
	return out && in
}

func abba1(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+1] {
			continue
		}
		if s[i] != s[i+3] {
			continue
		}
		if s[i+1] != s[i+2] {
			continue
		}
		return true
	}
	return false
}

func P2(input []string) (n int) {
	for _, s := range input {
		aba := map[string]struct{}{}
		xyx := map[string]struct{}{}
		for len(s) > 0 {
			i := strings.Index(s, "[")
			var out, in string
			if i == -1 {
				out, in = s, ""
				s = ""
			} else {
				j := strings.Index(s, "]")
				out, in = s[:i], s[i+1:j]
				s = s[j+1:]
			}
			for i := 0; i < len(out)-2; i++ {
				if out[i] == out[i+1] {
					continue
				}
				if out[i] != out[i+2] {
					continue
				}
				aba[out[i:i+3]] = struct{}{}
			}
			for i := 0; i < len(in)-2; i++ {
				if in[i] == in[i+1] {
					continue
				}
				if in[i] != in[i+2] {
					continue
				}
				xyx[in[i:i+3]] = struct{}{}
			}
		}
		ok := false
		for s := range aba {
			if _, ok = xyx[s[1:3]+s[1:2]]; ok {
				break
			}
		}
		if ok {
			n++
		}
	}
	return
}
