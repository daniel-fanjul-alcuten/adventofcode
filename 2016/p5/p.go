package p

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func P1(doorID string) (passwd string) {
	h := md5.New()
	b := []byte(nil)
	for i := 0; len(passwd) < 8; i++ {
		h.Reset()
		if _, err := fmt.Fprintf(h, "%s%d", doorID, i); err != nil {
			panic(err)
		}
		b = h.Sum(b[:0])
		h := hex.EncodeToString(b)
		if strings.HasPrefix(h, "00000") {
			passwd += h[5:6]
		}
	}
	return
}

func P2(doorID string) (passwd string) {
	h := md5.New()
	b := []byte(nil)
	m := map[int]string{}
	for i := 0; len(m) < 8; i++ {
		h.Reset()
		if _, err := fmt.Fprintf(h, "%s%d", doorID, i); err != nil {
			panic(err)
		}
		b = h.Sum(b[:0])
		h := hex.EncodeToString(b)
		if strings.HasPrefix(h, "00000") {
			pos := h[5]
			if pos >= '0' && pos <= '7' {
				key := int(pos) - '0'
				if _, ok := m[key]; !ok {
					m[key] = h[6:7]
				}
			}
		}
	}
	for i := 0; i < 8; i++ {
		passwd += m[i]
	}
	return
}
