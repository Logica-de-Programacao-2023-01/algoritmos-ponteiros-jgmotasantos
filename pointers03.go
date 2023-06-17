package main

import "fmt"

func reverseString(str *string) {
	runes := []rune(*str)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*str = string(runes)
}

func main() {
	s := "Hello, World!"
	fmt.Println("String original:", s)

	reverseString(&s)
	fmt.Println("String reversa:", s)
}
