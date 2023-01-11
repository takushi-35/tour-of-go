package main

import "fmt"

func main() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << uint(i) //i分1を２進数で左にずらしている。つまり2の累乗になっている。
  }

  for _, value := range pow {
    //_を使うことでRangeの戻り値を省略できる、今回は1つ目だが2つ目も省略可能
    fmt.Printf("%d\n", value)
  }
}
