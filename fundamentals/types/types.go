package types

import (
	"fmt"
	"math"
	"reflect"
)

func Setup() {
	fmt.Println("Types!")

	// Integer types
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer literal is", reflect.TypeOf(32000))

	// Without signals (Positives) ... uint8 uint16 uint21 uint64
	var b byte = 255
	fmt.Println("The byte is", reflect.TypeOf(b))

	// With signals ... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("The maximum value of int is", i1)
	fmt.Println("The type of i1 is", reflect.TypeOf(i1))

	var i2 rune = 'a' // represents a mapping of the Unicode table (int32)
	fmt.Println("The rune is", reflect.TypeOf(i2))

	// Real numbers
	var x float32 = 49.99
	// or
	// var x = float32(49.99)
	fmt.Println("The type of x is", reflect.TypeOf(x))
	fmt.Println("The type of the literal 49.99 is", reflect.TypeOf(49.99))

	// Boolean
	bo := true
	fmt.Println("The type of bo is", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// Strings
	s1 := "Hello, my name is Rafael"
	fmt.Println(s1 + "!")
	fmt.Println("The length of string s1 is", len(s1))

	// Strings with multiple lines
	s2 := `
		Hello
		my
		name
		is
		Rafael
	`
	fmt.Println("The length of string s2 is", len(s2))
}
