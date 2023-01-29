package main

import "fmt"

func main(){
  // declare static length array
  array1 := [5] int {1, 2, 3, 4, 5}
  fmt.Println(array1)

  // declare dynamic length array
  array2 := [...] int {1, 2, 3, 4, 5}
  fmt.Println(array2)

  //not init value -> value is 0
  var array3 [3] int
  fmt.Println(array3)
  array3[0] = 1
  array3[1] = 2
  array3[2] = 3
  fmt.Println(array3)

  // pointer array
  array4 := [3] int { 2, 5, 10}
  arrayPointer := &array4
  fmt.Println(arrayPointer)

  //root change -> all change
  arrayPointer[0] = 100
  fmt.Println(array4)

  // slice
  slice := [] int {1, 2, 3, 4, 5}
  slice1 := slice[1:3]
  slice2 := slice[:4]
  slice3 := slice[2:]
  fmt.Println(slice1)
  fmt.Println(slice2)
  fmt.Println(slice3)

  // declare when know length
  slice4 := make([] int, 5)
  fmt.Println(slice4)
  slice5 := make([] int, 5, 10)
  fmt.Println(len(slice5), cap(slice5))

  // append
  slice6 := [] int {1, 2, 3, 4, 5}
  slice6 = append(slice6, 6, 7, 8)
  fmt.Println(slice6)
}





