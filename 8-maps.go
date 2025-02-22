package main

import "fmt"

func main() {
	idade := map[string]int{} // Criando MAP vazio e atribuindo valores depois
	idade["Rafaela"] = 25
	idade["Isabel"] = 31
	fmt.Println(idade)
	fmt.Println(idade["Rafaela"])
	fmt.Println(idade["Isabel"])

	anoNasc := map[string]int{ // Criando o MAP já com os valores
		"Rafaela": 2000,
		"Isabel": 1994,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Rafaela"])

	anoNasc["golangDoZero"] = 2024 // Adicionando elementos no MAP
	fmt.Println(anoNasc)
}

// 2 - Maps: Heterogêneos 
//Pode misturar Tipos
//estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro
// map[k]v -> k = chave, v = valor

// map[string]int
// { "Rafaela": 25, "Isabel": 31}
// map[string]string
// {"Rafaela": "Souza", "Isabel": "Nunes"}