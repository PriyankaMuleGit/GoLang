//Q10. Consider map1 like { a:1, b:2 ,c:3}    and map2 like  {d:4,e:5,a:1 , c:3}
//. Merge these two maps. So, output should be  {a:1,b:2,c:3,d:4,e:5}

package main

import f "fmt"

func main() {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	f.Println(map1)
	map2 := map[string]int{"d": 4, "e": 5, "a": 1, "c": 3}
	f.Println(map2)
	finalMap := make(map[string]int)

	for i := range map1 {
		checkKey := false
		_, checkKey = finalMap[i]
		if checkKey == false {
			finalMap[i] = map1[i]
		}
	}
	for i := range map2 {
		checkKey := false
		_, checkKey = finalMap[i]
		if checkKey == false {
			finalMap[i] = map2[i]
		}
	}
	f.Println("Final Map :", finalMap)
}
