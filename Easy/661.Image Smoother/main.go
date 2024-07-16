package main

import (
	"fmt"
)

func main() {
	res := imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}})
	fmt.Println(res)
}

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	n := len(img[0])
	fmt.Println(m, n)
	res := make([][]int, len(img))
	for i := range res {
		res[i] = make([]int, n)
	}
	fmt.Println(res)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := img[i][j]
			count := 1
			if i > 0 && j > 0 {
				sum += img[i-1][j-1]
				count++
			}
			if i > 0 {
				sum += img[i-1][j]
				count++
			}
			if j < n-1 && i > 0 {
				sum += img[i-1][j+1]
				count++
			}
			if j > 0 {
				sum += img[i][j-1]
				count++
			}
			if j < n-1 {
				sum += img[i][j+1]
				count++
			}
			if j > 0 && i < m-1 {
				sum += img[i+1][j-1]
				count++
			}
			if i < m-1 {
				sum += img[i+1][j]
				count++
			}
			if i < m-1 && j < n-1 {
				sum += img[i+1][j+1]
				count++
			}
			fmt.Println(sum, count)
			res[i][j] = sum / count
		}
	}
	return res
}
