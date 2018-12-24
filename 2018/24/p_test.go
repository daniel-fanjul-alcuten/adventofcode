package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	s, u := P1(S1, S2, 0)
	if s != 1 {
		t.Error(s)
	}
	if u != 5216 {
		t.Error(u)
	}
}

func TestP1_X1(t *testing.T) {
	s, u := P1(X1, X2, 0)
	if s != 1 {
		t.Error(s)
	}
	if u != 14377 {
		t.Error(u)
	}
}

func TestP2_S(t *testing.T) {
	b, u := P2(S1, S2)
	if b != 1570 {
		t.Error(b)
	}
	if u != 51 {
		t.Error(u)
	}
}

func TestP2_X(t *testing.T) {
	b, u := P2(X1, X2)
	if b != 34 {
		t.Error(b)
	}
	if u != 6947 {
		t.Error(u)
	}
}

var S1 = []string{
	"17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2",
	"989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3",
}

var S2 = []string{
	"801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1",
	"4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4",
}

var X1 = []string{
	"916 units each with 3041 hit points (weak to cold, fire) with an attack that does 29 fire damage at initiative 13",
	"1959 units each with 7875 hit points (weak to cold; immune to slashing, bludgeoning) with an attack that does 38 radiation damage at initiative 20",
	"8933 units each with 5687 hit points with an attack that does 6 slashing damage at initiative 15",
	"938 units each with 8548 hit points with an attack that does 89 radiation damage at initiative 4",
	"1945 units each with 3360 hit points (immune to cold; weak to radiation) with an attack that does 16 cold damage at initiative 1",
	"2211 units each with 7794 hit points (weak to slashing) with an attack that does 30 fire damage at initiative 12",
	"24 units each with 3693 hit points with an attack that does 1502 fire damage at initiative 5",
	"2004 units each with 4141 hit points (immune to radiation) with an attack that does 18 slashing damage at initiative 19",
	"3862 units each with 3735 hit points (immune to bludgeoning, fire) with an attack that does 9 fire damage at initiative 10",
	"8831 units each with 3762 hit points (weak to radiation) with an attack that does 3 fire damage at initiative 7",
}

var X2 = []string{
	"578 units each with 55836 hit points with an attack that does 154 radiation damage at initiative 9",
	"476 units each with 55907 hit points (weak to fire) with an attack that does 208 cold damage at initiative 18",
	"496 units each with 33203 hit points (weak to fire, radiation; immune to cold, bludgeoning) with an attack that does 116 slashing damage at initiative 14",
	"683 units each with 12889 hit points (weak to fire) with an attack that does 35 bludgeoning damage at initiative 11",
	"1093 units each with 29789 hit points (immune to cold, fire) with an attack that does 51 radiation damage at initiative 17",
	"2448 units each with 40566 hit points (immune to bludgeoning, fire; weak to cold) with an attack that does 25 slashing damage at initiative 16",
	"1229 units each with 6831 hit points (weak to fire, cold; immune to slashing) with an attack that does 8 bludgeoning damage at initiative 8",
	"3680 units each with 34240 hit points (immune to bludgeoning; weak to fire, cold) with an attack that does 17 radiation damage at initiative 3",
	"4523 units each with 9788 hit points (immune to bludgeoning, fire, slashing) with an attack that does 3 bludgeoning damage at initiative 6",
	"587 units each with 49714 hit points (weak to bludgeoning) with an attack that does 161 fire damage at initiative 2",
}
