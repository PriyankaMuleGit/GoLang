package main

import "fmt"

func main() {
	var add = 0.0
	const total = 500.0
	var marks = 0.0
	fmt.Println("Enter 5 subject marks out of 100...")

	i := 1
	for i = 1; i <= 5; i++ {
		fmt.Println("Enter sub marks : ")
		fmt.Scan(&marks)
		add = add + marks
	}

	per := 1.0
	per = (add / total) * 100
	fmt.Println("Percentage = ", per)

	if per >= 60 && per <= 100 {
		fmt.Println("First class")
	} else if per > 50 && per < 60 {
		fmt.Println("Second class")
	} else {
		fmt.Println("Average marks")
	}
}
