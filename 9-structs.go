package main

import "fmt"

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipos diferentes

// type <nome da nossa estrutura> struct { <campos> }
type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Pessoa // herdando a struct Pessoa
	Tipo   string
}

func main() {
	// Inicializando pessoas com valores específicos
	fmt.Println(Pessoa{"Rafaela", 25})
	fmt.Println(Pessoa{Nome: "Isabel", Idade: 31})
	fmt.Println(Pessoa{Nome: "Rafaela"}) // O campo idade aparece com valor zero, pois não foi atribuído um valor explicitamente.

	// Criando uma instância de Pessoa e acessando seus campos
	p1 := Pessoa{Nome: "Bob", Idade: 2}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	// Alterando o campo Idade de p1
	p1.Idade = 3
	fmt.Println(p1.Idade)

	// Criando outro exemplo de pessoa
	p2 := Pessoa{Nome: "Patrick", Idade: 2}

	// Slice de Pessoas | lista do tipo Pessoa
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// Criando um map de alunos com listas de pessoas
	alunos := map[string][]Pessoa{} // estrutura chave (string) - valor (lista de pessoas)
	alunos["programação"] = pessoas
	fmt.Println(alunos)

	// Criando e inicializando um map de alunos com valores diretos
	alunos2 := map[string][]Pessoa{
		"Programação": {{Nome: "Rafaela", Idade: 25}, {Nome: "Bento", Idade: 4}},
		"Engenharia":  {{Nome: "Rafaela2", Idade: 25}, {Nome: "Bento2", Idade: 4}},
	}
	fmt.Println(alunos2)

	// Struct herdando campos de outra struct
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)  // Acessando o campo Nome da struct Pessoa
	fmt.Println(prof.Pessoa.Idade) // Acessando o campo Idade da struct Pessoa
	fmt.Println(prof.Tipo)         // Acessando o campo Tipo da struct Profissao
}



