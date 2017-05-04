package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Name struct {
	FirstName string
	LastName  string
}

type Address struct {
	Country    string
	PostalCode string
	Address    string
}

type Person struct {
	Name    Name
	Address Address
}

func main() {
	person := &Person{
		Name{"Cynthia", "Stone"},
		Address{"us", "02110", "1091 Lyon Avenue, Boston"},
	}
	spew.Dump(person)
}
