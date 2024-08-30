package unary

import "fmt"

func Setup() {
	fmt.Println("Unary!")

	x := 1
	y := 2

	// only postfix
	x++ // x += 1 or x = x + 1
	fmt.Println(x)

	y-- // y -= 1 or y = y - 1
	fmt.Println(y)
}
