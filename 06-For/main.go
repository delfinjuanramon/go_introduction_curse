// Ciclos:  FOR   Ojo no existe While ni Repeat

package main

import "fmt"


func main(){

	sum:= 0
	for i:= 0; i< 10; i++{
		sum++
	}
	fmt.Println("la suma es:", sum)

	sum= 1
	for sum < 1000 {
		sum++
	}
	fmt.Println("la suma nueva es:", sum)

	sum=0
	for {
		if sum > 100 {
			break
		}
		sum++		
	}
	fmt.Println("la suma con break es:", sum)

	miArreglo:= [6]int{2,4,6,8,10,12}
	fmt.Println()

	for i:= range miArreglo {
			fmt.Println("indice:", i, " Valor:", miArreglo[i])
	}
    fmt.Println()

	for i, valor:= range miArreglo {           // aqui tre por indice
			fmt.Println("indice:", i, " Valor:", valor)
	}
    fmt.Println()

	for _, valor:= range miArreglo {         // ojo tiene _ en lugar de i
			fmt.Println( " Valor:", valor)
	}
    fmt.Println()

	mapa1:= map[string] float64{
		"A": 12.3,
		"Pi": 3.1416,
		"E": 2.71,
		"J": 99.999,
	}

	for key, valor:= range mapa1 {           // aqui trae por clave o llave
			fmt.Println("Key:", key, " Valor:", valor)
	}
    fmt.Println()

	for z, v:= range mapa1 {                 // igual al anterior
			fmt.Println("y:", z, " Valor:", v)
	}
    fmt.Println()


	mapa3:= map[string][]int{
		"A": nil,
		"B": {2,34,1,9,4},
		"C": {4,5,2,1},
	}

	for key, valor:= range mapa3{
		fmt.Println("Llave: ", key)
		for _, v:= range valor{
			fmt.Println("Valor: ", v)
		}
		fmt.Println()
	}


}