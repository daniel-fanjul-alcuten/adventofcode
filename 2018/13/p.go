package p

type Position struct {
	X, Y int
}

type Cart struct {
	Char  byte
	Turns int
}

func P1(input []string, ticks int) *Position {
	carts1, carts2 := map[Position][]Cart{}, map[Position][]Cart{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			switch input[y][x] {
			case '<', '>', '^', 'v':
				carts1[Position{x, y}] = []Cart{{input[y][x], 0}}
			}
		}
	}
	for i := 0; i < ticks; i++ {
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				p := Position{x, y}
				if len(carts1[p]) == 0 {
					continue
				}
				cart := carts1[p][0]
				delete(carts1, p)
				next(input, &p, &cart)
				carts2[p] = append(carts2[p], cart)
				if len(carts1[p])+len(carts2[p]) > 1 {
					return &p
				}
			}
		}
		carts1, carts2 = carts2, carts1
	}
	return nil
}

func P2(input []string, ticks int) *Position {
	carts1, carts2 := map[Position][]Cart{}, map[Position][]Cart{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			switch input[y][x] {
			case '<', '>', '^', 'v':
				carts1[Position{x, y}] = []Cart{{input[y][x], 0}}
			}
		}
	}
	for i := 0; i < ticks; i++ {
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				p := Position{x, y}
				if len(carts1[p]) == 0 {
					continue
				}
				cart := carts1[p][0]
				delete(carts1, p)
				next(input, &p, &cart)
				carts2[p] = append(carts2[p], cart)
				if len(carts1[p])+len(carts2[p]) > 1 {
					delete(carts1, p)
					delete(carts2, p)
				}
			}
		}
		carts1, carts2 = carts2, carts1
		if len(carts1) == 0 {
			return &Position{-1, -1}
		} else if len(carts1) == 1 {
			for p := range carts1 {
				return &p
			}
		}
	}
	return nil
}

func next(input []string, p *Position, cart *Cart) {
	switch cart.Char {
	case '<':
		switch input[p.Y][p.X] {
		case '<', '>', '-':
			p.X--
		case '/':
			p.Y++
			cart.Char = 'v'
		case '\\':
			p.Y--
			cart.Char = '^'
		case '+':
			switch cart.Turns {
			case 0:
				p.Y++
				cart.Char = 'v'
				cart.Turns = 1
			case 1:
				p.X--
				cart.Turns = 2
			case 2:
				p.Y--
				cart.Char = '^'
				cart.Turns = 0
			}
		default:
			panic(string(input[p.Y][p.X]))
		}
	case '>':
		switch input[p.Y][p.X] {
		case '>', '<', '-':
			p.X++
		case '/':
			p.Y--
			cart.Char = '^'
		case '\\':
			p.Y++
			cart.Char = 'v'
		case '+':
			switch cart.Turns {
			case 0:
				p.Y--
				cart.Char = '^'
				cart.Turns = 1
			case 1:
				p.X++
				cart.Turns = 2
			case 2:
				p.Y++
				cart.Char = 'v'
				cart.Turns = 0
			}
		default:
			panic(string(input[p.Y][p.X]))
		}
	case '^':
		switch input[p.Y][p.X] {
		case '^', 'v', '|':
			p.Y--
		case '/':
			p.X++
			cart.Char = '>'
		case '\\':
			p.X--
			cart.Char = '<'
		case '+':
			switch cart.Turns {
			case 0:
				p.X--
				cart.Char = '<'
				cart.Turns = 1
			case 1:
				p.Y--
				cart.Turns = 2
			case 2:
				p.X++
				cart.Char = '>'
				cart.Turns = 0
			}
		default:
			panic(string(input[p.Y][p.X]))
		}
	case 'v':
		switch input[p.Y][p.X] {
		case 'v', '^', '|':
			p.Y++
		case '/':
			p.X--
			cart.Char = '<'
		case '\\':
			p.X++
			cart.Char = '>'
		case '+':
			switch cart.Turns {
			case 0:
				p.X++
				cart.Char = '>'
				cart.Turns = 1
			case 1:
				p.Y++
				cart.Turns = 2
			case 2:
				p.X--
				cart.Char = '<'
				cart.Turns = 0
			}
		default:
			panic(string(input[p.Y][p.X]))
		}
	default:
		panic(cart)
	}
}
