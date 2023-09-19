package main

import "fmt"

func main() {
	var sortedNumber = bubbleSort([]int{7, 5, 2, 13, 1})
	fmt.Println(sortedNumber)
}

func bubbleSort(arr []int) []int {
	var arrlen = len(arr) - 1

	for arrlen > 0 {
		var i = 0

		for i < arrlen {
			var nextIndex = i + 1

			if arr[i] > arr[nextIndex] {
				arr[i], arr[nextIndex] = arr[nextIndex], arr[i]
			}

			i += 1
		}

		arrlen -= 1
	}

	return arr
}
