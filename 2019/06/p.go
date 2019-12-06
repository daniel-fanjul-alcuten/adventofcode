package p

func P1(ii ...[]string) (n int) {
	l, m := map[string]struct{}{}, map[string]string{}
	for _, i := range ii {
		l[i[0]] = struct{}{}
		l[i[1]] = struct{}{}
		m[i[1]] = i[0]
	}
	for o := range l {
		o2, ok := m[o]
		for ok {
			n++
			o = o2
			o2, ok = m[o]
		}
	}
	return
}

func search(source, target string, s map[string][]string, v map[string]struct{}) (n int, ok bool) {
	if source == target {
		ok = true
		return
	}
	for _, t := range s[source] {
		if _, ok := v[t]; ok {
			continue
		}
		v[t] = struct{}{}
		nt, okt := search(t, target, s, v)
		if okt && (!ok || nt+1 < n) {
			n, ok = nt+1, okt
		}
		delete(v, t)
	}
	return
}

func P2(ii ...[]string) int {
	m, s := map[string]string{}, map[string][]string{}
	for _, i := range ii {
		m[i[1]] = i[0]
		s[i[0]] = append(s[i[0]], i[1])
		s[i[1]] = append(s[i[1]], i[0])
	}
	source, target := m["YOU"], m["SAN"]
	n, ok := search(source, target, s, map[string]struct{}{source: struct{}{}})
	if !ok {
		panic(ok)
	}
	return n
}
