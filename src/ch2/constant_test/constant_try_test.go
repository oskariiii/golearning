package constant_test

import (
	"fmt"
	"testing"
)

//const (
//	Monday = 1 + iota
//	Tuesday
//	Wednesday
//	Thursday
//	Friday
//	Saturday
//	Sunday
//)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestContantTry(t *testing.T) {
	//t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	var a = 1
	var b = 2
	a, b = b, a
	fmt.Println("a的值是: ", a)
	fmt.Println("b的值是: ", b)
}
func TestConstantTry1(t *testing.T) {
	t.Log(Readable)
	t.Log(Writable)
	t.Log(Executable)

	a := 7 // 0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
