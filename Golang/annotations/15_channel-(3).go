package main

import (
  "fmt"
  "sync"
)

func main(){
  var wg = sync.WaitGroup{}
  ch1 := make(chan int, 50)
  ch2 := make(chan int, 50)

  wg.Add(2)

  //receive channel
  go func (){
    for {
      select {
        case i, ok := <-ch1:
          if ok {
            fmt.Printf("Received from channel 1: %v\n", i)
          } else {
            break
          }
        case j, ok := <-ch2:
          if ok {
            fmt.Printf("Received from channel 2: %v\n", j)
          } else {
            break
          }
        default:
          fmt.Println("No data received")
          break
      }
    }
    wg.Done()
  }()

  //send channel
  go func (){
    ch1 <- 10
    close(ch1)

    ch2 <- 20
    close(ch2)
    wg.Done()
  }()

  wg.Wait()
}
