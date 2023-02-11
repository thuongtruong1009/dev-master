package main

import "fmt"

func main(){
    // defer
    a := "Hello"
    defer fmt.Println(a)  // defer will cache snapshot of a in stack
    a = "Hi"  // a change will not affect to defer

    // panic recover
    fmt.Println("start")
    panic("something bad happened")
    str := recover()
    fmt.Println(str)
    fmt.Println("end")
}
