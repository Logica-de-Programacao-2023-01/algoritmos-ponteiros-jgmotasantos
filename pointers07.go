package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValorConta(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	conta := Conta{
		Saldo:   1000.0,
		Titular: "João",
	}

	fmt.Println("Saldo atual:", conta.Saldo)

	valorAdicional := 500.0
	adicionarValorConta(&conta, valorAdicional)

	fmt.Println("Saldo atualizado:", conta.Saldo)
}
