// Uso de Maps

package main

import "fmt"

// Los mapas tienen una llave y un valor en cada elemento. Las llaves son unicas


func main(){

	mapa1:= make(map[int]string)   // Tipo de dato llave:int Tipo de valor: string
	mapa1[10]= "A"                 
	mapa1[20]= "B"
	mapa1[30]= "C"
	mapa1[74]= "D"
	mapa1[56]= "E"

	fmt.Println(mapa1)          
	fmt.Println(mapa1[10])   // este es la key 10, no el indice
	fmt.Println(mapa1[0])   // ojo este no existe porque no hay llave=0

	delete(mapa1,30)         // borra el elemento que tiene la llave 30	
	fmt.Println(mapa1)  

	mapa1[10]= "AZ"
	fmt.Println(mapa1)  

	mapa1[20]= ""
	fmt.Println(mapa1[20])    // tiene el valor de vacio: "". No muestra nada
	fmt.Println(mapa1[99])    // la llave 99 no exite. No muestra nada

	valor, bnd := mapa1[20]
	fmt.Println("El valor:",valor, "Bandera: ", bnd)  // Bnd da True porque Si existe
	valor1, bnd1 := mapa1[99]
	fmt.Println("El valor:",valor1, "Bandera: ", bnd1)  // Bnd da falso por que no existe

	mapa2:= map[byte]string{
		'F': "La efe",
		'C':  "la Ce",
		'J':  "La J",
	}
	fmt.Println("El mapa2:", mapa2)  // presenta los bytes por su ASCII


/*
	mapa1:= make(map[int]string)   // Tipo de dato llave:int Tipo de valor: string
	mapa1[10]= "A"                 
	mapa1[20]= "B"
	mapa1[30]= "C"
	mapa1[74]= "D"
	mapa1[56]= "E"

	fmt.Println(mapa1)          
	fmt.Println(mapa1[10])   // este es la key 10, no el indice
	fmt.Println(mapa1[0])   // ojo este no existe porque no hay llave=0

	delete(mapa1,30)         // borra el elemento que tiene la llave 30	
	fmt.Println(mapa1)  
	
	mapa1[10]= "AZ"
	fmt.Println(mapa1)  

	mapa1[20]= ""
	fmt.Println(mapa1[20])    // tiene el valor de vacio: "". No muestra nada
	fmt.Println(mapa1[99])    // la llave 99 no exite. No muestra nada

	valor, bnd := mapa1[20]
	fmt.Println("El valor:",valor, "Bandera: ", bnd)  // Bnd da True porque Si existe
	valor1, bnd1 := mapa1[99]
	fmt.Println("El valor:",valor1, "Bandera: ", bnd1)  // Bnd da falso por que no existe

	mapa2:= map[byte]string{
		'F': "La efe",
		'C':  "la Ce",
		'J':  "La J",
	}
	fmt.Println("El mapa2:", mapa2)  // presenta los bytes por su ASCII
	
*/
}