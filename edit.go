package ughw

func Unwrap() {
	some1(some2(some3(123)))
}

func Extract() int {
	some1(some2(some3(123)))

	some1(123)
	some2(123)
	some3(123)

	coolVar := some3(111)
	j := some2(coolVar)
	some1(coolVar)
	return coolVar + j
}

func Extract2() int {
	coolVar := some3(111)
	some1(coolVar)
	return 1
}

func Inline() {
	// i := "123.123"
	// j := 123
}

type ImplementExample interface {
	CoolFunc1(int) string
	CoolFunc2()
	CoolFunc3() (string, string, error)
}

type SampleStruct struct {
	A int
	B string
	C *int
}

func SelectOccurrence() {
	BadNameInt := 123
	BadNameString := "123"
	BadNameFloat := 1.0
	_, _, _ = BadNameInt, BadNameString, BadNameFloat
}

func cyclicExpand() {
	Int := 123
	_ = Int
}

func AutoComplete() {
	// panec()// tab vs enter
	// namestr := "123"
	// some4()
	// f4
	// ^q
	// ^si
	// find path
}

func InInspection() {
	someError1()
	return
}

func LTemplateAndPostfix() (string, error) {
	// new stack
	error1 := someError1()
	_ = error1

	// strings := []string{"1", "1", "1", "1",}
	// strings.a

	// FOR
	// 123

	return "", nil
}

func SomeFint(va SampleStruct) (string, error) {
	// Test check
	// Complete Current Statement
	return "", nil
}

func PartComplete() {
	// var a io.ReadWriteCloser
}

func ParametrInfo() {
	// reflect.FuncOf()
}

func Injection() {
	s := "SELECT * from COOL_TABLE"
	_ = s
}

// -----------------------------------------------------------

// <editor-fold desc="Cool folding">
// some1 cool func 1
func some1(a int) {
	return
}

// some2 cool func 2
func some2(a int) int {
	return a
}

// some3 cool func 3
func some3(a int) int {
	return a
}

// some4 cool func 4
func some4(a string) string {
	return a
}

// someError1 cool error
func someError1() error {
	return nil
}

// </editor-fold>
