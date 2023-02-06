package main

import (
  "fmt"
  "reflect"
)

type User struct {
  Id   int    `json:"id"`
  Name string `json:"name"`
}

func main() {
  u := User{1, "Tom"}
  t := reflect.TypeOf(u)
  field, _ := t.FieldByName("Name")
  fmt.Println(field)

  for i := 0; i < t.NumField(); i++ {
    f := t.Field(i) //split with white space
    fmt.Printf("%s", f.Tag.Get("json"))
  }
}
