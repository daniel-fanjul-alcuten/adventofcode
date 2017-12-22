package p22

import (
	"testing"
)

func TestS1(t *testing.T) {
	g := Grid{}
	g.Parse(S1)
	for {
		g.Burst()
		if g.NumBur == 7 {
			if n := g.NumInf; n != 5 {
				t.Error(n)
			}
		} else if g.NumBur == 70 {
			if n := g.NumInf; n != 41 {
				t.Error(n)
			}
		} else if g.NumBur == 10000 {
			if n := g.NumInf; n != 5587 {
				t.Error(n)
			}
		} else if g.NumBur > 10000 {
			break
		}
	}
}

func TestP1(t *testing.T) {
	g := Grid{}
	g.Parse(P1)
	for {
		g.Burst()
		if g.NumBur == 10000 {
			if n := g.NumInf; n != 5570 {
				t.Error(n)
			}
		} else if g.NumBur > 10000 {
			break
		}
	}
}

func TestS2(t *testing.T) {
	g := Grid{}
	g.Parse(S1)
	for {
		g.Burst2()
		if g.NumBur == 100 {
			if n := g.NumInf; n != 26 {
				t.Error(n)
			}
		} else if g.NumBur == 10000000 {
			if n := g.NumInf; n != 2511944 {
				t.Error(n)
			}
		} else if g.NumBur > 10000000 {
			break
		}
	}
}

func TestP2(t *testing.T) {
	g := Grid{}
	g.Parse(P1)
	for {
		g.Burst2()
		if g.NumBur == 10000000 {
			if n := g.NumInf; n != 2512022 {
				t.Error(n)
			}
		} else if g.NumBur > 10000000 {
			break
		}
	}
}
