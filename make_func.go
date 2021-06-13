package test

import (
	"fmt"
	"reflect"
)

func sum(args []reflect.Value) []reflect.Value {
	arg1, arg2 := args[0], args[1]
	if arg1.Kind() != arg2.Kind() {
		fmt.Println("type of arguments are not same")
		return nil
	}
	switch arg1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(arg1.Int() + arg2.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(arg1.Uint() + arg2.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(arg1.Float() + arg2.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(arg1.String() + arg2.String())}
	default:
		fmt.Println("not support this type: ", arg1.Type())
		return []reflect.Value{}
	}
}

func MakeSum(fptr interface{}) {
	// fptr is pointer to a func
	// obtain the func it self
	// then we use it to makefunc
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)
}
