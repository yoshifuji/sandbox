package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

/*
range of Int type
0 ～ 18,446,744,073,709,551,615	符号なし 64 ビット整数
*/
func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
