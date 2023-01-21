package main

import "fmt"

const a = 10

// iota start at 0 and increment by 1
const (
  red = iota
  green
  blue
)

// custom iota
const(
  _ = 1 << (10 * iota)
  KB
  MB
  GB
)

func main(){
  const i int16 = 16
  fmt.Println(i)

  // shadow constant
  const a = 12
  fmt.Println(a)

  fmt.Println(red, green, blue)

  fmt.Println(KB, MB, GB)
}

