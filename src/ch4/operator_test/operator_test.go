package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Eexcutable
)

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//	t.Log(a == c)   // 比较两个数组， 个数及元素相同才行
	t.Log(a == d)
}
