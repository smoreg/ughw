package trouble

type SomeStr struct{}

func (s SomeStr) Method1() {
	SomeFunc("123")
	s.SomeMethod()
}

func (s SomeStr) SomeMethod(str string) {}

func SomeFunc(s string) {}
