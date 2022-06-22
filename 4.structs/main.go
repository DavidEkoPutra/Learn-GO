package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//alex := person{"Alex", "Anderson"}
	//fmt.Println(alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Jay",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 21212,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
