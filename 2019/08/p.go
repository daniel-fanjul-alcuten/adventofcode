package p

import "strings"

func P1(w, h int, image string) int {
	minlayer, mincount := "", -1
	for i := 0; i < len(image); i += w * h {
		layer := image[i : i+w*h]
		count := strings.Count(layer, "0")
		if mincount < 0 || count < mincount {
			minlayer, mincount = layer, count
		}
	}
	c1 := strings.Count(minlayer, "1")
	c2 := strings.Count(minlayer, "2")
	return c1 * c2
}

func P2(w, h int, image string) (s string) {
	nlayers := len(image) / w / h
	layer := func(i int) string {
		return image[i*w*h : (i+1)*w*h]
	}
	for i := 0; i < w*h; i++ {
		found := false
		for l := 0; l < nlayers; l++ {
			switch p := layer(l)[i : i+1]; p {
			case "0", "1":
				s += p
				found = true
			case "2":
			default:
				panic(p)
			}
			if found {
				break
			}
		}
		if !found {
			s += "2"
		}
	}
	return
}
