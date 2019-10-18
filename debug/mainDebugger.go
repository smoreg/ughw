package main

import (
	"fmt"
)

func main() {
	x := 123
	_ = x
	funcName()
	fmt.Println(123) // TODO (Smoreg) : REMOVE IT BAKA
}

func funcName() {
	x := 111
	_ = x
	//Smart step into
	funcName3(funcName1(), funcName2())
}

func funcName1() int {
	return funcName2()
}

func funcName2() int {
	panic(1)
	return 1
}

func funcName3(x, y int) {

}
