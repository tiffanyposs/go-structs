package main

import "fmt"

// create a new type called person, struct type
type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}                        // relies on order, common in Go but not great
	chris := person{firstName: "Chris", lastName: "Anderson"} // init with properties

	fmt.Println(alex)
	fmt.Println(chris)
}

// $ go run main.go
