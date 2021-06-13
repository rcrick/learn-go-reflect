package benchmark

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

// Benchmark_Direct_New-16                 42326690                25.33 ns/op
func Benchmark_Direct_New(b *testing.B) {
	var p *Person
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = new(Person)
	}
	_ = p
}

//Benchmark_Reflect_New-16                28298707                42.64 ns/op
func Benchmark_Reflect_New(b *testing.B) {
	var p *Person
	pType := reflect.TypeOf(Person{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pVal := reflect.New(pType)
		p = pVal.Interface().(*Person)
	}
	_ = p
}

// Benchmark_Direct_Set-16                 45673189                25.48 ns/op
func Benchmark_Direct_Set(b *testing.B) {
	var p *Person
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = new(Person)
		p.Age = 1
		p.Name = "name"
	}
}

// Benchmark_Reflect_Set-16                27042921                40.57 ns/op
func Benchmark_Reflect_Set(b *testing.B) {
	pType := reflect.TypeOf(Person{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pVal := reflect.New(pType)
		var p = pVal.Interface().(*Person)
		p.Age = 1
		p.Name = "name"
	}
}

// Benchmark_Reflect_FiledByIndex-16       20692100                52.14 ns/op
func Benchmark_Reflect_FiledByIndex(b *testing.B) {
	pType := reflect.TypeOf(Person{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pVal := reflect.New(pType).Elem()
		pVal.Field(0).SetString("name")
		pVal.Field(1).SetInt(10)
	}
}

// Benchmark_Reflect_FiledByName-16         6617115               190.1 ns/op
// FieldByName will loop to find filed by name, and call .Field, to it slower than Field
func Benchmark_Reflect_FiledByName(b *testing.B) {
	pType := reflect.TypeOf(Person{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pVal := reflect.New(pType).Elem()
		pVal.FieldByName("Name").SetString("name")
		pVal.FieldByName("Age").SetInt(10)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/rcrick/learn-go-reflect/benchmark
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// Benchmark_Direct_New-16                 42326690                25.33 ns/op
// Benchmark_Reflect_New-16                28298707                42.64 ns/op

// Benchmark_Direct_Set-16                 45673189                25.48 ns/op
// Benchmark_Reflect_Set-16                27042921                40.57 ns/op
// Benchmark_Reflect_FiledByIndex-16       20692100                52.14 ns/op
// Benchmark_Reflect_FiledByName-16         6617115               190.1 ns/op
