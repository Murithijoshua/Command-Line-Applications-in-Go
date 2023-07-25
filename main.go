package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	// Creating a pointer to a new zero-value instance of the Person struct.
	personPtr := new(Person)

	// Accessing and modifying fields of the struct.
	personPtr.firstname = "Jane"
	personPtr.lastname = "Smith"
	personPtr.age = 25

	// Accessing the fields of the struct using the pointer.
	fmt.Println("First Name:", personPtr.firstname)
	fmt.Println("Last Name:", personPtr.lastname)
	fmt.Println("Age:", personPtr.age)
}
