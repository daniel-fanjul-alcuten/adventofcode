package p9

import (
	"fmt"
	"sort"
	"strings"
)

type Group []fmt.Stringer

func (g Group) String() (s string) {
	s += "{"
	for i, r := range g {
		if i > 0 {
			s += ","
		}
		s += r.String()
	}
	s += "}"
	return
}

type Garbage struct {
	Text  string
	Value int
}

func (g Garbage) String() (s string) {
	s += "<"
	s += string(g.Text)
	s += ">"
	return
}

type ParseFunc func(s string, i *int) (interface{}, error)

func ParseBytes(s string, i *int, m map[byte]ParseFunc, e ParseFunc) (r interface{}, err error) {
	bytes := make([]string, len(m))
	for k := range m {
		bytes = append(bytes, string(k))
	}
	sort.Strings(bytes)
	keys := strings.Join(bytes, "")
	if *i >= len(s) {
		if e == nil {
			err = fmt.Errorf("Expected any of '%v': %v", keys, "")
		} else {
			return e(s, i)
		}
	} else if f := m[s[*i]]; f != nil {
		return f(s, i)
	} else {
		if e == nil {
			err = fmt.Errorf("Expected any of '%v': %v", keys, s[*i:])
		} else {
			return e(s, i)
		}
	}
	return
}

func ParseGroup(s string, i *int) (g Group, err error) {
	if _, err = ParseBytes(s, i, map[byte]ParseFunc{
		'{': func(s string, i *int) (r interface{}, err error) {
			*i++
			return
		},
	}, nil); err != nil {
		return
	}
	ok := true
	for ok {
		var r interface{}
		if r, err = ParseBytes(s, i, map[byte]ParseFunc{
			'{': func(s string, i *int) (r interface{}, err error) {
				return ParseGroup(s, i)
			},
			'<': func(s string, i *int) (r interface{}, err error) {
				return ParseGarbage(s, i)
			},
		}, func(s string, i *int) (r interface{}, err error) {
			return
		}); err != nil {
			return
		}
		if r != nil {
			g = append(g, r.(fmt.Stringer))
		}
		if r, err = ParseBytes(s, i, map[byte]ParseFunc{
			',': func(s string, i *int) (r interface{}, err error) {
				*i++
				return
			},
		}, func(s string, i *int) (r interface{}, err error) {
			ok = false
			return
		}); err != nil {
			return
		}
	}
	if _, err = ParseBytes(s, i, map[byte]ParseFunc{
		'}': func(s string, i *int) (r interface{}, err error) {
			*i++
			return
		},
	}, nil); err != nil {
		return
	}
	return
}

func ParseGarbage(s string, i *int) (g Garbage, err error) {
	if _, err = ParseBytes(s, i, map[byte]ParseFunc{
		'<': func(s string, i *int) (r interface{}, err error) {
			*i++
			return
		},
	}, nil); err != nil {
		return
	}
	ok := true
	for ok {
		if _, err = ParseBytes(s, i, map[byte]ParseFunc{
			'>': func(s string, i *int) (r interface{}, err error) {
				*i++
				ok = false
				return
			},
			'!': func(s string, i *int) (r interface{}, err error) {
				g.Text += string(s[*i])
				*i++
				if *i < len(s) {
					g.Text += string(s[*i])
					*i++
				}
				return
			},
		}, func(s string, i *int) (r interface{}, err error) {
			g.Text += string(s[*i])
			g.Value += 1
			*i++
			return
		}); err != nil {
			return
		}
	}
	return
}

func (g Group) Score(h int) (s int) {
	s += h
	for _, r := range g {
		if g2, ok := r.(Group); ok {
			s += g2.Score(h + 1)
		}
	}
	return
}

func (g Group) GarbageValue() (v int) {
	for _, r := range g {
		if e, ok := r.(Group); ok {
			v += e.GarbageValue()
		} else if e, ok := r.(Garbage); ok {
			v += e.Value
		}
	}
	return
}
