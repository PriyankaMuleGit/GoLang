package main

import (
	f "fmt"
)

func main() {
	s1 := "Jay Jawan Jay Kisan"
	for i := range s1 {
		f.Println(string(s1[i]))
	}
}
