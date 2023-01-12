package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
    word := strings.Fields(s)
    res := make(map[string]int)
    for i := 0; i < len(word); i++ {
      res[word(i)] = res[word(i)] + 1
    }
    return res
}

func main() {
	wc.Test(WordCount)
}
