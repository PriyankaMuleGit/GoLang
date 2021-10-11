package main

import f "fmt"

func main() {
	var data = map[string]string{"Mr.x": "BJP", "Mr.y": "NCP", "Mr.z": "BJP"}
	final1 := make(map[string]int)
	for k, _ := range data {
		keyPresent := false
		count := 0
		count, keyPresent = final1[data[k]]
		if keyPresent == true {
			count++
			final1[data[k]] = count
		} else {
			final1[data[k]] = 1
		}
	}
	f.Println(final1)
}
