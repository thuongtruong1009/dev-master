package main

type CustomMap[T comparable, V int | string] map[T]int

func main() {
  m := make(CustomMap[int, string])
  m[1] = "Hello World"
}
