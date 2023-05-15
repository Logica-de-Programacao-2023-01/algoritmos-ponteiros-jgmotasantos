//Escreva uma função swap que receba dois ponteiros para int como argumentos e troque
//os valores apontados por eles.

package main

import "fmt"

func swap(ptr *int) {
	*ptr = *ptr * 2
}

func main() {
	var x int = 40
	var y int = 50

	fmt.Println("Antes de chamada a funçao o primeiro valor era: ", x)
	swap(&x)
	fmt.Println("Depois de chamada a funçao o primeiro valor ficou: ", x)

	fmt.Println("Antes de chamada a funçao o primeiro valor era: ", y)
	swap(&y)
	fmt.Println("Depois de chamada a funçao o primeiro valor ficou: ", y)

}
