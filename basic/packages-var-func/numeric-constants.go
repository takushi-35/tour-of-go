package main


import "fmt"

const (
  //1のbitを左に100ずらして大きな数字を作る
  //つまり1の後に0が100個続く２進数
  Big = 1 << 100
  //ビットを左に99ずらす。
  small = Big >> 99
)

func needInt(x int) int { return x*10 +1}
func needFloat(x float64) float64 {
  return x * 0.1
}

func main() {
  fmt.Println(needInt(small))
  fmt.Println(needFloat(small))
  fmt.Println(needFloat(Big))
}
