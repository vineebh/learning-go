package main

import "fmt"

func avg(){}

func main(){
	// var arr [10] float64
	arr := [...]float64{1.2,2.2,50,2,3}
	slice := arr[:]
	var avg, sum float64
	for _,v := range(slice){
		sum += v
	}
	avg = sum/float64(len(slice))
	fmt.Printf("%f\n",avg)
}
