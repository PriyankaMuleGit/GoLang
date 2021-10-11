//Q8. Consider a Map in Q7. Check if key 104 is there,
//if it is present the replace it’s value with Mr.L
//Otherwise add it with value Mr.L and display the map
package main

import (
	f "fmt"
)

func main() {
	map1 := map[int]string{
		101: "Mr.X",
		102: "Mr.Y",
		103: "Mr.Z",
		//	104: "Mr.X",
		105: "Mr.Z",
	}
	f.Println("Created Map :", map1)
	for i := range map1 {
		if i == 104 {
			map1[104] = "Mr.L"
		} else if map1[i] != "Mr.L" {
			map1[106] = "Mr.L"
		}
	}
	f.Println("New Map :", map1)
}
