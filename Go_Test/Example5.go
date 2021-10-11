// Q5.Create array of 10 integers. Create slice based on this.
// A.Write a code to display last five elements form slice                                                                                  A.Write a code to display last five elements form slice
// B.Write a code to display last element from slice

package main

import f "fmt"

func main() {
	myArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	f.Println("Original Array :", myArray)
	slice := myArray[1:]
	f.Println("CreatedSlice [1:] from myArray:", slice)
	slice1 := slice[len(slice)-5:]
	f.Println("Last 5 elements of createdSlice :", slice1)
	slice2 := slice[len(slice)-1:]
	f.Println("Last element of createdSlice :", slice2)
}
