package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1, 2} //typeの構造体
  v2 = Vertex{X: 1} // Y:0が暗黙の了解
  v3 = Vertex{} //X:0, Y:0 値を入れないと0になる
  p = &Vertex{1, 2} //タイプの格納ポインタ
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
