package slice_test

import "fmt"
import "testing"

func TestSlice(t *testing.T) {
	var slice = make([]int, 3, 5)
	printSlice(slice)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	fmt.Println(x[0])
}
