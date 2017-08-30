package math_test

import (
	. "example/math" //.表示调用时可以省略包名
	"testing"
)

func TestMethod1(t *testing.T) {
	//t.Error(`error`) //不会终止测试
	t.Log("log1")
}

func TestMethod2(t *testing.T) {
	//t.Fatal(`error`) //会终止测试
	t.Log("log2")
}

func TestMethod3(t *testing.T) {
	t.Log(Sum(1, 2)) //Sum定义自export_test.go
}

//基准测试
//go test -bench=.
//需要其他的功能测试通过的情况下才会执行
func BenchmarkMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(Sum(1, i))
	}
}
