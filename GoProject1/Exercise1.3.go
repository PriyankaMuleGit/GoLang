//4.	Write a GoLang program to enter length in centimeter and convert it into meter and kilometer

package main

import "fmt"

func main() {
	fmt.Println("Enter length in centimeters :")
	l := 0.0
	fmt.Scan(&l)

	fmt.Println(l, "Centimeters =", l/100, "Meters")
	fmt.Println(l, "Centimeters =", l/10000, "Kilometers")
}
