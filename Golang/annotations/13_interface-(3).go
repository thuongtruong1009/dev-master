package main

import (
  "fmt"
  "bytes"
)

type Writer interface{
  Write([]byte) (int, error)
}

type Closer interface{
  Close() error
}

type WriterCloser interface{
  Writer
  Closer
}

type BufferedWriterCloser struct{
  buffer *bytes.Buffer
}

func (bwc BufferedWriterCloser) Write(data []byte) (int, error){
  n, err := bwc.buffer.Write(data)
  if err != nil {
    return 0, err
  }
  v := make([]byte, 8)
  for bwc.buffer.Len() > 8 {
    _, err := bwc.buffer.Read(v)
    if err != nil {
      return 0, err
    }
    _, err = fmt.Println(string(v))
    if err != nil {
      return 0, err
    }
  }
  return n, nil
}

func (bwc *BufferedWriterCloser) Close() error{
  for bwc.buffer.Len() > 0{
    data := bwc.buffer.Next(8)
    _, err := fmt.Println(string(data))
    if err != nil {
      return err
    }
  }
  return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
  return &BufferedWriterCloser{
    buffer: bytes.NewBuffer([]byte{}),
  }
}

func main() {
  //traditional way
  var wc WriterCloser = NewBufferedWriterCloser()
  wc.Write([]byte("Welcome to hello world!"))
  wc.Close()

  //empty interface way
  var obj interface{} = NewBufferedWriterCloser()
  wc, err := obj.(WriterCloser)
  if err == false {
    fmt.Println("error")
  }
  wc.Write([]byte("Welcome to hello world!"))
  wc.Close()
}
