package relationaloperators

import (
	"fmt"
	"time"
)

func Setup() {
	fmt.Println("Relational Operators!")

	banana := "Banana"

	fmt.Println("Strings:", banana == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println(">=", 3 >= 2)
	fmt.Println("<=", 3 <= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Dates:", d1 == d2)
	fmt.Println("Dates with Equal:", d1.Equal(d2))

	type Person struct {
		Name string
	}

	p1 := Person{"John"}
	p2 := Person{"John"}
	fmt.Println("Persons:", p1 == p2)
}
