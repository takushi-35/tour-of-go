package main

import (
	"fmt"
)

func main () {
	defer fmt.Println("World") 
	//deferへの関数の引数はこの時点で渡されるが実行は呼び出し元の関数が実行されるときに行われる

	fmt.Println("Hello, ")
}