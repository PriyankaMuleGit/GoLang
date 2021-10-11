//Q4.Write a code to remove element from slice.

package main

import f "fmt"

func removeFunc(slice []int, n int) []int {
	return append(slice[:n], slice[n+1:]...)
}

func main() {
	mySlice := []int{55, 11, 33, 66, 22}
	var newSlice []int
	f.Println("Before deleting elements :", mySlice)
	f.Println("Deleting element of 2nd index")
	newSlice = removeFunc(mySlice, 2)
	f.Println("After deleting elements Slice: ", newSlice)
}
