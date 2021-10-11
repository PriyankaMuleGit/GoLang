package main

import f "fmt"

func main() {
	createArray()
}

func createArray() {
	var num [5]int
	var sum = 0
	f.Println("Enter 5 elements")
	for i := 0; i < len(num); i++ {
		f.Scan(&num[i])

		if num[i]%2 == 0 {
			sum += num[i]
		}
	}
	for i := 0; i < len(num); i++ {
		f.Print(num[i], "   ")
	}
	newArray := num
	newArray[4] = 111
	f.Println("\nNew Array =", newArray)

	f.Println("\nSum of Even numbers = ", sum)

	var num1 []int
	f.Println("Length", len(num1))

}
