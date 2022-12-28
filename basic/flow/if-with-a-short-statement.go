package main

import (
  "fmt"
  "math"
)


func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim { //pow関数は冪乗計算。xのn乗
    return v
  }
  return lim
}

func main () {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
