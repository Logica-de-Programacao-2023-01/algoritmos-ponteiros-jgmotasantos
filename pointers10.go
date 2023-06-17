package main

import "fmt"

func preencherPrimos(slice *[]int, tamanho int) {
	if tamanho <= 0 {
		return
	}

	numerosPrimos := []int{}
	numero := 2

	for len(numerosPrimos) < tamanho {
		if isPrimo(numero) {
			numerosPrimos = append(numerosPrimos, numero)
		}
		numero++
	}

	*slice = numerosPrimos
}

func isPrimo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	tamanho := 10
	slice := []int{}
	preencherPrimos(&slice, tamanho)

	fmt.Printf("Os primeiros %d números primos são: %v\n", tamanho, slice)
}
