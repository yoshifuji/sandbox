package main

import "fmt"

func split(sum int) (prm_x, prm_y int) {
	prm_x = sum * 4 / 9
	prm_y = sum - prm_x
	return
	// *same result below*
	// return prm_x, prm_y
}

func main() {
	fmt.Println(split(18))
}
