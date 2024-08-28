package prints

import "fmt"

func Setup() {
	fmt.Println("Prints!")
	fmt.Println("--------------")

	fmt.Print("Same")
	fmt.Print(" line.")

	x := 4.209348

	// Incorrect: fmt.Println("x: " + x)

	xs := fmt.Sprint(x)

	fmt.Println("x: " + xs)
	fmt.Println("x: ", x)

	fmt.Printf("x: %.2f\n", x)

	a := 10
	b := 2.999
	c := true
	d := "Hello"

	fmt.Printf("\n%d %f %.1f %t %s\n", a, b, b, c, d) // or %v
	fmt.Printf("%v %v %v %v\n", a, b, c, d)

	// -------------------------------------------------------------
	// # Exercise 1: Printing a Message
	// -------------------------------------------------------------

	// Objective:
	// Print the phrase "Hello, World!".

	// Code:
	fmt.Printf("Hello, World!\n")

	// -------------------------------------------------------------
	// # Exercise 2: Displaying Variables
	// -------------------------------------------------------------

	// Objective:
	// Declare three variables: a string, an int, and a float64.
	// Display these variables in the output.

	// Code:
	fmt.Printf("%s, %d, %f \n", "String", 1, 2.45)

	// -------------------------------------------------------------
	// # Exercise 3: Displaying Numbers in Different Formats
	// -------------------------------------------------------------

	// Objective:
	// Given an integer, display its value in decimal, binary, octal, and hexadecimal.

	// Code:
	num := 42
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Binary: %b\n", num)
	fmt.Printf("Octal: %o\n", num)
	fmt.Printf("Hexadecimal: %x\n", num)

	// -------------------------------------------------------------
	// # Exercise 4: Text Alignment
	// -------------------------------------------------------------

	// Objective:
	// Create a table that shows the names of five fruits and their respective prices,
	// ensuring proper text alignment.

	// Code:
	fmt.Printf("%-10s | %-10s\n", "Fruit", "Price")
	fmt.Println("---------------------")
	fmt.Printf("%-10s | %-10s\n", "Apple", "$ 1.25")

	// -------------------------------------------------------------
	// # Exercise 5: Adding Leading Zeros
	// -------------------------------------------------------------

	// Objective:
	// Display an integer with at least 6 digits, filling the leading spaces with zeros.

	// Code:
	num2 := 42
	fmt.Printf("Number: %06d\n", num2)

	// -------------------------------------------------------------
	// # Exercise 6: Displaying Special Characters
	// -------------------------------------------------------------

	// Objective:
	// Display a string containing a tab, a newline, and a backslash.

	// Code:
	fmt.Printf("Tab: \t Newline: \n Backslash: \\ \n")

	// -------------------------------------------------------------
	// # Exercise 7: Storing Formatted Output
	// -------------------------------------------------------------

	// Objective:
	// Capture the formatted output of a string and store the result in a variable.
	// Then, display the content of that variable.

	// Code:
	str := fmt.Sprintf("Number: %06d\n", num2)
	fmt.Println(str)

	// -------------------------------------------------------------
	// # Exercise 8: Customizing the Output of Structures
	// -------------------------------------------------------------

	// Objective:
	// Create a structure that represents a point in 2D (x, y) and display it in the format "(x, y)".

	// Code:
	type Point struct {
		x int
		y int
	}

	p := Point{1, 2}
	fmt.Printf("Point: (%d, %d)\n", p.x, p.y)

	// -------------------------------------------------------------
	// # Exercise 9: Formatting Date and Time
	// -------------------------------------------------------------

	// Objective:
	// Display the current date and time in the format "DD/MM/YYYY HH:MM:SS".

	// Code:
	fmt.Printf("Date and Time: %s\n", "01/01/2021 12:00:00")

	// -------------------------------------------------------------
	// # Exercise 10: Displaying Memory Address
	// -------------------------------------------------------------

	// Objective:
	// Declare an integer variable and display its memory address.

	// Code:
	num3 := 42
	fmt.Printf("Memory Address: %p\n", &num3)
}
