package main

import "fmt"

func generatePipeline(numbers []int) <-chan int {
  out := make(chan int)
  go func() {
    for _, n := range numbers {
      out <- n
    }
    close(out)
  }()
  return out
}

func fanOut(in <-chan int) <-chan int{
  out := make(chan int)
  go func() {
    for n := range in {
      out <- n*n
    }
    close(out)
  }()
  return out
}

func fanIn(inputChannel ...<-chan int) <-chan int {
  in := make(chan int)
  go func() {
    for _, c := range inputChannel {
      for n := range c {
        in <- n
      }
    }
  }()
  return in
}

func main() {
  randomNumbers := []int{}
  for i := 1; i <= 1000; i++ {
    randomNumbers = append(randomNumbers,i)
  }

  //generate pipeline
  inputChan := generatePipeline(randomNumbers)

  //fan-out
  c1 := fanOut(inputChan)
  c2 := fanOut(inputChan)
  c3 := fanOut(inputChan)
  c4 := fanOut(inputChan)

  //fan-in
  c := fanIn(c1,c2,c3,c4)

  sum := 0
  for i := 0; i < len(randomNumbers); i++ {
    // sum += randomNumbers[i] * randomNumbers[i]
    sum += <-c
  }
  fmt.Printf("Total sum of square: %d", sum)
}
