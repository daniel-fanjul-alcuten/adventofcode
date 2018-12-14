package p

func P1(input []string) (output string) {
	p := 5
	for i := 0; i < len(input); i++ {
		s := input[i]
		for j := 0; j < len(s); j++ {
			switch s[j] {
			case 'U':
				if p > 3 {
					p -= 3
				}
			case 'D':
				if p < 7 {
					p += 3
				}
			case 'L':
				if p%3 != 1 {
					p--
				}
			case 'R':
				if p%3 != 0 {
					p++
				}
			default:
				panic(s[j])
			}
		}
		output += string('0' + p)
	}
	return
}

func P2(input []string) (output string) {
	x, y, m := 0, 2, map[int]map[int]rune{
		0: map[int]rune{
			2: '1',
		},
		1: map[int]rune{
			1: '2',
			2: '3',
			3: '4',
		},
		2: map[int]rune{
			0: '5',
			1: '6',
			2: '7',
			3: '8',
			4: '9',
		},
		3: map[int]rune{
			1: 'A',
			2: 'B',
			3: 'C',
		},
		4: map[int]rune{
			2: 'D',
		},
	}
	for i := 0; i < len(input); i++ {
		s := input[i]
		for j := 0; j < len(s); j++ {
			switch s[j] {
			case 'L':
				if m[y][x-1] != 0 {
					x--
				}
			case 'R':
				if m[y][x+1] != 0 {
					x++
				}
			case 'U':
				if m[y-1][x] != 0 {
					y--
				}
			case 'D':
				if m[y+1][x] != 0 {
					y++
				}
			default:
				panic(s[j])
			}
		}
		output += string(m[y][x])
	}
	return
}
