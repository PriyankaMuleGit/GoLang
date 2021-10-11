package main

import "fmt"

func main() {
	// s := []int{1, 2, 3, 674, 67}
	// fmt.Println("before Sorting", s)
	// sort.Ints(s)
	// fmt.Println("Sorted slice", s)

	// sum := 0
	// for i := 0; i < len(s); i++ {
	// 	if s[i]%2 == 0 {
	// 		sum += s[i]
	// 	}
	// }
	// fmt.Println("Sum:", sum)

	s1 := []string{"Mumbai", "Mumbai", "Nashik", "Pune", "Pune"}
	cities := make(map[string]int)

	checkPresent := false
	count := 0
	for i := 0; i < len(s1); i++ {
		count, checkPresent = cities[s1[i]]
		if checkPresent == true {
			count = count + 1
			cities[s1[i]] = count
		} else {
			cities[s1[i]] = 1
		}
	}
	fmt.Println("maps :", cities)
	// var flag = 0
	// var flag2 = 0
	// var flag3 = 0
	// for i := 0; i < len(s1); i++ {
	// 	if s1[i] == "Pune" {
	// 		flag++
	// 	}
	// 	if s1[i] == "Mumbai" {
	// 		flag2++
	// 	}
	// 	if s1[i] == "Nashik" {
	// 		flag3++
	// 	}

	// }
	// fmt.Println("Pune repeated", flag, "times")
	// fmt.Println("Mumbai repeated", flag2, "times")
	// fmt.Println("Nashik repeated", flag3, "times")

}
