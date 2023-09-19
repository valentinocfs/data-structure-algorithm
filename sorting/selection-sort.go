package main

import "fmt"

func main() {
	var sortedNumber = selectionSort([]int{7, 5, 2, 9, 1})
	fmt.Println(sortedNumber)
}

func selectionSort(arr []int) []int {
	var startIndex = 0
	var arrLen = len(arr) - 1

	for startIndex < arrLen {
		var lowestValue = arr[startIndex]
		var lowestValueIndex = startIndex
		var secondIndex = startIndex + 1

		for i := secondIndex; i <= arrLen; i++ {
			if arr[i] < lowestValue {
				lowestValue = arr[i]
				lowestValueIndex = i
			}
		}

		arr[startIndex], arr[lowestValueIndex] = arr[lowestValueIndex], arr[startIndex]

		startIndex++
	}

	return arr
}
