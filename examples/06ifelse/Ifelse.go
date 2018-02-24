package main

import "fmt"

func main() {

	var flag = true
	if flag {
		fmt.Println("it is true")
	}

	if 4 % 2 == 0 {
		fmt.Println("4 is even")
	} else {
		fmt.Println("4 is even")
	}

	if 9 % 3 == 0 {
		fmt.Println("9 is deviced by 3")
	}

	if k := 7; k < 0 {
		fmt.Println("is negative")
	} else if k < 10 {
		fmt.Println(k, " has 1 digit")
	} else {
		fmt.Println(k, " has multiple digits")
	}

}
