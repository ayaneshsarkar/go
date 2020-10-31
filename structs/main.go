package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var ayanesh person

	ayanesh.firstName = "Andy"
	ayanesh.lastName = "Heisenberg"

	fmt.Println(ayanesh)
	fmt.Printf("%+v", ayanesh)
}
