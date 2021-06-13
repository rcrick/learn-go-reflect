package benchmark

import (
	"sort"
	"testing"
)

// Benchmark_StructPointerMetodCall-16     1000000000               0.2387 ns/op
func Benchmark_StructPointerMetodCall(b *testing.B) {
	s := &Struct{}
	for i := 0; i < b.N; i++ {
		s.Foo()
	}
}

// Benchmark_StructInterfaceCall-16        1000000000               1.208 ns/op
func Benchmark_StructInterfaceCall(b *testing.B) {
	var s Iface = &Struct{}
	for i := 0; i < b.N; i++ {
		s.Foo()
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/rcrick/learn-go-reflect/benchmark
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// Benchmark_StructPointerMetodCall-16     1000000000               0.2387 ns/op
// Benchmark_StructInterfaceCall-16        1000000000               1.208 ns/op

type SortIface interface {
	Num() int
}

type Sort struct {
	number int
}

func (s Sort) Num() int {
	return s.number
}

type SourtIfaceByNumber []SortIface

func (s SourtIfaceByNumber) Len() int           { return len(s) }
func (s SourtIfaceByNumber) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SourtIfaceByNumber) Less(i, j int) bool { return s[i].Num() < s[j].Num() }

type SourtByNumber []Sort

func (s SourtByNumber) Len() int           { return len(s) }
func (s SourtByNumber) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SourtByNumber) Less(i, j int) bool { return s[i].Num() < s[j].Num() }

// Benchmark_SortIface-16                 6         166848069 ns/op
func Benchmark_SortIface(b *testing.B) {
	s := make(SourtIfaceByNumber, 1000000)
	for i := 0; i < 1000000; i++ {
		s[i] = Sort{i}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(s)
	}
}

// Benchmark_SortStruct-16               19          59353636 ns/op
func Benchmark_SortStruct(b *testing.B) {
	s := make(SourtByNumber, 1000000)
	for i := 0; i < 1000000; i++ {
		s[i] = Sort{i}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(s)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/rcrick/learn-go-reflect/benchmark
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// Benchmark_SortIface-16                 6         166848069 ns/op
// Benchmark_SortStruct-16               19          59353636 ns/op