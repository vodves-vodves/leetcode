package main

import (
	"fmt"
	"sort"
)

func main() {
	res := maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4)
	fmt.Println(res)

}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	res := 0
	fmt.Println(boxTypes)
	for _, boxType := range boxTypes {
		if truckSize > 0 {
			if truckSize >= boxType[0] {
				res += boxType[0] * boxType[1]
				truckSize -= boxType[0]
			} else {
				res += truckSize * boxType[1]
				truckSize -= truckSize
			}
		}
	}
	return res
}
