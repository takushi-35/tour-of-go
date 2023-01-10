package main

import "fmt"

func main() {
	//make関数でsliceを動的サイズで作成することができる
	//一つ目の引数に型名、2つ目にlength,3つ目にcapacityを指定できる。
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}