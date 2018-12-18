package p

func P1(input string, n int) (last string, safe int) {
	row, row2 := []byte(input), make([]byte, 0, len(input))
	for i := 0; i < len(row); i++ {
		if row[i] == '.' {
			safe++
		}
	}
	for i := 0; i < n; i++ {
		for i := 0; i < len(row); i++ {
			left := byte('.')
			if i-1 >= 0 {
				left = row[i-1]
			}
			center := row[i]
			right := byte('.')
			if i+1 < len(row) {
				right = row[i+1]
			}
			next := byte('.')
			if left == '^' && center == '^' && right == '.' ||
				left == '.' && center == '^' && right == '^' ||
				left == '^' && center == '.' && right == '.' ||
				left == '.' && center == '.' && right == '^' {
				next = '^'
			} else {
				safe++
			}
			row2 = append(row2, next)
		}
		row, row2 = row2, row[:0]
	}
	last = string(row)
	return
}
