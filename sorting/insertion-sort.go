package main

import "fmt"

func main() {
	var sortedNumber = insertionSort([]int{7, 5, 2, 9, 1})
	fmt.Println(sortedNumber)
}

func insertionSort(arr []int) []int {
	var arrLen = len(arr) - 1

	for i := 1; i <= arrLen; i++ {
		var position = i - 1
		var temp = arr[i]

		for position >= 0 {
			if arr[position] > temp {
				arr[position+1] = arr[position]
				position--
			} else {
				break
			}
		}

		arr[position+1] = temp
	}

	return arr
}
