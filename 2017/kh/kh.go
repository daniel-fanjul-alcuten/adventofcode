package kh

import "encoding/hex"

func KnotHash(bb []byte) (h []byte, x string) {
	bytes, position, skip := [256]byte{}, 0, 0
	for i := range bytes {
		bytes[i] = byte(i)
	}
	for _ = range [64]int{} {
		for _, d := range append(bb, 17, 31, 73, 47, 23) {
			i, j := position, (position+int(d)-1+len(bytes))%len(bytes)
			for t := d; t > 1; t -= 2 {
				bytes[i], bytes[j] = bytes[j], bytes[i]
				i, j = (i+1)%len(bytes), (j-1+len(bytes))%len(bytes)
			}
			position = (position + int(d) + skip) % len(bytes)
			skip += 1
		}
	}
	for i, j := 0, 0; i < len(bytes); i, j = i+16, j+1 {
		t := bytes[i]
		for j := i + 1; j < i+16; j++ {
			t ^= bytes[j]
		}
		h = append(h, t)
	}
	x = hex.EncodeToString(h[:])
	return
}
