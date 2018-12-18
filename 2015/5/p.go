package p

import "strings"

func P1(input []string) (output int) {
	for _, in := range input {
		v, t2 := 0, false
		for j := 0; j < len(in); j++ {
			c := in[j]
			switch c {
			case 'a', 'e', 'i', 'o', 'u':
				v++
			}
			if !t2 && j > 0 && c == in[j-1] {
				t2 = true
			}
		}
		if v < 3 {
			continue
		}
		if !t2 {
			continue
		}
		if strings.Contains(in, "ab") {
			continue
		}
		if strings.Contains(in, "cd") {
			continue
		}
		if strings.Contains(in, "pq") {
			continue
		}
		if strings.Contains(in, "xy") {
			continue
		}
		output++
	}
	return
}

func P2(input []string) (output int) {
	for _, in := range input {
		found := false
		for i := 0; i < len(in)-1; i++ {
			found = strings.Index(in[i+2:], in[i:i+2]) != -1
			if found {
				break
			}
		}
		if !found {
			continue
		}
		found = false
		for i := 0; i < len(in)-2; i++ {
			found = in[i] == in[i+2]
			if found {
				break
			}
		}
		if !found {
			continue
		}
		output++
	}
	return
}
