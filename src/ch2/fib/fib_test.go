package fib

import "testing"
import "fmt"

func TestFib(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// a:=1
	// b:=1
	var (
		a = 1
		b = 1
	)

	fmt.Println(a)
	for i := 1; i < 5; i++ {
		fmt.Println(b)
		tmp := a
		a = b
		b = tmp + a

	}
	fmt.Println()
}
