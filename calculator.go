package main

import (
	"flag"
	"fmt"
)

func main() {
	// Calculator
	var first int
	var op string
	var second int

	flag.IntVar(&first, "first", 0, "-first= ")
	flag.StringVar(&op, "op", "", "-op= ")
	flag.IntVar(&second, "second", 0, "-second= ")

	flag.Parse()

	switch {
	case op == "+":
		sum := first + second
		fmt.Printf("Result of %d + %d = %d \n", first, second, sum)
	case op == "-":
		substract := first - second
		fmt.Printf("Result of %d - %d = %d \n", first, second, substract)
	case op == "*":
		multiply := first * second
		fmt.Printf("Result of %d * %d = %d \n", first, second, multiply)
	case op == "/":
		divided := first / second
		fmt.Printf("Result of %d / %d = %d \n", first, second, divided)
	case op == "%":
		modulo := first % second
		fmt.Printf("Result of %d % %d = %d \n", first, second, modulo)
	default:
		fmt.Println("Silahkan coba lagi \n")
	}

}
