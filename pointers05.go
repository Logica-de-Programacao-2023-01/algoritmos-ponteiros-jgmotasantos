package main

import (
	"fmt"
	"math"
)

func atualizarValorDaVariavel(value *float64) {
	*value = (*value + math.Pi) / 2
}

func main() {
	num := 3.14
	fmt.Println("Valor inicial:", num)

	atualizarValorDaVariavel(&num)
	fmt.Println("Novo valor:", num)
}
