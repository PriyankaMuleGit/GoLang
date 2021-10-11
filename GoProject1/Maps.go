package main

import f "fmt"

func main() {
	var items = map[string]int{"item1": 700, "item2": 800, "kurti": 400}

	for k, v := range items {
		f.Println("Value of ", k, " => ", v)
	}

	f.Println()
	for k, v := range items {
		if k == "kurti" {
			items[k] = v + 500
			f.Println("Cost increased by 500 :", k, v)
		}

	}

	f.Println()

	for k, v := range items {
		if v >= 500 {
			f.Println("itens having price above 500", k, v)
		}
	}

	// emp := make(map[int]string)
	// emp[1] = "Priyanka"
	// emp[2] = "Poonam"
	// f.Println("Map : ", emp)
}
