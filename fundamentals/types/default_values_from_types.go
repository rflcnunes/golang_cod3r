package types

import (
	"fmt"
	"reflect"
)

func DefaultValues() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e)

	// -------------------------------------------------------------
	// # Exercise 1: Default value of an int
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type int without assigning a value.
	// Print the value of this variable to check the default value.

	// Code:
	var f int
	fmt.Println(f, reflect.TypeOf(f))

	// -------------------------------------------------------------
	// # Exercise 2: Default value of a float64
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type float64 without assigning a value.
	// Print the value of this variable to check the default value.

	// Code:
	var g float64
	fmt.Println(g, reflect.TypeOf(g))

	// -------------------------------------------------------------
	// # Exercise 3: Default value of a string
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type string without assigning a value.
	// Print the value of this variable to check the default value.

	// Code:
	var h string
	fmt.Println(h, reflect.TypeOf(h))

	// -------------------------------------------------------------
	// # Exercise 4: Default value of a bool
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type bool without assigning a value.
	// Print the value of this variable to check the default value.

	// Code:
	var i bool
	fmt.Println(i, reflect.TypeOf(i))

	// -------------------------------------------------------------
	// # Exercise 5: Default value of an array
	// -------------------------------------------------------------

	// Objective:
	// Declare an array of 3 integers without assigning values to the elements.
	// Print the values of the elements to check the default values.

	// Code:
	var j [3]int
	// or
	// var j []int
	fmt.Println(j, reflect.TypeOf(j))
}
