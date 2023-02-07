package main

import "fmt"

func main(){
  m := map[string] int { "one": 1, "two": 2, "three": 3 }

  if m["one"] == 1 {
    fmt.Println("one is 1")
  }

  if age, isExist := m["one"]; isExist == true{
    fmt.Println("age is", age)
  } else {
    fmt.Println("age is not exist")
  }


  // switch
  number := 1
  switch number {
    case 1:
      fmt.Println("one")
      // fallthrough // continue to next case
      // break // not execute below
      fmt.Println("continue")
    case 2:
      fmt.Println("two")
    default:
      fmt.Println("other")
  }

  //interface
  var i interface{} = 1
  switch i.(type){
    case int:
      fmt.Println("int")
    case float64:
      fmt.Println("float64")
    case string:
      fmt.Println("string")
    default:
      fmt.Println("other")
  }
}
