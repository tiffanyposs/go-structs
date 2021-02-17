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