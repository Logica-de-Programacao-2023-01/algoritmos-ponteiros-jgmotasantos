package main

import "fmt"

func inverterBool(ptrBool *bool) {
	*ptrBool = !(*ptrBool)
}

func main() {
	valor := true
	fmt.Println("Anteriormente o valor era: ", valor)
	inverterBool(&valor)
	fmt.Println("Apos chamar a funçao o valor ficou:", valor)

}
