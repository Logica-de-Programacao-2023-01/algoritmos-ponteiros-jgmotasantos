package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func alterarTituloLivro(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro Exemplo",
		Autor:  "Anônimo",
	}

	fmt.Println("Título original:", livro.Titulo)

	alterarTituloLivro(&livro)
	fmt.Println("Título alterado:", livro.Titulo)
}
