package main

import "fmt"

type bot interface{
  getGreeting() string
}

type engBot struct{}

type spanBot struct{}

func (engBot) getGreeting() string{
  return "Hello"
}

func (spanBot) getGreeting() string{
  return "Hola!"
}

func print(b bot){
  fmt.Println(b.getGreeting())
}

func main(){
  en := engBot{}
  sp := spanBot{}

  print(en)
  print(sp)
}
