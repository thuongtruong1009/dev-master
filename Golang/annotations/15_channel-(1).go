package main

import (
  "fmt"
  "sync"
)

func main(){
  var wg = sync.WaitGroup{}
  ch := make(chan int)
  wg.Add(2)

  //receive channel
  go func (ch <-chan int){
    i := <- ch
    fmt.Println("Received from channel", i)
    wg.Done()
  }(ch)

  //send channel
  go func (ch chan<- int){
    a := 10
    ch <- a
    a = 20 //not effect channel because channel got snapshot of a
    wg.Done()
  }(ch)

  wg.Wait()
}
