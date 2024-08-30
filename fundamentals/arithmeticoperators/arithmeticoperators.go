package arithmeticoperators

import (
	"fmt"
	"math"
)

func Setup() {
	fmt.Println("Arithmetic Operators!")

	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Subtraction =>", a-b)
	fmt.Println("Division =>", a/b)
	fmt.Println("Multiplication =>", a*b)
	fmt.Println("Modulus =>", a%b)

	// bitwise
	fmt.Println("And =>", a&b) // 11 & 10 = 10
	fmt.Println("Or =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	// other operations using math
	c := 3.0
	d := 2.0

	fmt.Println("Max =>", math.Max(float64(a), float64(b)))
	fmt.Println("Min =>", math.Min(c, d))
	fmt.Println("Exponentiation =>", math.Pow(c, d))
}
