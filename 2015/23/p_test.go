package p

import "testing"

func TestP1_S(t *testing.T) {
	if a, _ := P1(S, 0); a != 2 {
		t.Error(a)
	}
}

func TestP1_X(t *testing.T) {
	if _, b := P1(X, 0); b != 307 {
		t.Error(b)
	}
}

func TestP2_X(t *testing.T) {
	if _, b := P1(X, 1); b != 160 {
		t.Error(b)
	}
}

var S = []string{
	"inc a",
	"jio a, +2",
	"tpl a",
	"inc a",
}

var X = []string{
	"jio a, +18",
	"inc a",
	"tpl a",
	"inc a",
	"tpl a",
	"tpl a",
	"tpl a",
	"inc a",
	"tpl a",
	"inc a",
	"tpl a",
	"inc a",
	"inc a",
	"tpl a",
	"tpl a",
	"tpl a",
	"inc a",
	"jmp +22",
	"tpl a",
	"inc a",
	"tpl a",
	"inc a",
	"inc a",
	"tpl a",
	"inc a",
	"tpl a",
	"inc a",
	"inc a",
	"tpl a",
	"tpl a",
	"inc a",
	"inc a",
	"tpl a",
	"inc a",
	"inc a",
	"tpl a",
	"inc a",
	"inc a",
	"tpl a",
	"jio a, +8",
	"inc b",
	"jie a, +4",
	"tpl a",
	"inc a",
	"jmp +2",
	"hlf a",
	"jmp -7",
}
