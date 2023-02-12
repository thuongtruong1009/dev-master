package main

import "fmt"

type myStruct struct {
  number int
}

func main() {
  var a int = 12
  var b *int = &a
  fmt.Println(a, *b)

  a = 24
  fmt.Println(a, *b)

  *b = 36
  fmt.Println(a, *b)

  //struct
  var c *myStruct = new(myStruct) // &myStruct{number: 12}
  fmt.Println(c.number)

  (*c).number = 24
  fmt.Println((*c).number)
}
