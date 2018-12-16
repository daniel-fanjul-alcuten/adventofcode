package p

var S = []Disc{
	{0, 4, 5},
	{0, 1, 2},
}

var X1 = []Disc{
	{0, 1, 13},
	{0, 10, 19},
	{0, 2, 3},
	{0, 1, 7},
	{0, 3, 5},
	{0, 5, 17},
}

var X2 = append(X1,
	Disc{0, 0, 11},
)

var X3 = append(X2,
	Disc{0, 0, 31},
	Disc{0, 0, 37},
	Disc{0, 0, 41},
	Disc{0, 0, 43},
	Disc{0, 0, 47},
	Disc{0, 0, 53},
	Disc{0, 0, 59},
)
