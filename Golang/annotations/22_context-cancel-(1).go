package main

import (
  "fmt"
  "context"
)

func main() {
  taskfn := func(ctx context.Context) <-chan int {
    ch := make(chan int)
    n := 1

    go func()  {
      for {
        select {
        case <-ctx.Done():
          fmt.Println("taskfn: Done")
          return
        case ch <- n:
          n++
        }
      }
    }()
  }
  return ch
}

ctx, cancel := context.WithContext(ctx.Background())

defer cancel()

for n := range taskfn(ctx){
  fmt.Println(n)
  if n == 5 {
    break
  }
}
