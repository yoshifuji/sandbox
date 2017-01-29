package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		basic
	*/
	i := 32

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	/*
	   function
	*/
	num := 10
	inc1(num)
	fmt.Println(num)

	inc2(&num)
	fmt.Println(num)
}

func inc1(i int) {
	i++
	fmt.Println("inc1: i = " + strconv.Itoa(i))
}

func inc2(i *int) {
	*i++
	fmt.Println("inc2: i = " + strconv.Itoa(*i))
}
