package main

import (
	"fmt"
	"log"
)

type Foo struct {
	Bar string
	Buz string
}

func main() {
	foo := retStructure()
	log.Printf("%v", foo)
	log.Printf("%v", foo.Bar)
	log.Printf("%T", foo)

	fmt.Print("logging end")
}

func retStructure() (foo *Foo){
	foo = &Foo{"bar", "baz"}
	return foo
}