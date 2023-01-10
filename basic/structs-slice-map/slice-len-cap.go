package main

import "fmt"

func main() {
	
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//sliceの要素を0に削る。元の要素数(cap)には影響なし
	s = s[:0]
	printSlice(s)

	//スライスの長さを拡張してみる。元の要素数(cap)には影響なし
	s = s[:4]
	printSlice(s)

	//スライスの最初の2つの要素を落とす。元の要素数(cap)にも影響あり
	s = s[2:]
	printSlice(s)

	s = s[0: 3]
	printSlice(s)

}

func printSlice(s []int) {
	//len(length)は配列の要素数
	//cap(capacity)はスライスの最初の要素から数えて、元となる配列の要素数
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
