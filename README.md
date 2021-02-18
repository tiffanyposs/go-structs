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