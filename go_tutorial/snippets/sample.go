package main
//http://qiita.com/hayashier/items/65cd8c0592caebd5a61d
//Thanks: hayashier++
/*
型
変数・定数
演算子
for
if, switch
関数
  定義方法
  可変長パラメータ
  名前付き戻り値
  関数リテラル
遅延実行
構造体、インタフェース
配列
スライス
マップ
エラーハンドリング
エラー処理
  パニック、リカバリ
Goルーチン
  チャネル: 作成、クローズ
 */

import (
	"fmt"
)

/*
const
 */
const (
	ZERO  = iota
	ONE
	TWO
	THREE
	FOUR
)
const f float32 = 123.345

func main() {

	//fconst()
	conds()
	//floop()
	//varLength(1,2,3,4,5)

	/*
	d, c, b := fname()
	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(b)
	*/



}

func fconst() {
	fmt.Println(f)
	fmt.Println(ZERO)
	fmt.Println(ONE)
}

func conds(){
	x := 9
	for x > 0 {
		x--
		fmt.Println("Count: ", x)
	}

	y := 1
	if y > 0 {
		fmt.Println(y)
	}

	i := 3
	switch i {
		case 1: fmt.Println("hoge")
		case 2:fmt.Println("fuga")
		case 3: fmt.Println("piyo")
		default: fmt.Println("default")
	}
}

func floop(){
	/*
	Loop
	 */
	// Pattern1
	for value := 0;; {
		value++
		fmt.Println(value)
		if value % 10 == 0 {
			break
		}
	}
	// Pattern2
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	// Pattern3
	items := map[int]int{0:10, 1:20}
	for key, value := range items {
		fmt.Println("key:", key, " value:", value)
	}
	// Pattern4
	dictionary := map[string]int{
		"a": 123,
		"b": 456,
		"c": 789,
	}
	for key, value := range dictionary {
		fmt.Println("key:", key, " value:", value)
	}
}

/*
variable length
 */
func varLength(args ...int) {
	fmt.Println(args)
}

/*
named returns
 */
func fname()(dog string, cat string, bird string){
	dog = "wan"
	cat = "nya-"
	bird = "ka-"
	return
}

