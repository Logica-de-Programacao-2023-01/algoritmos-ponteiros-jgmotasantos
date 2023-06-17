package main

import "fmt"

type Book struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Book) {
	desconto := livro.Preco * 0.1
	livro.Preco -= desconto
}

func main() {
	livro := Book{
		Titulo: "Aprenda Go",
		Autor:  "John Doe",
		Preco:  100.0,
	}

	fmt.Println("Preço original:", livro.Preco)

	aplicarDesconto(&livro)
	fmt.Println("Preço com desconto:", livro.Preco)
}
