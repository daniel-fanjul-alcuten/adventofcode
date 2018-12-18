package p

import "strings"

func P1(s string) (t string) {
	i := 0
	for i+1 < len(s) {
		var b1 bool
		if s[i] >= 'a' && s[i] <= 'z' {
			b1 = false
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			b1 = true
		} else {
			panic(s[i])
		}
		var b2 bool
		if s[i+1] >= 'a' && s[i+1] <= 'z' {
			b2 = false
		} else if s[i+1] >= 'A' && s[i+1] <= 'Z' {
			b2 = true
		} else {
			panic(s[i+1])
		}
		react := b1 && !b2 && s[i+1]-'a' == s[i]-'A' ||
			b2 && !b1 && s[i]-'a' == s[i+1]-'A'
		if !react {
			i++
			continue
		}
		s = s[:i] + s[i+2:]
		if i > 0 {
			i = i - 1
		}
	}
	t = s
	return
}

func P2(s string) (t string) {
	m := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		m[s[i]] = struct{}{}
	}
	t = s
	for b := range m {
		s := strings.Replace(s, strings.ToLower(string(b)), "", -1)
		s = strings.Replace(s, strings.ToUpper(string(b)), "", -1)
		p1 := P1(s)
		if len(p1) < len(t) {
			t = p1
		}
	}
	return
}
