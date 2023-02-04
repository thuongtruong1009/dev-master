package main

import "fmt"

type contactInfo struct{
  email string
  zipCode int
}

type person struct{
  firstName string
  lastName string
  contact contactInfo
}

func (PointerToPerson *person) updateFirstName(newFirstName string){
  (*PointerToPerson).firstName = newFirstName
}

func (p person) print(){
  fmt.Printf("%+v", p)
}

func main() {
  alex := person{
    firstName: "Alex",
    lastName: "Anderson",
    contact: contactInfo{
      email: "sample@email.com",
      zipCode: 700000,
    },
  }
  alex.updateFirstName("Alexis")
  alex.print()
}
