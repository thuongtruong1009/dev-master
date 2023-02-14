package main

import "fmt"

func findMax(a, b int) (max *int) {
  if(a > b) {
    max = &a
  }else{
    max = &b
  }
  return
}

func findSum(s string, a []int, b ...int) (sum int){
  for _, v := range b {
    sum += v
  }
  fmt.Println(s, a)
  return
}

func calDevide(a, b int) (int, error){
  result := 0
  if(b == 0) {
    return 0, fmt.Errorf("b can not be zero")
  }
  result = a / b
  return result, nil
}

func nested(){
  for i := 0; i < 5; i++ {
    func (i int){
      fmt.Println(i)
    }(i)
  }
}

// implement method interface
type greeter struct{
  greeting string
  name string
}

func ( g greeter) greetNotChange(){
  fmt.Println(g.greeting, g.name)
  g.name = "Golang"
}
func ( g *greeter) greetChange(){
  fmt.Println(g.greeting, g.name)
  g.name = "Golang"
}

func main() {
  a :=10
  b :=20

  maxPointer := findMax(a, b)
  fmt.Println(*maxPointer)
  fmt.Println(maxPointer)

  sum := findSum("Hello", []int{1, 2, 3, 4, 5}, 1,2,3,4,5)
  fmt.Println(sum)

  res, err := calDevide(10, 0)
  if(err != nil) {
    fmt.Println(err)
  }
  fmt.Println(res)

  nested()

  g := greeter{greeting: "Hello", name: "World"}
  g.greetNotChange()
  fmt.Println(g.name)

  g.greetChange()
  fmt.Println(g.name)
}
