package ecode

import (
	"fmt"
	"testing"
)

func TestEcode(t *testing.T) {
	code1 := New(1, "test")
	code2 := New(2, "test2")
	code3 := New(3, "test2")

	fmt.Println("m", code1.Message())
	fmt.Println(code1)
	fmt.Println("m", code2.Message())
	fmt.Println(code2)
	fmt.Println("m", code3.Message())
	fmt.Println(code3)
	_ = New(3, "test")
}
