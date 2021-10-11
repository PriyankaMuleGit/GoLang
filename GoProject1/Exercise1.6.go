//7.	Write a GoLang program to enter any number and calculate its square.
package main

import "fmt"

func main() {
	fmt.Println("Enter Number :")
	no := 0
	fmt.Scan(&no)
	fmt.Println("Square of a number", no, "is : ", no*no)
}
