# Structs

A `struct` are a data structure, a collection of properties that are all related.

Define a struct by listing property names and their types

```go

type person struct {
	firstName string
	lastName  string
}

```

Create an instance of a struct in two ways.

This way is common but more risky because it relies on the order of the defined struct. If the order of the struct changes the values could map in a different way.

```go

func main() {
	alex := person{"Alex", "Anderson"}
}

```

This way is less risky because you are defining what property maps to which value.

```go
func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
}

```


The third way is to init an instance and then update the properies after

```go

func main() {
	var matt person   // inits matt with type of person with zero values ("" for strings)
	fmt.Println(matt) // Println will appear as { }, but in reality firstName and lastName got set to a zero value

	matt.firstName = "Matt"
	matt.lastName = "Anderson"

	fmt.Printf("%+v \n", matt) // Printf with "%+v" will print all of the fields on matt
}

```

## Embedded Structs

Below is an example of embedding a struct inside of a struct.

```go

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // set type to new contactInfo struct
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", jim)
}

```

You can also do a shortened version if the key name is the same as the struct type

```go

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo // set type to new contactInfo struct
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

	fmt.Printf("%+v", jim)
}

```

## Pointers

Memorory on your machine (RAM) can be thought of as a bunch of different slots/boxes. Each box can store some data.

Go is a "Pass by Value" language, which means that when you pass a value in Go it creates a new instance in the memory.

The below example doesn't do what you might think it does at first glance:

```go

// create a new type called person, struct type
type person struct {
	firstName string
	lastName  string
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
	}

	jim.print()
	jim.updateName("Jimmy") // doesn't actually update firstName to "Jimmy"
	jim.print() // firstName is still "
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName // p is a copy of jim, doesn't actually update the original object in the memorys
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

```

* `&<variable>` - Give me the memory address of the value this value is pointing at
* `*<type>` - We are working with a pointer that points towars a certain type
* `*<pointer>` - Give me the value this memory address is pointing at ( we want to manipulate the value the pointer is referencing)


Here is an example of using pointers (this code will now work):

```go

// create a new type called person, struct type
type person struct {
	firstName string
	lastName  string
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
	}

	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName 
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

```


Here is a simplified version of using pointers:

```go

// create a new type called person, struct type
type person struct {
	firstName string
	lastName  string
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
	}

	jim.print()
	jimPointer.updateName("Jimmy")
	jim.print()
}

// Go will automatically know to switch person a pointer to person
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName 
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

```

Cheat Sheet:

| `address` | `value`                                       |
| ------  | ----------------------------------------------- |
| 0001    | person{firstName: "Jim", lastName: "Party"}     |
| 0002    | person{firstName: "Alex", lastName: "Anderson"} |

Turn `address` into `value` with `*address`
Turn `value` into `address` with `&value`

---

Slices work differently than structs, you don't have to pass the pointer:

```go

package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice) // updates to bye
}

func updateSlice(s []string) {
	s[0] = "bye"
}

```

## Value Types vs Reference Types

### Example with Arrays and Slices

Arrays:

* Primative data structure
* Can't be resized
* Rarely used directly

Slices:

* Can grow and shrink
* Used 99% of the time for a list of elements

### Slice structures

Slices actually have a `pointer` towards an array in the data, and also have `capacity` and `length` property.

Since Go is a pass by value language, when you pass `0001` slice into a method, it does make a copy but a slice actually the copy still points towards `0002` for it's data. This is why you don't have to use pointers with Slices in Go (like you do with structs).

This is called a **reference type** in Go.

| `address` | `value`                                                                  |
| ----------| ------------------------------------------------------------------------ |
| 0001      | Original Slice: pointer to 0002 array, length, capacity                  |
| 0002      | []string{ "banana", "kiwi" }                                             |
| 0002      | Slice 0001 passed into function: pointer to 0002 array, length, capacity |

### Reference vs Value Types

Value Types:

Use pointers to change these in functions.

* Int
* Float
* String
* Bool
* Structs

Reference Types:

Don't use pointers with these.

* Slices
* Maps
* Channels
* Pointers
* Functions