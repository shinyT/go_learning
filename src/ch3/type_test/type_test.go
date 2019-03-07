package type_test

import "testing"

type MyInt int64

func TestTypeImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a) // 任意类型之间不能隐式类型转换
	var c MyInt
	c = MyInt(b) // 即使 别名也必须显示转换
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1  // 不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var str string
	t.Log("*" + str + "*")
	t.Log(len(str))

}
