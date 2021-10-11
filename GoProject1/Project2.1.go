package main

import "fmt"

func checkEvenOdd(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func checkDigit(no int) int {
	flag := 0

	for no != 0 {
		flag++
		no = no / 10
	}
	return flag

}

func main() {
	fmt.Println("Enter a number")
	num := 1
	fmt.Scan(&num)
	// if checkEvenOdd(num) {
	// 	fmt.Println("Even")
	// } else {
	// 	fmt.Println("Odd")
	// }
	dig := 0
	dig = checkDigit(num)
	fmt.Println("Number of digits :", dig+1)

	var s []string
	if s == nil {
		fmt.Println("Nilllll")
	}
}
