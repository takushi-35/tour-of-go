package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//ポインタレシーバを使いレシーバがさす変数を変更する。レシーバの変数を更新するには*を使う必要がある。
//ここではVertexを更新している。*をつけなければメイン関数のところの初期で宣言した、3,4のまま使われることになる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
