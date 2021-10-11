package main

import f "fmt"

type Item struct {
	item  string
	price int
}

func main() {
	i1 := Item{"Kurti", 800}
	i2 := Item{"Dupatta", 300}
	i3 := Item{"Jacket", 700}
	i4 := Item{"Jeans", 800}
	i5 := Item{"T-shirt", 400}

	var arr [5]Item
	arr[0] = i1
	arr[1] = i2
	arr[2] = i3
	arr[3] = i4
	arr[4] = i5

	total := 0
	for i := range arr {
		// f.Println("Array index :", i, "=", arr[i])
		total += arr[i].price
	}

	f.Println("Average of prices is :", total/len(arr))
	f.Println(arr)

	slice := make([]Item, 5)
	slice = arr[len(arr)-2:]
	f.Println("Last 2 items :", slice)

}
