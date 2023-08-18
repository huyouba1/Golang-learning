package strutil

import "testing"

// 基准测试
func BenchmarkString2Intv1(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		String2Intv1(i)
	}
}

func BenchmarkString2Intv2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		String2Intv2(i)
	}
}

/*
go test -run none -bench=. ./...
-run none  忽略测试

内存测试
go test -run none -bench=. -benchmem ./...

*/
