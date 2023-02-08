package main

import "fmt"

func main(){
    // for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // while loop
    i := 0
    for i < 10 {
      if(i == 4){
        i++
        continue
      }
      fmt.Println(i)
      i++
    }

    // infinite loop
    for {
        fmt.Println("infinite loop")
        break
    }

    // label loop
    Loop:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            if j == 2 {
                break Loop
            }
            fmt.Println(i, j)
        }
    }

    // range loop
    arr := []int{1, 2, 3, 4, 5}
    for i, v := range arr {
        fmt.Println(i, v)
    }
}
