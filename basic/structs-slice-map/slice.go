package main

import "fmt"

func main() {
	primes := [6] int {2, 3, 5, 7, 11, 13}

	var s []int = primes[1 : 4] //配列の2つ目から4つ目まで。最初の指定の値は含むが、後ろの配列分は含まない。
	
	fmt.Println(s)
	fmt.Println(primes[0:3])

}