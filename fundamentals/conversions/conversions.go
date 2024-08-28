package conversions

import (
	"fmt"
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
}
