//3.	Write a GoLang program to enter radius of a circle and find its diameter, circumference and area.
package main

import "fmt"

func main() {
	var r float32
	fmt.Println("Enter Radius :")
	fmt.Scan(&r)
	fmt.Println("Diameter of Circle", 2*r)
	fmt.Println("Circumference of a circle ", 2*3.14*r)
}
