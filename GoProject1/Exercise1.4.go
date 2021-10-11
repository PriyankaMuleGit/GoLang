//5.	Write a GoLang program to convert days into years, weeks and days.
package main

import "fmt"

func main() {
	fmt.Println("Enter Days :")
	days := 0
	fmt.Scan(&days)

	years := 0
	weeks := 0
	dayss := 0

	years = days / 365
	weeks = (days % 365) / 7
	dayss = days - (years*365 + weeks*7)
	fmt.Println("Years =", years)
	fmt.Println("Weeks =", weeks)
	fmt.Println("Dayss =", dayss)

}
