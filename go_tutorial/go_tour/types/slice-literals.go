package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13, 15}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{10, true},
		{11, false},
		{12, false},
		{13, true},
		{14, true},
		{15, false},
		{16, true},
		{17, true},
		{18, true},
		{19, true},
	}
	fmt.Println(s)

	/*
	default slice
	 */
	s1 := s[1:4]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := s[8:]
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := s[:3]
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	/*
	nil slice
	 */
	var t []int
	fmt.Println(t, len(t), cap(t))
	if t == nil {
		fmt.Println("nil!")
	}

}
