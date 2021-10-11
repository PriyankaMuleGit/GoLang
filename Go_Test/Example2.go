//Q2. Create a Slice using make method. Put elements in the slice.
// Write a code to find out position of element in sclice using ways for loop and range.
package main

import f "fmt"

func main() {
	myslice := make([]string, 4)
	myslice[0] = "Seeta"
	myslice[1] = "Rama"
	myslice[2] = "Shyam"
	myslice[3] = "Geeta"
	f.Println(myslice)

	// for i := 0; i < len(myslice); i++ {
	// 	f.Println(myslice[i], " is at position :", i)
	// }
	//or

	for i := range myslice {
		f.Println(myslice[i], " is at position :", i)
	}

}
