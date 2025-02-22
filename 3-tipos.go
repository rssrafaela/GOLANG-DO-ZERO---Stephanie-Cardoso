package main

import "fmt"

func main(){
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	fmt.Printf("Type: %T - Value: %v\n", "rafa", "rafa")
	fmt.Printf("Type: %T - Value: %v\n",1, 1)
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)
}

//Tipos:
// bool (true/false)
// string - sequência de bytes
// int
// float (float64/gloat32) - decimal