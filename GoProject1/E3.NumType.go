package main

import f "fmt"

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
func isPrime(n1 int) bool {
	if n1 == 1 {
		return true
	}
	for i := 2; i <= n1/2; i++ {
		if n1%i == 0 {
			return true
		} else {
			return false
		}
	}
}

func main() {
	s1 := []int{22, 53, 67, 11, 88}
	for i := range s1 {
		if isEven(s1[i]) == true {
			f.Println(s1[i], " is Even")
		} else {
			f.Println(s1[i], " is Odd")
		}
		if isPrime(s1[i]) == true {
			f.Println(s1[i], " is Prime")
		} else {
			f.Println(s1[i], " is not Prime")
		}
	}
}
