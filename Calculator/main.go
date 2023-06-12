package main

import (
	"fmt"
	"gocons/Calculator/logic"
)

func main() {

	x, y := 4, 2
	fmt.Println("Addition of two numbers is", logic.Addition(x, y))
	fmt.Println("Subtraction of two numbers is", logic.Subtraction(x, y))
	fmt.Println("Multipication of two numbers is", logic.Multipication(x, y))
	Quotient, Remainder := logic.Divide(x, y)
	fmt.Println("Quotient and Remainder of two numbers  are", (Quotient), "and", Remainder)
}
