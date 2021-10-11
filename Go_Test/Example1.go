//Q1.Write a code to concatenate two slices.
package main

import f "fmt"

func main() {
	var s1_Cities = []string{"Mumbai", "Nashik", "Pune"}
	var s2_Cities = []string{"Dehli", "Jaipur"}

	//Concatinate 2 slices
	final_Slice := append(s1_Cities, s2_Cities...)
	f.Println("The final Slice after concatenating s.s1_Citiesand s2s2_Cities :", final_Slice)
}
