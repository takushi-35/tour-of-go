package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

//引数をインターフェスMで返す
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	
	i = &T{"Hello"} //iにTのポインタ経由でHelloを代入
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}

func describe(i I) {
	//代入された関数をその型と一緒に返すさらに、その関数名と渡した引数名を返却
	fmt.Printf("(%v, %T)\n", i, i)
}