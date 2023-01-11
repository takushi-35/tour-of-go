package main

import "fmt"

func main() {
  i, j := 42, 2701

  p := &i //iが格納されているPointer
  fmt.Println(*p) //ポインタを通して変数iの値を出力
  *p = 21 //ピンタを通してiへ値を代入している。
  fmt.Println(i)

  p = &j

  *p = *p /37
  fmt.Println(j)
}
