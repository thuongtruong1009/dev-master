package main

import "fmt"

// Access modifiers scope
var export1 string = "Other modules not access -> private"
var Export2 string = "Other modules can access -> public"

//Global variables
var (
  i1 int = 1
  i2 float64 = 2.5
  i3 string = "Hello"
  i4 bool = true
)

func test(){
  fmt.Println(i1, i2, i3, i4)
}

func main() {
  //overwrite i1 value -> shadow
  var i1 int
  i1 = 10
  fmt.Println(i1)

  //shorthand for init and assign single variable
  i2 := 20.5
  fmt.Printf("%v %T \n", i2, i2)

  var i3 bool
  i3 = (1 == 1)
  fmt.Println(i3)

  //shorthand for init and assign multiple variables
  i4, i5 := "Hello", true
  fmt.Println(i4, i5)

  test()
}
