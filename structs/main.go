package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	dwight := person{
		firstName: "Dwight",
		lastName:  "Schrute",
		contactInfo: contactInfo{
			email:   "dwight@dwight.fearless",
			pinCode: 700154,
		},
	}

	dwight.updateName("Danger")
	dwight.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
