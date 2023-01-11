package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

func main() {
  //mapsはキー(data)を利用して要素を指定できる。
  //例えばyear=365, week=5などキーに数字を紐付けたりできる
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex {
    40.68422, -74.39967,
  }
  fmt.Println(m["Bell Labs"])
}
