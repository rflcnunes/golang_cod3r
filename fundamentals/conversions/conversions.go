package conversions

import (
	"fmt"
	"reflect"
	"strconv"
)

func Setup() {
	fmt.Println("Conversions!")

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// Warning
	fmt.Println("Test " + string(97))

	// int to string
	fmt.Println("INTO TO STRING: ", strconv.Itoa(97))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num, num-122)

	// string to bool
	b, _ := strconv.ParseBool("true") // 1 or true its true, another number or text is false
	if !b {
		fmt.Println("It's false!")
	} else {
		fmt.Println("It's true!")
	}

	// -------------------------------------------------------------
	// # Exercise 1: Default Value and Conversion from int to float64
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type int without assigning a value and a variable of type float64.
	// Assign the default value of the int variable to the float64 variable using type conversion.
	// Display the value of the float64 variable.

	// Code:
	var excInt int
	var excFloat float64

	excInt = 2
	excFloat = float64(excInt)

	fmt.Println("The type of excFloat and value is ", reflect.TypeOf(excFloat), excFloat)

	// -------------------------------------------------------------
	// # Exercise 2: Default Value and Conversion from float64 to int
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type float64 without assigning a value and a variable of type int.
	// Assign the default value of the float64 variable to the int variable using type conversion.
	// Display the value of the int variable.

	// Code:
	var excFloat2 float64
	var excInt2 int

	excFloat2 = 2.5
	excInt2 = int(excFloat2)

	fmt.Println("The type of excInt2 is and value is ", reflect.TypeOf(excInt2), excInt2)

	// -------------------------------------------------------------
	// # Exercise 3: Conversion from string to int
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type string with a numeric value (e.g., "42").
	// Convert this string to int and display the conversion result.

	// Code:
	var excString string = "42"

	excInt3, _ := strconv.Atoi(excString)

	fmt.Println("The type of excInt3 and value is ", reflect.TypeOf(excInt3), excInt3)

	// -------------------------------------------------------------
	// # Exercise 4: Conversion from int to string
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type int with a value (e.g., 123).
	// Convert this int to string and display the conversion result.

	// Code:
	var excInt4 int = 123

	excString2 := strconv.Itoa(excInt4)

	fmt.Println("The type of excString2 and value is ", reflect.TypeOf(excString2), excString2)

	// -------------------------------------------------------------
	// # Exercise 5: Default Value and Conversion from bool to string
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type bool without assigning a value and a variable of type string.
	// Assign the default value of the bool variable to the string variable using type conversion.
	// Display the value of the string variable.

	// Code:
	var excBool bool
	var excString3 string

	excBool = true
	excString3 = strconv.FormatBool(excBool)

	fmt.Println("The type of excString3 is and value ", reflect.TypeOf(excString3), excString3)

	// -------------------------------------------------------------
	// # Exercise 6: Conversion between numeric types
	// -------------------------------------------------------------

	// Objective:
	// Declare variables of different numeric types (e.g., int, float32).
	// Convert the value of one variable to another numeric type and display the conversion result.

	// Code:
	var excInt5 int = 42
	var excFloat3 float32 = float32(excInt5)

	fmt.Println("The type of excFloat3 is and value ", reflect.TypeOf(excFloat3), excFloat3)

	// -------------------------------------------------------------
	// # Exercise 7: Conversion from byte to string
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type byte and assign a value (e.g., 65).
	// Convert the byte to a string that represents the corresponding character and display the result.

	// Code:
	var excByte byte = 65

	excString4 := string(excByte)

	fmt.Println("The type of excString4 is and value ", reflect.TypeOf(excString4), excString4)

	// -------------------------------------------------------------
	// # Exercise 8: Conversion from string to float64
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type string with a decimal numeric value (e.g., "3.14").
	// Convert this string to float64 and display the conversion result.

	// Code:
	var excString5 string = "3.14"

	excFloat4, _ := strconv.ParseFloat(excString5, 64)

	fmt.Println("The type of excFloat4 and value is ", reflect.TypeOf(excFloat4), excFloat4)

	// -------------------------------------------------------------
	// # Exercise 9: Conversion from int to complex128
	// -------------------------------------------------------------

	// Objective:
	// Declare a variable of type int with a value (e.g., 5).
	// Convert this int to complex128 and display the conversion result.

	// Code:
	var excInt6 int = 5

	excComplex := complex(float64(excInt6), 0)

	fmt.Println("The type of excComplex is ", reflect.TypeOf(excComplex), "and value is ", excComplex)

	// -------------------------------------------------------------
	// # Exercise 10: Conversion between types using conversion functions
	// -------------------------------------------------------------

	// Objective:
	// Use conversion functions (such as `strconv.Atoi`, `strconv.FormatFloat`) to convert between string and numeric types.
	// Display the conversion results.

	// Code:
	excString6 := "42"

	excInt7, _ := strconv.Atoi(excString6)

	fmt.Println("The type of excInt7 and value is ", reflect.TypeOf(excInt7), excInt7)
}
