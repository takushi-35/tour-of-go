package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"])

  //要素の更新
  m["Answer"] = 48
  fmt.Println("The value", m["Answer"])

  //要素の削除
  delete(m, "Answer")
  fmt.Println("The Value", m["Answer"])

  v, ok := m["Answer"] //仮にm(key)に要素があれば2つ目の変数(ok)にtrueが格納される
  fmt.Println("The value:", v, "Present?", ok)
}
