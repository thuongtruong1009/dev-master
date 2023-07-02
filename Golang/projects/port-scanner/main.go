package main

import (
   "net"
   "fmt"
   "strconv"
   "time"
)

func PortChecker() {
   for i := 1; i < 65535; i++ {
      port := strconv.FormatInt(int64(i), 10)
      conn, err := net.Dial("tcp", "127.0.0.1:" + port)
      if err == nil {
         fmt.Println("Port",i, "open")
         conn.Close()
      }
   }
}

func PortScanner(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	fmt.Println("Port Scanning")
	open := PortScanner("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)
}
