package main

import "fmt"

type Num interface {
  int | float64
}

func add[T Num](a T, b T) T {
  return a + b
}

func main() {
  fmt.Println(add(1, 2))
  fmt.Println(add(1.1, 2.2))
}
