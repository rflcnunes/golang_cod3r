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
}
