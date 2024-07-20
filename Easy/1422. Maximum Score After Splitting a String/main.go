package main

import (
	"fmt"
)

func main() {
	res := maxScore("00101101010101111110")
	fmt.Println(res)
}

func maxScore(s string) int {
	var maxl, maxr, max int
	for split := 0; split < len(s)-1; split++ {
		for i := 0; i <= split; i++ {
			if s[i] == '0' {
				maxl++
			}
		}
		for i := split + 1; i < len(s); i++ {
			if s[i] == '1' {
				maxr++
			}
		}
		if max < maxl+maxr {
			max = maxl + maxr
		}
		maxl, maxr = 0, 0
	}

	return max
}
