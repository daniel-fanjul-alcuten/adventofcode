package p

var S = Puzzle{0, 0, Floors{
	Floor{"HM", "LM"},
	Floor{"HG"},
	Floor{"LG"},
	nil,
}}

var X1 = Puzzle{0, 0, Floors{
	Floor{"SG", "SM", "PG", "PM"},
	Floor{"TG", "RG", "RM", "CG", "CM"},
	Floor{"TM"},
	nil,
}}

var X2 = Puzzle{0, 0, Floors{
	Floor{"SG", "SM", "PG", "PM", "EG", "EM", "DG", "DM"},
	Floor{"TG", "RG", "RM", "CG", "CM"},
	Floor{"TM"},
	nil,
}}
