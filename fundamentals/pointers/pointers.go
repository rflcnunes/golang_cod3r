package pointers

import "fmt"

func Setup() {
	fmt.Println("POINTERS!")

	i := 1

	// Go does not have pointer arithmetic
	var p *int = nil

	p = &i

	fmt.Println("Pointer p:", *p)

	*p++

	fmt.Println("Pointer p:", *p)
}
