package ifelseif

import "fmt"

func gradeForConcept(grade float64) string {
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 8 && grade < 9 {
		return "B"
	} else if grade >= 6 && grade < 8 {
		return "C"
	} else if grade >= 3 && grade < 5 {
		return "D"
	} else if grade < 1 || grade > 10 {
		return "Invalid grade"
	} else {
		return "E"
	}
}

func Setup() {
	fmt.Println("IF / ELSEIF")

	grade := gradeForConcept(4)
	fmt.Println(grade)

	grade2 := gradeForConcept(8)
	fmt.Println(grade2)

	grade3 := gradeForConcept(10)
	fmt.Println(grade3)

	invalidGrande := gradeForConcept(15)
	fmt.Println(invalidGrande)
}
