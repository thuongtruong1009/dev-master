package main

import "fmt"

func fib(n int) int {
  if n <= 1 {
    return n
  }
  return fib(n-1) + fib(n-2)
}

func worker(jobs <-chan int, results chan<- int) {
  for n := range jobs {
    results <- fib(n)
  }
}

func main() {
  number := 1000
  numberOfWorker := 5 // number of threads in CPU
  jobs := make(chan int, number)
  results := make(chan int, number)

  for i := 0; i < numberOfWorker; i++ {
    go worker(jobs, results)
  }

  for i := 0; i < number; i++ {
    jobs <- i
  }

  close(jobs)

  for j := 0; j < number; j++ {
    fmt.Println(<-results)
  }
}
