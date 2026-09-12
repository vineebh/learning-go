package main

import "fmt"

func main() {
	unsorted := [...]int{1, 2, 3, 5, 2, 1, 0}
	fmt.Println(unsorted)
	sorted := bubbleSort(unsorted[:])
	fmt.Println(sorted)
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
