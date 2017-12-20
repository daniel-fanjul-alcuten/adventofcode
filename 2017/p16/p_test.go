package p16

import (
	"testing"
)

func TestSample(t *testing.T) {
	d := Dance("abcde")
	d = d.Spin(1)
	if s := string(d); s != "eabcd" {
		t.Error(s)
	}
	d = d.Exchange(3, 4)
	if s := string(d); s != "eabdc" {
		t.Error(s)
	}
	d = d.Partner('e', 'b')
	if s := string(d); s != "baedc" {
		t.Error(s)
	}
}

func TestDanceN(t *testing.T) {
	d := Dance("abcde").DanceN("s1,x3/4,pe/b")
	if s := string(d); s != "baedc" {
		t.Error(s)
	}
}

func TestP1(t *testing.T) {
	d := Dance("abcdefghijklmnop").DanceN(P1)
	if s := string(d); s != "ociedpjbmfnkhlga" {
		t.Error(s)
	}
}

func TestP2(t *testing.T) {
	d, m := Dance("abcdefghijklmnop"), map[string]int{}
	i, l := 0, 1*1000*1000*1000
	for i < l {
		m[string(d)] = i
		d, i = d.DanceN(P1), i+1
		if j, ok := m[string(d)]; ok {
			i += ((l - i) / (i - j)) * (i - j)
		}
	}
	if s := string(d); s != "gnflbkojhicpmead" {
		t.Error(s)
	}
}
