package p

import (
	"fmt"
	"strings"
	"time"
)

type Event struct {
	Time  time.Time
	Guard int
	Event string
}

func Parse(ss []string) (ee []Event, err error) {
	for _, s := range ss {
		i := strings.Index(s, "[")
		if i == -1 {
			err = fmt.Errorf("No i")
			return
		}
		j := strings.Index(s, "]")
		if j == -1 {
			err = fmt.Errorf("No j")
			return
		}
		var e Event
		if e.Time, err = time.Parse("2006-01-02 15:04", s[i+1:j]); err != nil {
			return
		}
		s = s[j+2:]
		if strings.HasPrefix(s, "Guard ") {
			if _, err = fmt.Sscanf(s, "Guard #%d begins shift", &e.Guard); err != nil {
				err = fmt.Errorf("For %s: %s", s, err)
				return
			}
			s = s[strings.Index(s, "begins shift"):]
		}
		e.Event = s
		ee = append(ee, e)
	}
	return
}

func P1(ee []Event) (guard int, minute int) {
	g, s, t := 0, false, ee[0].Time
	gm, gmMaxK, gmMaxV := map[int]int{}, 0, 0
	gmt := map[int]map[int]int{}
	for _, e := range ee {
		if g > 0 {
			for t.Before(e.Time) {
				if s {
					gm[g]++
					if gm[g] > gmMaxV {
						gmMaxK, gmMaxV = g, gm[g]
					}
					m1 := gmt[g]
					if m1 == nil {
						m1 = map[int]int{}
						gmt[g] = m1
					}
					m1[t.Minute()]++
				}
				t = t.Add(time.Minute)
			}
		}
		if e.Guard != 0 {
			g = e.Guard
		}
		s = e.Event == "falls asleep"
	}
	guard = gmMaxK
	m1, maxK, maxV := gmt[guard], 0, 0
	for m, v := range m1 {
		if v > maxV {
			maxK, maxV = m, v
		}
	}
	minute = maxK
	return
}

func P2(ee []Event) (guard int, minute int) {
	g, s, t := 0, false, ee[0].Time
	gmt, gmtMaxG, gmtMaxM, gmtMaxT := map[int]map[int]int{}, 0, 0, -1
	for _, e := range ee {
		if g > 0 {
			for t.Before(e.Time) {
				if s {
					m1 := gmt[g]
					if m1 == nil {
						m1 = map[int]int{}
						gmt[g] = m1
					}
					min := t.Minute()
					m1[min]++
					if m1[min] > gmtMaxT {
						gmtMaxG, gmtMaxM, gmtMaxT = g, min, m1[min]
					}
				}
				t = t.Add(time.Minute)
			}
		}
		if e.Guard != 0 {
			g = e.Guard
		}
		s = e.Event == "falls asleep"
	}
	guard, minute = gmtMaxG, gmtMaxM
	return
}
