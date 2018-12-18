package p25

type Operation struct {
	Write byte
	Move  int
	Next  byte
}

type Machine struct {
	State  byte
	Steps  int
	States map[byte][]Operation
}

var S1 = Machine{
	'a', 6, map[byte][]Operation{
		'a': []Operation{
			{1, 1, 'b'}, {0, -1, 'b'},
		},
		'b': []Operation{
			{1, -1, 'a'}, {1, 1, 'a'},
		},
	},
}

var P1 = Machine{
	'a', 12317297, map[byte][]Operation{
		'a': []Operation{
			{1, 1, 'b'}, {0, -1, 'd'},
		},
		'b': []Operation{
			{1, 1, 'c'}, {0, 1, 'f'},
		},
		'c': []Operation{
			{1, -1, 'c'}, {1, -1, 'a'},
		},
		'd': []Operation{
			{0, -1, 'e'}, {1, 1, 'a'},
		},
		'e': []Operation{
			{1, -1, 'a'}, {0, 1, 'b'},
		},
		'f': []Operation{
			{0, 1, 'c'}, {0, 1, 'e'},
		},
	},
}
