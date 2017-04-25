package main

/*
http://dev.classmethod.jp/go/golang-6/
*/
import (
	"fmt"
	"strconv"
)

type Car interface {
	run(int) string
	stop()
}

type MyCar struct {
	name  string
	speed int
}

func (u *MyCar) run(speed int) string {
	/*
		"Atoi" : A(string) to Int(10 decimal number)
		"Itoa" : Int to A
		"ParseBool" : String to bool
	*/
	u.speed = speed
	return strconv.Itoa(speed) + " km/hour running" //A to Int
}

func (u *MyCar) stop() {
	fmt.Println("stopping")
	u.speed = 0
}

func main() {
	myCar := &MyCar{name: "foobar", speed: 0}
	myCar.run(100)
	fmt.Println(myCar.speed)
	myCar.stop()
}
