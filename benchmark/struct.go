package benchmark

type Iface interface {
	Foo()
}

type Struct struct {
	Field1 string
}

// Benchmark_StructMetodCall-16            1000000000               0.2437 ns/op
// 这个方法和小面的没差别
func (s *Struct) Foo() {}

// Benchmark_StructPointerMetodCall-16     1000000000               0.2369 ns/op
// func (s Struct) Foo() {}
