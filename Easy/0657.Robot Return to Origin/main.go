package main

import (
	"fmt"
)

func main() {
	fmt.Println(judgeCircle("RRDD"))
}

func judgeCircle(moves string) bool {
	i, j := 0, 0
	for _, val := range moves {
		switch string(val) {
		case "U":
			i += 1
		case "D":
			i -= 1
		case "R":
			j += 1
		case "L":
			j -= 1
		}
	}
	if i == 0 && j == 0 {
		return true
	}
	return false
}
