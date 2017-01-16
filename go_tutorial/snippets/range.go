package main

import "fmt"

/*
http://dev.classmethod.jp/go/golang-7/
*/
func main() {
	nums := []int{2, 3, 4}

	for i := 0; i < len(nums); i++ {
		fmt.Print(fmt.Sprintf("index: %d, value: %d\n", i, nums[i]))
	}

	for i, v := range nums {
		fmt.Print(fmt.Sprintf("index: %d, value: %d \n", i, v))
	}
}
