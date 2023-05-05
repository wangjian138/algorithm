package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	sorted := append(quickSort(less), pivot)
	sorted = append(sorted, quickSort(greater)...)

	return sorted
}

func main() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 10}
	fmt.Println("Unsorted array:", arr)

	sorted := quickSort(arr)
	fmt.Println("Sorted array:", sorted)
}
