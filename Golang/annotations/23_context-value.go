package main

import (
  "fmt"
  "context"
)

// send value to context
func enrichContext(ctx context.Context) context.Context {
  return context.WithValue(ctx, "request-id", "12345")
}

// get value from context
func doSomethingCool(ctx context.Context) {
  rID := ctx.Value("request-id")
  fmt.Println(rID)
  for {
    select {
    case <-ctx.Done():
      fmt.Println("timed out")
      return
    default:
      fmt.Println("working...")
    }
    time.Sleep(500 * time.Millisecond)
  }
}

func main() {
  fmt.Println("Go context tutorial")
  ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)

  defer cancel()

  ctx = enrichContext(ctx)
  doSomethingCool(ctx)

  select {
  case <-ctx.Done():
    fmt.Println("Context is exceeded the deadline")
  }
  time.Sleep(2 * time.Second)
}
