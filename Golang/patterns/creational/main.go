package main

import (
  "fmt"
  "time"
  "patterns/creational/singleton"
  "patterns/creational/builder"
)

func main() {
  // Singleton ------------------------------------------
  for i := 0; i < 10; i++ {
    go func() {
      fmt.Printf("%p\n", singleton.GetInstance())
    }()
  }
  time.Sleep(10 * time.Second)

  // Builder------------------------------------------
  normalBuilder := builder.GetBuilder("normal")
  iglooBuilder := builder.GetBuilder("igloo")

  director := builder.NewDirector(normalBuilder)

  normalHouse := director.BuildHouse()
  fmt.Printf("Normal house door type: %s\n", normalHouse.GetDoorType())
  fmt.Printf("Normal house window type: %s\n", normalHouse.GetWindowType())
  fmt.Printf("Normal house num floor: %d\n", normalHouse.GetNumFloor())

  director.SetBuilder(iglooBuilder)

  iglooHouse := director.BuildHouse()
  fmt.Printf("\nIgloo house door type: %s\n", iglooHouse.GetDoorType())
  fmt.Printf("Igloo house window type: %s\n", iglooHouse.GetWindowType())
  fmt.Printf("Igloo house num floor: %d", iglooHouse.GetNumFloor())
}
