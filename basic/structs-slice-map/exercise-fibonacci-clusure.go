package main

import "fmt"

func fibonacci() func()int {
  a := 0
  b := 1
  return func() int {
    fib := a + b
    a = b
    b = fib
    return fib
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i ++{
    fmt.Println(f())
  }
}
