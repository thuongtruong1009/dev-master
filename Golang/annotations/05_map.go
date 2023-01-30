package main

import "fmt"

func main() {
  studentName := make(map[string] int)
  studentName = map[string] int {
    "John": 10,
    "Peter": 20,
    "Sam": 30,
  }
  fmt.Println(studentName)

  // Add new element
  studentName["Raj"] = 40
  fmt.Println(studentName)

  // Delete element
  delete(studentName, "Sam")
  fmt.Println(studentName)

  // Check if element exists
  if _, ok := studentName["Sam"]; ok {
    fmt.Println("Sam exists")
  } else {
    fmt.Println("Sam does not exist")
  }
}
