package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	res := reverseVowels("helliokjjuofa")
	fmt.Println(res)
}

func reverseVowels(s string) string {

	r := []rune(s)

	ls := len(s) - 1

	vowels := []string{"a", "e", "i", "o", "u"}

	for i, val := range r {
		val := string(val)
		if i >= ls {
			break
		}
		if slices.Contains(vowels, val) || slices.Contains(vowels, strings.ToLower(val)) {
			for j := ls; j > 0; j-- {
				if slices.Contains(vowels, string(r[j])) || slices.Contains(vowels, strings.ToLower(string(r[j]))) {
					r[i], r[j] = r[j], r[i]
					ls = j - 1
					break
				}
			}
		}
	}

	return string(r)
}
