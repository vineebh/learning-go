package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Printf("fizz")
		} else if i%5 == 0 {
			fmt.Printf("buzz")
		} else {
			fmt.Printf("%d", i)
		}
		/*
			switch {
			case i%3==0 && i%5==0:
				fmt.Printf("fizzbuzz")
			case i%3==0:
				fmt.Printf("fizz")
			case i%5==0:
				fmt.Printf("buzz")
			default:
				fmt.Printf("%d",i)
			}
		*/
		fmt.Println()
	}
}

func main() {
	fizzbuzz()
}
