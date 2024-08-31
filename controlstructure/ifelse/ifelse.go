package ifelse

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved with grade", grade)
	} else {
		fmt.Println("Disapproved with grade", grade)
	}
}

func printResultOtherValidMethod(grade float64) {
	value := true

	if grade >= 6 && (value || !value) {
		fmt.Println("Approved with grade", grade)
	} else {
		fmt.Println("Disapproved with grade", grade)
	}
}

func Setup() {
	fmt.Println("IF / ELSE")

	printResult(7.3)
	printResult(3.0)

	printResultOtherValidMethod(6)
	printResultOtherValidMethod(4)
}
