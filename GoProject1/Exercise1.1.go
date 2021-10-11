//1.	Write a GoLang program to enter length and breadth of a rectangle and find its perimeter
//2.	Write a GoLang program to enter length and breadth of a rectangle and find its area
package main

import "fmt"

func main() {
	var l, b float32
	fmt.Println("Enter length :")
	fmt.Scan(&l)
	fmt.Println("Enter breadth :")
	fmt.Scan(&b)

	fmt.Println("Perimeter of a rectangle =", 2*(l+b))
	fmt.Println("Area of a rectangle =", l*b)
}
