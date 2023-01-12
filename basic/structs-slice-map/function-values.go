package main


import (
  "fmt"
  "math"
)

//変数(関数)に引数3,4を入れてその結果を返す変数
func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

func main() {
  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*y + y*y)
  }
  fmt.Println(hypot(5, 12))

  //goでは関数も変数にできる。
  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))
}
