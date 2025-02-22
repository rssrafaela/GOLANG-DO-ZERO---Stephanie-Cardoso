package main

import "fmt"

func main(){
	//Variáveis
	// var + nome da varivel + tipo
	var nome string
	nome = "Rafaela"
	fmt.Println(nome)

	nome = "Isabel"
	fmt.Println(nome)

	var idade int
	idade = 25
	fmt.Println(idade)

	var a = "Rafaela"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

	// Constantes 
	const idadeRafaela = 25
	fmt.Println(idadeRafaela)
}