package main

import (
  "fmt"
  "net/http"
  "os"
  "os/signal"
  "sync"
  "time"
  "context"
)

func main() {
  var wg sync.WaitGroup

  wg.Add(1)

  go func() {
    defer wg.Done()

    time.Sleep(5 * time.Second)
  }()

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)

  fmt.Println("server started!")
  srv := http.Server{Addr: ":8080"}

  go func() {
    <-c
    fmt.Println("server shutting down...")
    go func(){
      for {
        fmt.Println("waiting for goroutines to finish...")
        time.Sleep(1 * time.Second)
      }
    }()
    wg.Wait()

    if err:= srv.Shutdown(context.Background()); err != nil {
      fmt.Println("server shutdown error:", err)
    }
  }()
  fmt.Println(srv.ListenAndServe())
}
