package main

import (
  "fmt"
  "sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func sayHello(){
  fmt.Printf("Hello #%v\n", counter)
  m.RUnlock()
  wg.Done()
}

func increment(){
  counter++
  m.Unlock()
  wg.Done()
}

func main(){
  for i := 0; i < 10; i++ {
    wg.Add(2)
    m.RLock() // prevent other can't change counter -> lock to read
    go sayHello()
    m.Lock()
    go increment()
  }
  wg.Wait()
}
