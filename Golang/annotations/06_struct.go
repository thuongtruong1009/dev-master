package main

import "fmt"

type Student struct {
  name string
  age int
  isMale bool
  subjects []string
}

// extends struct
type Student2 struct {
  Student
  grade int
}

func main() {
  // When create new struct the same type, can remove key
  student := Student{
    name: "John",
    age: 10,
    isMale: true,
    subjects: []string{"Math", "English"},
  }
  fmt.Println(student)

  student.name = "Doe"
  fmt.Println(student.name)


  // Declare and assign value
  studentNew1 := struct {name string; age int}{"John", 10}
  fmt.Println(studentNew1)

  // Copy struct
  studentNew2 := studentNew1
  studentNew2.name = "Doe"
  fmt.Println(studentNew2)

  // extends struct
  studentExtend := Student{}
  studentExtend.name = "John"
  studentExtend.age = 10
  studentExtend.isMale = true
  studentExtend.subjects = []string{"Math", "English"}
  studentExtend.grade = 10
}

