package benchmark

import (
	"reflect"
	"testing"
)

type Person struct {
	Name  string
	Age   int
	Sex   string
	Score int
}

// Benchmark_Direct_Set-16                         1000000000               0.2491 ns/op
func Benchmark_Direct_Set(b *testing.B) {
	var p = new(Person)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Name = "name"
		p.Age = 1
		p.Sex = "male"
		p.Score = 100
	}
}

// Benchmark_Reflect_Set-16                        838860561                1.415 ns/op
func Benchmark_Reflect_Set(b *testing.B) {
	pTyp := reflect.TypeOf(Person{})
	p := reflect.New(pTyp).Interface().(*Person)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Name = "name"
		p.Age = 1
		p.Sex = "male"
		p.Score = 100
	}
}

// Benchmark_Reflect_SetFiledByIndex-16            50546955                23.33 ns/op
func Benchmark_Reflect_SetFiledByIndex(b *testing.B) {
	pType := reflect.TypeOf(Person{})
	pVal := reflect.New(pType).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pVal.Field(0).SetString("name")
		pVal.Field(1).SetInt(1)
		pVal.Field(2).SetString("male")
		pVal.Field(3).SetInt(100)
	}

}

// Benchmark_Reflect_SetFiledByName-16              4385136               264.9 ns/op
func Benchmark_Reflect_SetFiledByName(b *testing.B) {
	pTyp := reflect.TypeOf(Person{})
	pVal := reflect.New(pTyp).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pVal.FieldByName("Name").SetString("name")
		pVal.FieldByName("Age").SetInt(1)
		pVal.FieldByName("Sex").SetString("male")
		pVal.FieldByName("Score").SetInt(100)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/rcrick/learn-go-reflect/benchmark2
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// Benchmark_Direct_Set-16                         1000000000               0.2491 ns/op
// Benchmark_Reflect_Set-16                        838860561                1.415 ns/op
// Benchmark_Reflect_SetFiledByIndex-16            50546955                23.33 ns/op
// Benchmark_Reflect_SetFiledByName-16              4385136               264.9 ns/op
