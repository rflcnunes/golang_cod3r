package constvar

import (
	"fmt"
	"math"
)

func Constvar() {
	const PI float64 = 3.1415
	var ray = 3.2 // (float64) type inferred from compiler

	// reduces form to create an var
	area := PI * math.Pow(ray, 2)

	fmt.Println("Area is", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "test"

	fmt.Println(g, h, i)
}
