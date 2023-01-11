package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
  for i, v := range pow {
    //sliceやmapsを一つ一つ反復処理するために利用される。
    //またその時に2つの変数を返す。一つ目が、indexで2つ目がindexの場所の要素(値)
    fmt.Printf("2**%d = %d\n", i, v)
    fmt.Println(i, v)
  }
}
