package main

import "fmt"

func main() {
	//function object
	foo := func(msg string) {
		fmt.Println(msg)
	}
	foo("hello")

	//function as args
	bar := func(f func(string)) {
		f("world!")
	}
	bar(foo)

	baz := func(f func(string)) func(string, int){
		return func(msg string, num int){
			for num > 0 {
				f(msg)
				num--
			}
		}
	}
	baz(foo)("aaa", 3)
}
