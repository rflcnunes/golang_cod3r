package noternary

import "fmt"

func getResults(grade float64) string {
	// return grade >= 6 ? "Approved" : "Disapproved"
	if grade >= 6 {
		return "Approved"
	}

	return "Disapproved"
}

func Setup() {
	fmt.Println("GO DOES NOT HAVE TERNARIES")

	fmt.Println(getResults(6.2))
	fmt.Println(getResults(4.2))
}
