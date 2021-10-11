//13.	Write a GoLang program to check whether a number is divisible by 5 and 11 or not.
//14.	Write a GoLang program to check whether a year is leap year or not.
//15.	Write a GoLang program to check whether a character is alphabet or not.
//16.	Write a GoLang program to input any alphabet and check whether it is vowel or consonant.
//17.	Write a GoLang program to input any character and check whether it is alphabet, digit or special character.
package main

import "fmt"

func main() {
	fmt.Println("Enter Number to check no is divisible by 5 and 11 :")
	num := 0
	fmt.Scan(&num)
	if num%5 == 0 && num%11 == 0 {
		fmt.Println(num, "Divisible by both 5 & 9")
	} else {
		fmt.Println(num, " is not Divisible by both 5 & 9")
	}

	println("=========================================================================")

	fmt.Println("Enter Year to check it is leap or not :")
	year := 0
	fmt.Scan(&year)
	if (year%4 == 0) && ((year%400 == 0) || (year%100 != 0)) {
		fmt.Println(year, "It's a leap year ")
	} else {
		fmt.Println(year, "It's not a leap year ")
	}

	println("=========================================================================")

	var ch = "null"
	fmt.Print("Enter a character to check whether its alphabet or not :")
	fmt.Scan(&ch)
	if ch >= "a" && ch <= "z" || ch >= "A " && ch <= "Z" {
		fmt.Println(ch, "is an alphabet")
	} else {
		fmt.Println(ch, "is not an alphabet")
	}

	println("=========================================================================")

	fmt.Print("Enter an alphabet to check whether its vowel or consonant:")
	var a = "null"
	fmt.Scan(&a)
	if a == "a" || a == "e" || a == "i" || a == "o" || a == "u" || a == "A" || a == "E" || a == "I" || a == "O" || a == "U" {
		fmt.Println(a, " is an vowel")
	} else {
		fmt.Println(a, " is a Consonant")
	}
	println("=========================================================================")

	var chr = "null"
	fmt.Print("Enter a character to check whether its an alphabet or a Digit or Special Symbol :")
	fmt.Scan(&chr)
	if chr >= "a" && chr <= "z" || chr >= "A " && chr <= "Z" {
		fmt.Println(chr, "is an alphabet")
	} else if chr >= "0" && chr <= "9'" {
		fmt.Println(chr, "is a digit")
	} else {
		fmt.Println(chr, "is a special Character")
	}
}
