package main

import "fmt"

func main() {
	var r float32
	fmt.Println("Enter Radius :")
	fmt.Scan(&r)

	var area float32 = 3.14 * r * r
	var peri float32 = 2 * 3.14 * r

	fmt.Println("Area of Circle = ", area)
	fmt.Println("Perimeter of Circle = ", peri)

}
