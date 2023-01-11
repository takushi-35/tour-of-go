package main

import (
  "fmt"
  "strings"
)

func main() {
  board := [][]string{
    //３目並べ
    []string{"_", "_", "_"  },
    []string{"_", "_", "_"  },
    []string{"_", "_", "_"  },
  }

  //配列に格納。一つの目が行、2つ目が列
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
}
