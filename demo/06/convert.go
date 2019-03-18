package main

import "fmt"

type MyString = string
type MyString2 string
type MyString3 = MyString2

func (MyString3) t() {

}

func main() {
	fmt.Println(string(-1))

	var s1 MyString2
	s1 = "abc"
	switch interface{}(s1).(type) {
	case string:
		fmt.Println("string")
	case MyString2:
		fmt.Println("MyString2")
	}

	ok := string(s1)
	fmt.Println(ok)

}
