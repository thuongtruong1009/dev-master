package main

import "fmt"

func main(){
  //1. Integer types
  // Signed Integer: int8, int16, int32, int64 -> positive
  // Unsigned Integer: uint8, uint16, uint32, uint64 -> negative
  var a int8 = 10 // 1010s
  var b int8 = 3  // 0011
  fmt.Println(a & b) // -> AND (same bit 1: 1 else: 0) -> 0010
  fmt.Println(a | b) // -> OR (same bit 0: 0 else: 1) -> 1011
  fmt.Println(a ^ b) // -> XOR (same bit: 0 else: 1) -> 1001
  fmt.Println(a &^ b) // -> NAND -> 1000
  fmt.Println(a << 2) // -> Left shift  -> 0010 1000
  fmt.Println(a >> 2) // -> Right shift -> 0000 0010

  //2. Floating point types
  // float32, float64
  var c float32 = 10.5e2
  var d float32 = 3.5E2
  fmt.Println(c + d)

  //3. Complex types
  // complex64, complex128
  var e complex64 = 1 + 2i
  var f complex64 = complex(3, 4)
  fmt.Println(e + f, real(e + f), imag(e + f))

  //4. Boolean types
  // bool
  var g bool = true
  var h bool = false
  fmt.Println(g && h)

  //5. String types
  // string
  var i string = "Hello"
  var j string = "World"
  fmt.Println(i + j)

  //6. Byte types (alias for uint8 -> unsigned symbol)
  // byte
  var k byte = 'a'
  var l byte = 'b'
  fmt.Println(k + l)

  // combine string and byte
  var m string = "Hello"
  var n byte = 'a'
  fmt.Println(m + string(n))

  //7. Rune types (alias for int32 -> signed symbol)
  // rune
  var o rune = 'a'
  var p rune = 'b'
  fmt.Println(o, p)


  //8. Array
  var r []byte = []byte{'a', 'b', 'c'}
  var s []byte = []byte(m)
  fmt.Println(r, s)

}
