package main

import (
  "fmt"
  "patterns/creational/factory/abstract_factory"
)

func main() {
  adidasFactory := abstract_factory.GetSportsFactory("adidas")
  nikeFactory := abstract_factory.GetSportsFactory("nike")

  adidasShoe := adidasFactory.MakeShoe()
  adidasShort := adidasFactory.MakeShort()

  nikeShoe := nikeFactory.MakeShoe()
  nikeShort := nikeFactory.MakeShort()

  fmt.Printf("adidas shoe: %v", adidasShoe)
  fmt.Printf("adidas short: %v", adidasShort)

  fmt.Printf("nike shoe: %v", nikeShoe)
  fmt.Printf("nike short: %v", nikeShort)
}
