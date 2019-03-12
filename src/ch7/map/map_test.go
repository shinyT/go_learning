package map_test

import "testing"

func TestInitMap(t *testing.T) {
	var m1 map[int]int
	m1 = make(map[int]int)
	m1[1] = 5

	m2 := map[int]int{1: 10, 2: 20, 3: 30}

	m3 := map[int]int{}
	m3[1] = 111

	t.Logf("m1's length is %d", len(m1))
	t.Logf("m2's length is %d", len(m2))
	t.Logf("m3's length is %d", len(m3))

}

func TestAccessNoExistingMap(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	m1[1] = 111
	if v, ok := m1[1]; ok {
		t.Logf("m1[1] is exist and value is %d", v)
	} else {
		t.Log("m1[1] is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m1 {
		t.Logf("m1[%s] is %d", k, v)
	}
}
