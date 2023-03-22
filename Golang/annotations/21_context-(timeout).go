package main

import (
  "context"
  "fmt"
  "time"
  "net/http"
)

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  req, err := http.NewRequest("GET", "http://example.com", nil)
  if err != nil {
    panic(err)
  }

  req = req.WithContext(ctx)

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    panic(err)
  }

  fmt.Println(resp.Status)
}
