package main

import (
  "fmt"
  "sync"
)

func main(){
  var wg = sync.WaitGroup{}
  ch := make(chan int, 50)
  wg.Add(2)

  //receive channel
  go func (ch <-chan int){
    for {
      if i, ok := <-ch; ok {
        fmt.Println("Received from channel", i)
      } else {
        break
      }
    }
    wg.Done()
  }(ch)

  //send channel
  go func (ch chan<- int){
    ch <- 10
    ch <- 20
    ch <- 30
    close(ch) // stop send and receive
    wg.Done()
  }(ch)

  wg.Wait()
}
