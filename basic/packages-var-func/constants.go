package main

import "fmt"

const Pi = 3.14 //constは定数。定数は文字、文字列、boolean、数値(numeric)で使える

func main() {
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
