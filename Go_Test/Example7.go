//Q7. Write a code to show keys with same values in map.
//Eg. in Map { A :100 , B:100, C:300,D:400 ,E:300 , F:300}
//Out put : A,B with value 100  , C,E,F with value 300

package main

import f "fmt"

func main() {
	map1 := make(map[string]int)
	map1["A"] = 100
	map1["B"] = 100
	map1["C"] = 300
	map1["D"] = 400
	map1["E"] = 300
	map1["F"] = 300
	f.Println("My map :", map1)

	map2 := make(map[int]string)

	checkKey := false
	for i := range map1 {
		_, checkKey = map2[map1[i]]
		if checkKey == true {
			map2[map1[i]] = (map2[map1[i]] + "," + i)
		} else {
			map2[map1[i]] = i
		}
	}
	for i := range map2 {
		f.Println(map2[i], "with value :", i)
	}
}
