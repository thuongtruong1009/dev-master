package main

import "fmt"

func print(c map[string] string){
  for key, value := range c {
    fmt.Println("hex code for ", key, "is", value)
  }
}

func main(){
  // colors := make(map[string] string)
  colors := map[string] string{
    "red": "#ff0000",
    "green": "#4bf745",
  }

  colors["white"] = "#ffffff"

  delete(colors, "red")

  print(colors)
}
