package logics

import "fmt"

func purchases(work1, work2 bool) (bool, bool, bool) {
	purchaseTv50 := work1 && work2
	purchaseTv32 := work1 != work2
	purchaseIceCream := work1 || work2

	return purchaseTv50, purchaseTv32, purchaseIceCream
}

func Setup() {
	fmt.Println("Logics")

	tv50, tv32, iceCream := purchases(true, true)

	// Examples of unary operators:
	positive := +5           // Example of + (positive)
	negative := -5           // Example of - (negative)
	notIceCream := !iceCream // Example of ! (logical negation)

	fmt.Printf("TV 50: %t, TV 32: %t, Ice Cream: %t, Healthy: %t\n", tv50, tv32, iceCream, notIceCream)
	fmt.Printf("Positive: %d, Negative: %d, Not Ice Cream: %t\n", positive, negative, notIceCream)
}
