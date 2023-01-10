package main

import "fmt"

func main() {
	var a [2] string
	a[0] = "Hello"
	a[1] = "World!"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6] int {2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	b := a[1]
	fmt.Println(b)

	//arrayではある変数に配列を格納後、その変数の値を上書きしても元の配列に影響はない。
	b = "Go"
	fmt.Println(b)
	fmt.Println(a)

}