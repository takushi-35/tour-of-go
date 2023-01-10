package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	//sliceのゼロ値はnil。nil sliceは０の長さと容量を持っている

	if s == nil {
		fmt.Println("nil!")
	}

}