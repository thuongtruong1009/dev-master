package main

import (
  "fmt"
  "time"
)

func count(name string){
  for i := 1; i<=5; i++ {
   fmt.Println(name, i)
    time.Sleep(time.Second)
  }
}

func main() {
  go count("sheep")
  go count("fish")

  time.Sleep(time.Second * 10)
}
