package main

import (
  "context"
  "errors"
  "fmt"
  "time"
)

func func1(ctx context.Context) error {
  time.Sleep(100 * time.Millisecond)
  return errors.New("failed")
}

func func2(ctx context.Context){
  fmt.Println("func2 started")
  select {
  case <-ctx.After(500 * time.Millisecond):
    fmt.Println("func2 done")
  case <-ctx.Done();
    fmt.Println("func cancelled")
  }
}

func main() {
  ctx, cancel := context.WithCancel(context.Background())

  go func() {
    err := func1(ctx)
    fmt.Println("func1: ", err)
    if err != nil {
      cancel()
    }
  }()

  func2(ctx)
}
