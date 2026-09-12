package main

import "fmt"

func main() {
	var count int
	fmt.Printf("Enter the lenght for the Fibonacci series:\t")
	fmt.Scanf("%d", &count)
	fmt.Printf("%v\n", fibonacci(count))
}

func fibonacci(count int) []int {
	fibSeries := make([]int, 0, count)

	for i := 0; i < count; i++ {
		switch i {
		case 0:
			fibSeries = append(fibSeries, 0)
		case 1:
			fibSeries = append(fibSeries, 1)
		default:
			fibSeries = append(fibSeries, fibSeries[i-1]+fibSeries[i-2])
		}
	}

	return fibSeries
}
