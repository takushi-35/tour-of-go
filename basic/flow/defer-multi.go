package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++{
		//deferが複数ある場合、LIFO(Last-in-first-out)で実行される
		defer fmt.Println(i)
	}
	fmt.Println("done")
}