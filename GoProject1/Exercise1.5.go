//6.	Write a GoLang program to find power of any number x ^ y.

package main

import "fmt"

func main() {
	fmt.Println("Enter base x:")
	x := 0
	fmt.Scan(&x)
	fmt.Println("Enter exponent y:")
	y := 0
	fmt.Scan(&y)

	var result = 1
	for i := 1; i <= y; i++ {
		result *= x
	}
	fmt.Println("Result", x, "^", y, "=", result)
}
