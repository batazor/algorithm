package main

import (
  "fmt"
)

// gets the power series of integer a and tuple of square of a
// and cube of a
func powerSeries(a int) (square int, cube int, error error) {
   square = a*a
   cube = square*a
   return square, cube, nil
}

func main() {
  square, cube, err := powerSeries(3)

  if err != nil {
    fmt.Println("Error ", err)
    return
  }

  fmt.Println("Square ", square, "Cube", cube)
}
