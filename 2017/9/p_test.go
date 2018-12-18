package p9

import (
	"testing"
)

func TestX(t *testing.T) {
	i := 0
	if g, err := ParseGroup(X, &i); err != nil {
		t.Error(err)
	} else if s := g.String(); s != X {
		t.Error(s)
	} else if v := g.Score(0); v != 14421 {
		t.Error(v)
	} else if v := g.GarbageValue(); v != 6817 {
		t.Error(v)
	}
}
