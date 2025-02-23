package main

import (
	"fmt"
	"time"  // É necessário importar o pacote "time" para usar o Sleep
)

// Loops
// Laços de repitção
// Repetir tarefas

func main(){

	sum := 0
	// i++ -> soma 1 
	// i-- -> subtraindo 1
	for i := 0; i < 10; i++{
		fmt.Println("Sum:", sum)
		fmt.Println("Indice:", i)
		sum += i
		// é a mesma coisa que: sum = sum + i
		// sum -= i -> sum = sum -i
	}
	fmt.Println(sum)

	// // loop infinito, para forçar a parada aperta CTRL + C
	for {
		fmt.Println("Golang do zero")
		 	time.Sleep(2 * time.Second)
	} 

	// // for range
	frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}
