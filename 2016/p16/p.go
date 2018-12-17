package p

func P1(length int, input string) (disk, checksum string) {
	a := []byte(input)
	for len(a) < length {
		b := make([]byte, 0, len(a))
		for i := len(a) - 1; i >= 0; i-- {
			switch c := a[i]; c {
			case '0':
				b = append(b, '1')
			case '1':
				b = append(b, '0')
			default:
				panic(c)
			}
		}
		a = append(append(a, '0'), b...)
	}
	a = a[:length]
	disk = string(a)
	for {
		for i := 0; i < len(a)/2; i++ {
			if a[2*i] == a[2*i+1] {
				a[i] = '1'
			} else {
				a[i] = '0'
			}
		}
		a = a[:len(a)/2]
		if len(a)%2 == 1 {
			break
		}
	}
	checksum = string(a)
	return
}
