package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
	//sliceを使うときはデフォルト値があるので規定値を省略すればデフォルトが設定される。
	//下限が0で上限がスライスの長さとなる。
}