package main

import "fmt"

func avg(slice []float64) (avg float64) {
	sum := 0.0
	if len(slice) == 0 {
		avg = 0
	} else {
		for _, i := range slice {
			sum += i
		}
		avg = sum / float64(len(slice))
	}
	return avg
}

func main() {
	arr := [...]float64{1.1, 2.2, 3.3}
	average := avg(arr[:])
	fmt.Printf("%.2f\n", average)
}
