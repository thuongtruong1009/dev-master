package main

import (
  "fmt"
  "golang.org/x/exp/constraints"
)

func mapValues[T any](values []T, mf func(T) T) []T {
  var newValues []int
  for _, v := range values {
    newValue := mf(v)
    newValues = append(newValues, newValue)
  }
  return newValues
}

func main() {
  result := mapValues([]int{1, 2, 3}, func(n int) int {
    return n * 2
  })

  fmt.Println(result)
}
