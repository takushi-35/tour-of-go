package main

import "fmt"

func main() {
  //appendでスライスへ新しい要素を追加することができる。
  //ここではまずからの配列を用意
  var s []int
  printSlice(s)


  //nilだった配列に0を追加
  s = append(s, 0)
  printSlice(s)

  //次に1も配列に追加。(capも拡張されている)
  s = append(s, 1)
  printSlice(s)

  //配列の長さとcapが拡張。容量の拡張は足りない分の保管ではなく、元の容量の倍分が拡張される
  s = append(s, 2, 3, 4)
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
