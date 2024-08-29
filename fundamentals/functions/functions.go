package functions

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func print(a int, b int) {
	fmt.Println(sum(a, b))
}

func Setup() {
	fmt.Println("Functions!")

	print(2, 3)
}
