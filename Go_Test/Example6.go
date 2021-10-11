//Q6.Consider Map of rollno as key and teacher name as value like
//{ 101 : Mr.X , 102 : Mr.Y , 103 : Mr.Z , 104 : Mr.X , 105 : Mr.Z }
//from this map create a Map like
//{Mr.X : 2 , Mr.Y : 1 , Mr.Z : 2}

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
	f.Println("Created Map :", map1)

	map2 := make(map[string]int)
	checkKey := false
	count := 0
	for i := range map1 {
		count, checkKey = map2[map1[i]]
		if checkKey == true {
			count = count + 1
			map2[map1[i]] = count
		} else {
			map2[map1[i]] = 1
		}
	}
	f.Println("New Map :", map2)
}
