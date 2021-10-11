package main

import f "fmt"

type item struct {
	itemm string
	price int
}

func maxPrice(n1 int, n2 int) bool {
	return true
}

func main() {
	i := item{"Kurti", 600}
	f.Println("i :", i)
	i1 := item{"Dupatta", 300}
	f.Println("i1 :", i1)

	p := &i
	f.Println("Address of i :", p)
	f.Println("Item of i using ptr:", p.itemm)
	f.Println("Price of i using ptr:", p.price)

	if maxPrice(i.price, i1.price) == true {
		f.Println("Max price :", i.price)
	} else {
		f.Println("Max price :", i1.price)
	}

	p = getitem("a", 1000)
	f.Println("Item added :", p)

	f.Println("Total Bill=", i.price+i1.price)
}
func getitem(s string, i int) *item {
	i5 := item{s, i}
	return &i5
}
