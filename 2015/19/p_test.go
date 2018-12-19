package p

import "testing"

func TestP1(t *testing.T) {
	if n := len(P1(S1, S2)); n != 4 {
		t.Error(n)
	}
	if n := len(P1("HOHOHO", S2)); n != 7 {
		t.Error(n)
	}
	if n := len(P1(X1, X2)); n != 509 {
		t.Error(n)
	}
}

var S1 = "HOH"
var S2 = [][2]string{
	{"e", "H"},
	{"e", "O"},
	{"H", "HO"},
	{"H", "OH"},
	{"O", "HH"},
}

var X1 = "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"
var X2 = [][2]string{
	{"Al", "ThF"},
	{"Al", "ThRnFAr"},
	{"B", "BCa"},
	{"B", "TiB"},
	{"B", "TiRnFAr"},
	{"Ca", "CaCa"},
	{"Ca", "PB"},
	{"Ca", "PRnFAr"},
	{"Ca", "SiRnFYFAr"},
	{"Ca", "SiRnMgAr"},
	{"Ca", "SiTh"},
	{"F", "CaF"},
	{"F", "PMg"},
	{"F", "SiAl"},
	{"H", "CRnAlAr"},
	{"H", "CRnFYFYFAr"},
	{"H", "CRnFYMgAr"},
	{"H", "CRnMgYFAr"},
	{"H", "HCa"},
	{"H", "NRnFYFAr"},
	{"H", "NRnMgAr"},
	{"H", "NTh"},
	{"H", "OB"},
	{"H", "ORnFAr"},
	{"Mg", "BF"},
	{"Mg", "TiMg"},
	{"N", "CRnFAr"},
	{"N", "HSi"},
	{"O", "CRnFYFAr"},
	{"O", "CRnMgAr"},
	{"O", "HP"},
	{"O", "NRnFAr"},
	{"O", "OTi"},
	{"P", "CaP"},
	{"P", "PTi"},
	{"P", "SiRnFAr"},
	{"Si", "CaSi"},
	{"Th", "ThCa"},
	{"Ti", "BP"},
	{"Ti", "TiTi"},
	{"e", "HF"},
	{"e", "NAl"},
	{"e", "OMg"},
}
