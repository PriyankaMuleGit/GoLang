//Q3. Write a code to copy one slice to other slice

package main

import f "fmt"

func main() {
	oldSlice := []int{55, 11, 33, 66, 22}
	//var newSlice []int
	newSlice := make([]int, 5, 10)

	f.Println("Before Copying...")
	f.Println(oldSlice)
	f.Println(newSlice)

	//newSlice = oldSlice
	//or

	for i := range oldSlice {
		newSlice[i] = oldSlice[i]
	}

	f.Println("After Copying...")
	f.Println(oldSlice)
	f.Println(newSlice)
}
