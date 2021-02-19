package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// create a new type called person, struct type
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.print()
	jim.updateName("Jimmy")
	jim.print()
}

// *person is like "receive pointer that points towards a person type"
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName //
}

func (p person) print() {
	fmt.Printf("%+v \n\n\n", p)
}

// $ go run main.go
