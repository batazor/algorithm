package main

import "fmt"
import "math/big"

func main() {
  c := make(chan *big.Int, 2)
  c <- big.NewInt(1)
  c <- big.NewInt(0)

  for {
    n := <- c
    fmt.Println(n)
    c <- big.NewInt(0).Add(n, <- c)
    c <- n
  }
}
