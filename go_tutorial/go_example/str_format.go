package main

import (
	"fmt"
	"os"
)

/*
http://qiita.com/taji-taji/items/77845ef744da7c88a6fe
*/
func main() {
	//Print
	fmt.Println("Hello", "world!")
	//Fprint
	fmt.Fprint(os.Stdout, "Hello world!")
	//Sprint
	hello_s := fmt.Sprint("Hello wolrd!")
	fmt.Print(hello_s)
	//Printf
	hello_f := "Hello world!"
	fmt.Printf("%s\n", hello_f)
	fmt.Printf("%v\n", hello_f)
}
