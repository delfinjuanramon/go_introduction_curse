package main

import "fmt"

func main() {
	// Comillas dobles para def String y comillas Simples para def Caracteres
	var A byte = 'A'
	var a byte = 'a'
	var R byte = 82              // este es el ASCII de R
	var s byte = 115			// este es el ASCII de s
	vector:= []byte{65,97,82,115}

	fmt.Println()
	fmt.Println("Valor en ASII code:")
	fmt.Println(A)
	fmt.Println(a)
	fmt.Println(R)
	fmt.Println(s)
	fmt.Println((vector))
	
	fmt.Println()
	fmt.Println("Valor en String code:")
	fmt.Println(string(A))
	fmt.Println(string(a))
	fmt.Println(string(R))
	fmt.Println(string(s))
	fmt.Println(string(vector))

}	
