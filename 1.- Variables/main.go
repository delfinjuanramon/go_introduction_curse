package main

import (
	"fmt"
	"unsafe"
	"strconv"
)


func main() {        // Variables y Constantes
	var miIntVar  int
	miIntVar = -12
	fmt.Println("Mi variable entera es: ", miIntVar)

	var miIntVarPos uint
	miIntVarPos= 12
	fmt.Println("Mi variable positiva entera es: ", miIntVarPos )

	var miStringVar string
	miStringVar = "Hola"
	fmt.Println("Mi variable tipo String vale: ", miStringVar)
	
	var miBooleanVar bool
	miBooleanVar = true;
	fmt.Println("Mi variable booleana vale: ", miBooleanVar)

	fmt.Println("La direccion de memoria de la variable miStringVar es: ", &miStringVar)

	miIntVar2 := 12      //para declararla e instanciarla al mismo tiempo se usa :=
	fmt.Println("miIntVar2 vale: ", miIntVar2) 

	const miConstante1 = "A12345"
	fmt.Println("Mi vconstante de string 1 vale: ", miConstante1)

	const miConstInt int = 6789
	fmt.Println("la constante miConstInt vale: ", miConstInt)

	// fmt,printf   da formato a los valores

	fmt.Println("")

	var mi16BitsVar  int16
	
	fmt.Printf("Variable:  Tipo: %T, bytes: %d, bits: %d\n", 	miIntVar , unsafe.Sizeof(miIntVar), 8 * unsafe.Sizeof(miIntVar))

	fmt.Printf("Variable:  Tipo: %T, bytes: %d, bits: %d\n", mi16BitsVar, unsafe.Sizeof(mi16BitsVar), 8 * unsafe.Sizeof(mi16BitsVar))

	fmt.Println("")

	var miVarFloat32 float32
	var miVarFloat64 float64

	fmt.Printf("Variable:  Tipo: %T, bytes: %d, bits: %d\n", 	miVarFloat32 , unsafe.Sizeof(miVarFloat32), 8 * unsafe.Sizeof(miVarFloat32))

	fmt.Printf("Variable:  Tipo: %T, bytes: %d, bits: %d\n", miVarFloat64, unsafe.Sizeof(miVarFloat64), 8 * unsafe.Sizeof(miVarFloat64))

	fmt.Println("")

	{       // Simplemente con poner asi, se define un scope o alcance y la variables 
			// definidas en el, solo viven dentro de el
		var MiIntVar34 int
		fmt.Println("Mi variable entera es: ", MiIntVar34)
	}
	// Si trato de usar la variable fueras, la marca como indefinida
	// fmt.Println("Mi variable entera es: ", MiIntVar34)

    // cad= fmt.Sprintf(MiIntVar34)  convierte el string a la variable 


    MiIntVar99, err := strconv.ParseInt("1234567", 0, 64)  // desde bit 0 y de 6a bits
	fmt.Println("Error: ", err)
	fmt.Println(MiIntVar99)
	fmt.Println("La variable MiIntVar99 tiene:  Tipo: %T, Valor: %d\n", MiIntVar99, MiIntVar99)

}



