package main

import "fmt"

func updateVariableValue(value *int) {
	lastTwoDigits := *value % 100
	sum := (lastTwoDigits / 10) + (lastTwoDigits % 10)
	*value = sum
}

func main() {
	num := 1234
	fmt.Println("Valor inicial:", num)

	updateVariableValue(&num)
	fmt.Println("Novo valor:", num)
}
