package p

type Character struct {
	HP   int
	D, A int
}

func P1(p, e Character, m int, hard bool) (mana int) {
	maxHP := p.HP
	var battle func(p, e Character, m, cm, se, pe, re int, turn bool)
	battle = func(p, e Character, m, cm, se, pe, re int, turn bool) {
		win := func(cm int) {
			if mana == 0 || cm < mana {
				mana = cm
			}
		}
		if hard && !turn {
			p.HP--
			if p.HP <= 0 {
				return
			}
		}
		if se > 0 {
			se--
			if se == 0 {
				p.A -= 7
			}
		}
		if pe > 0 {
			pe--
			e.HP -= 3
			if e.HP <= 0 {
				win(cm)
				return
			}
		}
		if re > 0 {
			re--
			m += 101
		}
		if turn {
			d := e.D - p.A
			if d < 1 {
				d = 1
			}
			p.HP -= d
			if p.HP <= 0 {
				return
			}
			battle(p, e, m, cm, se, pe, re, !turn)
			return
		}
		for _, sp := range []string{"MM", "D", "S", "P", "R"} {
			p, e, m, cm, se, pe, re := p, e, m, cm, se, pe, re
			switch sp {
			case "MM":
				if m < 53 {
					break
				}
				m -= 53
				cm += 53
				e.HP -= 4
				if e.HP <= 0 {
					win(cm)
				} else {
					battle(p, e, m, cm, se, pe, re, !turn)
				}
			case "D":
				if m < 73 {
					break
				}
				m -= 73
				cm += 73
				e.HP -= 2
				if e.HP <= 0 {
					win(cm)
				} else {
					p.HP += 2
					if p.HP > maxHP {
						p.HP = maxHP
					}
					battle(p, e, m, cm, se, pe, re, !turn)
				}
			case "S":
				if se > 0 {
					break
				}
				if m < 113 {
					break
				}
				m -= 113
				cm += 113
				p.A += 7
				se = 6
				battle(p, e, m, cm, se, pe, re, !turn)
			case "P":
				if pe > 0 {
					break
				}
				if m < 173 {
					break
				}
				m -= 173
				cm += 173
				pe = 6
				battle(p, e, m, cm, se, pe, re, !turn)
			case "R":
				if re > 0 {
					break
				}
				if m < 229 {
					break
				}
				m -= 229
				cm += 229
				re = 5
				battle(p, e, m, cm, se, pe, re, !turn)
			}
		}
	}
	battle(p, e, m, 0, 0, 0, 0, false)
	return
}
