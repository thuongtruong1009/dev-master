package main

import (
  "fmt"
  "time"
  "strconv"
)

// var zero =
//   `███
//    █ █
//    █ █
//    █ █
//    ███`

// var one =
//   `██
//     █
//     █
//     █
//    ███`

// var two =
//   `███
//      █
//    ███
//    █
//    ███`

// var three =
//   `███
//      █
//    ███
//      █
//    ███`

// var four =
//   `█ █
//    █ █
//    ███
//      █
//      █`

// var five =
//   `███
//    █
//    ███
//      █
//    ███`

// var six =
//   `███
//    █
//    ███
//    █ █
//    ███`

// var seven =
// `███
//     █
//     █
//     █
//     █`

// var eight =
//   `███
//    █ █
//    ███
//    █ █
//    ███`

// var nine =
//   `███
//    █ █
//    ███
//      █
//    ███`

// var colon =
//   `
//     ░

//     ░
//       `

func clearConsole() {
  fmt.Print("\033[H\033[2J")
}

func formatTime(t time.Time) string {
  hour := t.Hour()
  minute := t.Minute()
  second := t.Second()

  hourStr := strconv.Itoa(hour)
  minuteStr := strconv.Itoa(minute)
  secondStr := strconv.Itoa(second)

  if hour < 10 {
    hourStr = fmt.Sprintf("0%s", hourStr)
  }
  if minute < 10 {
    minuteStr = fmt.Sprintf("0%s", minuteStr)
  }
  if second < 10 {
    secondStr = fmt.Sprintf("0%s", secondStr)
  }

  return fmt.Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
}

func main() {
  msgTime := make(chan time.Time)

  go func() {
    for {
      time.Sleep(1 * time.Second)
      msgTime <- time.Now()
    }
  }()

  for {
    select {
    case t := <-msgTime:
      clearConsole()
      fmt.Println(formatTime(t))
    }
  }
}
