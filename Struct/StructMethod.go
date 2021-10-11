package main

import f "fmt"

type item struct {
	id   int
	itm  string
	cost int
}
type integer int
func main() {
	i0 := item{1, "A", 400}
	i1 := item{2, "B", 600}

	f.Println(i0)
	f.Println(i1)

	i1.checkGrade()
	p:=integer(i0.cost).updateCost(200)
	f.Println("-------",p)
}

func (rec item) checkGrade() {
	if rec.cost > 500 {
		f.Println(rec.itm, "is A grade item")
	} else {
		f.Println(rec.itm, "is not A grade item")
	}
}

func (rece integer) updateCost(n int) item {
	i5:=rece.
    return i5
}



