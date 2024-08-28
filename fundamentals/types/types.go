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

	// -------------------------------------------------------------
	// # Exercise 1: Variable Declaration
	// -------------------------------------------------------------

	// Objective:
	// Declare variables using basic types: int, float64, string, and bool.
	// Assign values to these variables and display them.

	// Code:
	var exInt int
	var exFloat float64
	var exStr string
	var exBool bool

	exInt = 25
	exFloat = 25.7
	exStr = "Hello friend"
	exBool = false

	fmt.Printf("\n %d, %f, %s, %t \n", exInt, exFloat, exStr, exBool)

	// -------------------------------------------------------------
	// # Exercise 2: Type Conversion
	// -------------------------------------------------------------

	// Objective:
	// Convert a value from int to float64 and vice versa.
	// Display the results of the conversions.

	// Code:
	var exInt2 int = 25
	var exFloat2 float64 = 25.7

	fmt.Printf("\n %f, %d \n", float64(exInt2), int(exFloat2))

	// -------------------------------------------------------------
	// # Exercise 3: Arrays
	// -------------------------------------------------------------

	// Objective:
	// Declare an array of 5 integers.
	// Assign values to the elements of the array and display all elements.

	// Code:
	var exArray [5]int
	exArray[0] = 1
	exArray[1] = 2
	exArray[2] = 3
	exArray[3] = 4
	exArray[4] = 5

	fmt.Printf("\n %d, %d, %d, %d, %d \n", exArray[0], exArray[1], exArray[2], exArray[3], exArray[4])
	fmt.Println(exArray, len(exArray))

	// -------------------------------------------------------------
	// # Exercise 4: Slices
	// -------------------------------------------------------------

	// Objective:
	// Create a slice of strings with three elements.
	// Add two more elements to the slice and display the complete slice.

	// Code:
	exSlice := []string{"Hello", "World", "and"}
	exSlice = append(exSlice, "Good", "Morning")

	fmt.Println(exSlice)

	// -------------------------------------------------------------
	// # Exercise 5: Maps
	// -------------------------------------------------------------

	// Objective:
	// Create a map that associates country names with their respective ISO codes.
	// Add at least three countries to the map and display the code of a specific country.

	// Code:
	exMap := map[string]string{
		"Brazil":        "BR",
		"United States": "US",
		"Germany":       "DE",
	}

	fmt.Println(exMap["Brazil"])
	fmt.Println(exMap)

	// -------------------------------------------------------------
	// # Exercise 6: Simple Structs
	// -------------------------------------------------------------

	// Objective:
	// Create a struct that represents a person, with fields for name, age, and city.
	// Create an instance of that struct and display its values.

	// Code:
	var simpleStruct struct {
		name string
		age  int
	}

	simpleStruct.name = "Rafael"
	simpleStruct.age = 25

	fmt.Println(simpleStruct)

	// -------------------------------------------------------------
	// # Exercise 7: Nested Structs
	// -------------------------------------------------------------

	// Objective:
	// Create a struct that represents an address, with fields for street, city, and state.
	// Nest that struct inside a struct that represents a person.
	// Display the complete address of a person.

	// Code:
	type Address struct {
		street string
		city   string
	}

	type Person struct {
		name    string
		age     int
		address Address
	}

	var nestedStruct Person
	nestedStruct.name = "Rafael"
	nestedStruct.age = 25
	nestedStruct.address = Address{"Rua 1", "São Paulo"}

	fmt.Println(nestedStruct)

	// -------------------------------------------------------------
	// # Exercise 8: Pointers
	// -------------------------------------------------------------

	// Objective:
	// Declare an int variable and a pointer to that variable.
	// Modify the value of the variable using the pointer and display the new value.

	// Code:

	var exInt3 int = 25
	var exPointer *int = &exInt3

	*exPointer = 30

	fmt.Println(exInt3)

	// -------------------------------------------------------------
	// # Exercise 9: Constants
	// -------------------------------------------------------------

	// Objective:
	// Declare a constant for the gravity of Earth (9.8) and use it to calculate
	// the force exerted on an object with a mass of 10 kg. Display the result.

	// Code:
	const gravity float64 = 9.8
	var mass float64 = 10

	fmt.Printf("The force is: %f", gravity*mass)

	// -------------------------------------------------------------
	// # Exercise 10: Custom Type
	// -------------------------------------------------------------

	// Objective:
	// Create a custom type called "Temperature" as an alias for float64.
	// Declare a variable of that type, assign a value to it, and display it.

	// Code:
	type Temperature float64

	var temp Temperature = 25.5

	fmt.Println(temp)

	// -------------------------------------------------------------

	// Default values
	DefaultValues()
}
