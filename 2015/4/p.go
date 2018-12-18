package p

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func P1(input string) (output int) {
	i, h, hh := 1, md5.New(), []byte(nil)
	for {
		h.Reset()
		if _, err := fmt.Fprintf(h, "%v%d", input, i); err != nil {
			panic(err)
		}
		hh = h.Sum(hh[:0])
		h := hex.EncodeToString(hh)
		if strings.HasPrefix(h, "00000") {
			return i
		}
		i++
	}
}

func P2(input string) (output int) {
	i, h, hh := 1, md5.New(), []byte(nil)
	for {
		h.Reset()
		if _, err := fmt.Fprintf(h, "%v%d", input, i); err != nil {
			panic(err)
		}
		hh = h.Sum(hh[:0])
		h := hex.EncodeToString(hh)
		if strings.HasPrefix(h, "000000") {
			return i
		}
		i++
	}
}
