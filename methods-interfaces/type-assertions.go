package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string) //iにstringの型を保持していればsにiの値を代入し、okにtrueを代入する
	fmt.Println(s, ok)

	f, ok := i.(float64) //iにfloatの型を保持していないためfには0が代入される。okにはfalseが代入される
	fmt.Println(f, ok)


	f = i.(float64) //panicになる
	fmt.Println(f)



}