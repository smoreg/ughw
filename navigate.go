package ughw

import (
	"fmt"
)

//region CH
type NavInt interface {
	Func1(x int) int
	Func2(x int) int
	Func3(x int) int
	Func(x int) int
}

type NavStr struct{}

func (n NavStr) Func1(x int) int {
	return n.Func2(x)
}

func (n NavStr) Func2(x int) int {
	return n.Func3(x)
}

func (n NavStr) Func3(x int) int {
	j := 3
	_ = j
	// return n.Func(x)
	return NavInt(n).Func(x)
}

func (n NavStr) Func(x int) int {
	fmt.Println(123) // TODO (Smoreg) : REMOVE IT BAKA
	return x
}

func (n NavStr) Func4(x int) int {
	return n.Func2(x)
}

func (n NavStr) Func5(x int) int {
	return NavInt(n).Func2(x)
}

//endregion
//region TypeH
type NavStr2 struct {
	NavStr
}

func (n2 NavStr2) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (n2 NavStr2) Sum(b []byte) []byte {
	panic("implement me")
}

func (n2 NavStr2) Reset() {
	panic("implement me")
}

func (n2 NavStr2) Size() int {
	panic("implement me")
}

func (n2 NavStr2) BlockSize() int {
	panic("implement me")
}

func (n2 NavStr2) Sum64() uint64 {
	panic("implement me")
}

func (n NavStr2) Error() string {
	panic("implement me")
}

//endregion
