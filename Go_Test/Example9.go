//Delete all elements from map. ( check for two ways to do it)
package main

import f "fmt"

func main() {
	map1 := map[int]string{
		101: "Mr.X",
		102: "Mr.Y",
		103: "Mr.Z",
		104: "Mr.X",
		105: "Mr.Z",
	}
	f.Println("Map : ", map1)

	checkKey := false
	for i := range map1 {
		_, checkKey = map1[i]
		if checkKey == true {
			delete(map1, i)
		}
	}

	// or
	// map1 = nil
	f.Println(map1)
}
