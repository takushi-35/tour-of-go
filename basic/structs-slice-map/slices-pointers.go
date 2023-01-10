package main

import "fmt"

func main() {
	names := [4] string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	//slicesの要素を変更すると元の配列の値も変更される。
	//arrayではある変数に配列を格納後、その変数の値を上書きしても元の配列に影響はない。
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}