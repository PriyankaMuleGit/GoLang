//9.	Write a GoLang program to enter P, T, R and calculate Simple Interest.
//10.	Write a GoLang program to enter P, T, R and calculate Compound Interest.
package main

import "fmt"

func main() {
	var p, r, t, n float64
	var result = 1.0
	fmt.Println("Enter Principle Initial value :")
	fmt.Scan(&p)
	fmt.Println("Enter interest rate :")
	fmt.Scan(&r)
	fmt.Println("Enter time(in year) :")
	fmt.Scan(&t)

	fmt.Println("Simple Interest =", p*r*t/100)

	fmt.Println("Enter number of times compoundet in time :")
	fmt.Scan(&n)

	x := 1.0
	for i := 1.0; i <= n; i++ {
		x = (1 + (r / 100))
		result *= x
	}
	fmt.Println("-----------", result)
	fmt.Println("Compound Interest =", p*result)

}
