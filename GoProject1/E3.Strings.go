package main

import (
	f "fmt"
	"strings"
)

func main() {
	s1 := "Hey! I'm Priyanka"
	s2 := "hey! im dhanashree"
	s3 := "Hey! I'm Priyanka"
	f.Println()
	if s1 == s2 {
		f.Println("s1 and s2 are equal")
	} else {
		f.Println("s1 and s2 are not equal")
	}
	if s1 == s3 {
		f.Println("s1 and s3 are equal")
	} else {
		f.Println("s1 and s3 are not equal")
	}
	f.Println()
	f.Println("Length of a string s1 =", len(s1))
	f.Println()
	f.Println("s1 to Upper Case ", strings.ToUpper(s1))
	f.Println("s1 to Lower Case ", strings.ToLower(s1))
	f.Println("s2 to Title ", strings.Title(s2))
	f.Println()
	f.Println("String S1 contains Priyanka :", strings.Contains(s1, "Priyanka"))

}
