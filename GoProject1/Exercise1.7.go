//8.	Write a GoLang program to enter base and height of a triangle and find its area.

package main

import "fmt"

func main() {
	var b, h float32
	fmt.Println("Enter base :")
	fmt.Scan(&b)
	fmt.Println("Enter height :")
	fmt.Scan(&h)

	fmt.Println("Area of a Triangle =", 0.5*b*h)
}
