package main

import "fmt"

func modifyVariableValue(ptr *int) {
	*ptr = 42
}

func main() {
	var value int
	fmt.Println("Valor inicial:", value)

	modifyVariableValue(&value)
	fmt.Println("Valor modificado:", value)

	value = 0
}

