package map_ext_test

import "testing"

func TestMapWithFuncVal(t *testing.T) {
	funMap := map[int]func(op int) int{}
	funMap[1] = func(n int) int { return n }
	funMap[2] = func(n int) int { return n * n }
	funMap[3] = func(n int) int { return n * n * n }
	t.Log(funMap[1](2), funMap[2](2), funMap[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true

	n := 1
	if mySet[n] {
		t.Logf("mySet[%d] is exist", n)
	} else {
		t.Logf("mySet[%d] is not exist", n)
	}

	mySet[2] = true
	t.Logf("mySet's length is %d", len(mySet))
	delete(mySet, 1)
	t.Logf("mySet's length is %d", len(mySet))

}
