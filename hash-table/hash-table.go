package main

import (
	"fmt"
)

func main() {
	// Colission handling => Open Addressing & Separate Chaining

	var findIndex = twoSum([]int{11, 7, 9, 2}, 9)
	fmt.Println(findIndex)
}

func twoSum(nums []int, target int) []int {
	var temp map[int]int
	temp = make(map[int]int)

	for i, val := range nums {
		temp[val] = i
		var curr = target - val

		if temp[curr] == 1 {
			return []int{i, temp[curr]}
		}
	}

	return []int{}
}
