package main

import(
  "fmt"
  "time"
  "sync"
)

func count(name string, wg *sync.WaitGroup){
  for i := 1; i<=5; i++ {
    fmt.Println(name, i)
    time.Sleep(time.Second)
  }
  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  wg.Add(2) // wait 2 goroutines

  go count("sheep", &wg)
  go count("fish", &wg)

  wg.Wait()
  fmt.Println("Done")
}
