package main

import "fmt"

func for_loop(){
	for i := 0;i<10;i++{
		//fmt.Println(i)
		fmt.Printf("%d\n",i+1)
	}
}

func goto_loop(){
	var i int= 0
	here:
	i += 1
	fmt.Println(i)
	if i<10{
		goto here
	}
}

func array_loop(){
	var arr [10] int
	for i := range(10){
		arr[i]=i+1
	}
	/*for i := range(10){
		fmt.Println(arr[i])
	}*/
	fmt.Printf("%v\n",arr)
}

func main(){
	// for_loop()
	// goto_loop()
	array_loop()
}
