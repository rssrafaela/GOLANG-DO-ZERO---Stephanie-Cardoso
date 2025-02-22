package main

import "fmt"

func main() {
	// Chama a função soma passando 42 e 13
	fmt.Println(soma(42, 13))

	nome, sobrenome := printaNomeCompleto("RAFAELA", "SOUZA")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}

// Função começando com letra minúscula:
// Função é PRIVADA
// Função ela só pode ser utilizada no próprio pacote

// Caso coloque dois parametros, precisa ser colocado os parenteses.
func printaNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}


// Função começando com letra maiúscula:
// Função é PÚBLICA
// Função ela só pode ser utilizada fora do próprio pacote
// Como utilizaria ela fora do pacote : main.PrintaNome(nome)

func PrintaNome(nome string) string {
	return nome
}

// Função soma que retorna a soma de dois inteiros
func soma(x int, y int) int {
	return x + y
}


// Segunda opção:

// func main(){
//	somaDosValores := soma(42, 13)
//	fmt.Println(somaDosValores)
//}