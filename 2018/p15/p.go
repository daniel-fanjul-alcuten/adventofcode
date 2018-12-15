package p

import (
	"sort"
)

type Position struct {
	X, Y int
}

func (p Position) Adjacents() Positions {
	return Positions{
		Position{p.X, p.Y - 1},
		Position{p.X - 1, p.Y},
		Position{p.X + 1, p.Y},
		Position{p.X, p.Y + 1},
	}
}

func (p Position) Less(q Position) bool {
	if p.Y < q.Y {
		return true
	} else if p.Y > q.Y {
		return false
	}
	return p.X < q.X
}

type Positions []Position

func (pp *Positions) Len() int {
	return len(*pp)
}

func (pp *Positions) Less(i, j int) bool {
	return (*pp)[i].Less((*pp)[j])
}

func (pp *Positions) Swap(i, j int) {
	(*pp)[i], (*pp)[j] = (*pp)[j], (*pp)[i]
}

type Unit struct {
	Team        rune
	Position    Position
	AttackPower int
	HitPoints   int
}

type Puzzle struct {
	input          []string
	walls          map[Position]struct{}
	units1, units2 map[Position]*Unit
}

func New(elfAttackPower int, input []string) *Puzzle {
	walls := map[Position]struct{}{}
	units1 := map[Position]*Unit{}
	units2 := map[Position]*Unit{}
	for y, s := range input {
		for x, r := range s {
			switch r {
			case '#':
				p := Position{x, y}
				walls[p] = struct{}{}
			case 'E':
				p := Position{x, y}
				units1[p] = &Unit{r, p, elfAttackPower, 200}
			case 'G':
				p := Position{x, y}
				units1[p] = &Unit{r, p, 3, 200}
			}
		}
	}
	return &Puzzle{input, walls, units1, units2}
}

func (z *Puzzle) P1() (output int) {
	for r := 1; ; r++ {
		full := z.Round()
		m := map[rune]int{}
		for _, u := range z.units1 {
			m[u.Team] += u.HitPoints
		}
		if len(m) < 2 {
			if !full {
				r--
			}
			output = r
			for _, hp := range m {
				output *= hp
			}
			return
		}
	}
}

func (z *Puzzle) P2() (output int) {
	elfAttackPower := 3
	for {
		*z = *New(elfAttackPower, z.input)
		numElves := 0
		for _, u := range z.units1 {
			if u.Team == 'E' {
				numElves++
			}
		}
		output = z.P1()
		for _, u := range z.units1 {
			if u.Team == 'E' {
				numElves--
			}
		}
		if numElves == 0 {
			return
		}
		elfAttackPower++
	}
}

func (z *Puzzle) Round() (full bool) {
	full = true
	pp := make(Positions, 0, len(z.units1))
	for p := range z.units1 {
		pp = append(pp, p)
	}
	sort.Sort(&pp)
	for _, p := range pp {
		u := z.units1[p]
		if u != nil && u.HitPoints > 0 {
			targets := 0
			for _, t := range z.units1 {
				if t.Team != u.Team && t.HitPoints > 0 {
					targets++
				}
			}
			for _, t := range z.units2 {
				if t.Team != u.Team && t.HitPoints > 0 {
					targets++
				}
			}
			if targets == 0 {
				full = false
			}
			z.Unit(u)
		}
	}
	z.units1, z.units2 = z.units2, z.units1
	for p, u := range z.units1 {
		if u.HitPoints <= 0 {
			delete(z.units1, p)
		}
	}
	for p, u := range z.units2 {
		if u.HitPoints <= 0 {
			delete(z.units2, p)
		}
	}
	return
}

func (z *Puzzle) Unit(u *Unit) {
	paths := []Positions{Positions{u.Position}}
	if path := z.FindPath(u.Team, paths); path != nil {
		delete(z.units1, u.Position)
		z.units2[u.Position] = u
	} else {
		for {
			paths = z.Explore(paths)
			if len(paths) == 0 {
				delete(z.units1, u.Position)
				z.units2[u.Position] = u
				break
			}
			if path := z.FindPath(u.Team, paths); path != nil {
				delete(z.units1, u.Position)
				u.Position = path[1]
				z.units2[u.Position] = u
				break
			}
		}
	}
	var enemy *Unit
	for _, p := range u.Position.Adjacents() {
		e := z.units1[p]
		if e == nil || e.Team == u.Team || e.HitPoints <= 0 {
			e = z.units2[p]
		}
		if e == nil || e.Team == u.Team || e.HitPoints <= 0 {
			continue
		}
		if enemy == nil ||
			e.HitPoints < enemy.HitPoints ||
			e.HitPoints == enemy.HitPoints && e.Position.Less(enemy.Position) {
			enemy = e
		}
	}
	if enemy != nil {
		enemy.HitPoints -= u.AttackPower
	}
}

func (z *Puzzle) FindPath(team rune, paths []Positions) (path Positions) {
	for _, pp := range paths {
		p := pp[len(pp)-1]
		for _, q := range p.Adjacents() {
			e := z.units1[q]
			if e == nil || e.Team == team || e.HitPoints <= 0 {
				e = z.units2[q]
			}
			if e == nil || e.Team == team || e.HitPoints <= 0 {
				continue
			}
			if path == nil ||
				p.Less(path[len(path)-1]) ||
				p == path[len(path)-1] && len(pp) > 1 && pp[1].Less(path[1]) {
				path = pp
			}
		}
	}
	return
}

func (z *Puzzle) Explore(ipaths []Positions) (opaths []Positions) {
	best, exclude := map[Position]Positions{}, map[Position]struct{}{}
	for _, path := range ipaths {
		p := path[len(path)-1]
		if best[p] == nil || len(path) > 1 && path[1].Less(best[p][1]) {
			best[p] = path
		}
		for _, p := range path {
			exclude[p] = struct{}{}
		}
	}
	for _, path := range best {
		p := path[len(path)-1]
		for _, q := range p.Adjacents() {
			if _, ok := exclude[q]; ok {
				continue
			}
			if _, ok := z.walls[q]; ok {
				continue
			}
			if z.units1[q] != nil && z.units1[q].HitPoints > 0 {
				continue
			}
			if z.units2[q] != nil && z.units2[q].HitPoints > 0 {
				continue
			}
			opaths = append(opaths,
				append(append(Positions(nil), path...), q),
			)
		}
	}
	return
}
