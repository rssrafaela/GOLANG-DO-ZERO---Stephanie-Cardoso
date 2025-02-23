package main 

import "fmt"

func main(){

// Exemplo 1
	posicao := 2
	switch posicao {
		case 1: 
			fmt.Println("Primeiro lugar")
		case 2:
			fmt.Println("Segundo lugar")
		case 3: 
			fmt.Println("Terceiro lugar")
	}

// Exemplo 2
	nome := "bob"
	switch nome {
	case "Rafaela":
		fmt.Println("É uma pessoa")
	case "bento":
		fmt.Println("É um cachorro")
	case "bob":
		fmt.Println("É um personagem")
	}
}