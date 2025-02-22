package main

import "fmt"

func main() {

	// Arrays
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:3]) // Pega o valor até antes da posição 3.

	// Slices
	var slice []string // slice não tem tamanho
	slice = make([]string, 5) // determina um valor inicial, GO deixa adicionar mais elementos.
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)
	fmt.Println(slice[2]) // Atribui vazio
	fmt.Println(len(slice)) // Tamanho do slice

	numPares := []int{2, 4, 6, 8, 10, 12} // Criando o slice já com os valores
	fmt.Println(numPares)

	numPares = append(numPares, 14, 16) // adicionando itens ao slice
	fmt.Println(numPares)

}
// LISTAS

// 1 - Arrays e Slices: Homogêneos 
// todos os elementos tem o mesmo tipo
// [1, 2, 3,, 4, 5, 6] - []int
// ["rafaela", "isabel", "golang"] - []string
// [1.001, 2.002, 3.003, 4.004, 5.005, 6.006] - []float64

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ser outro
// map[string]int
// {"rafaela": 25, "isabel": 31}
// map [string]string
// {"Rafaela": "Souza", "Isabel": "Nunes"}

// Array

// Tamanho fixo, de zero ou mais elemtnos do mesmo tipo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho do array
// Por conta do tamanho fixo, não é tão usado. Só em casos específicos

// Slice

// Tipo o array, mas sem tamanho fixo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho do array
// Função append() usada para adicionar valores no slice