package main

import "fmt"

type Animal interface {
  Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
  return "Woof!"
}

func main() {
  var d1 Animal = Dog{}
  fmt.Println(d1.Speak())
}
