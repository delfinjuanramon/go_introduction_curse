// Vectores y Slices

package main

import ( 
	"fmt"
	"unsafe"
)

// En los vectores, todos los elementos son del mismo tipo. Son Arrays
// los vectores no pueden cambiar de tamaño. Son fijos

func main(){
	var miInt int
			
	fmt.Println(miInt)
	fmt.Printf("Tipo:  %T, bytes: %d, bits %d\n", miInt, unsafe.Sizeof(miInt) , unsafe.Sizeof(miInt)*8)

	var miArray [6]int      // Los arrays son vectores. Tienen un tamaño fijo no se pueden cambiar
	fmt.Println(miArray)

	miArray[1]= 2
	miArray[2]= 5
	miArray[3]= 9
	fmt.Println(miArray)
	fmt.Println("miArray Tamaño: ", unsafe.Sizeof(miArray))

	// Slices

	var slice1 []int
	fmt.Printf("Slice1.  Tamaño: %d, Valor: %v\n", len(slice1), slice1)

	slice1 = append(slice1, 10,20, 30, 40,50)
	fmt.Printf("Slice1.  Tamaño: %d, Valor: %v\n", len(slice1), slice1)

    otroSlice:= miArray[1:3]    // Toma elemento 1 hasta el 3-1 o sea el elemento2
	fmt.Printf("otroSlice.  Tamaño: %d, Valor: %v\n", len(otroSlice), otroSlice)

	fmt.Println()
	fmt.Println("dir miArray[2]:  ", &miArray[2] )
	fmt.Println("dir otroSlice[0]:",  &otroSlice[0])
	
	slice3:= make([]int, 3)
	fmt.Println()
	fmt.Printf("slice3:  Tamaño: %d, Valor: %v\n", len(slice3), slice3)

	// Maps
	//      Tienen la estructura: lklave, valor.   Las llaves son unicas
	

/*	
	var miArreglo [6]int
	fmt.Println(miArreglo)
	
	otroArray:= [3]string{"valor1", "valor2", "valor3"}
	fmt.Println(otroArray)

	miArreglo[1]= 2
	miArreglo[2]= 5
	miArreglo[3]= 9
	miArreglo[4]= 11
	miArreglo[5]= 15
	fmt.Println(miArreglo)
	fmt.Println()

	//   SLICES
	// POdemos tomar en un slice una porcion de un vector. Si lo cambiamos en los slices, se cambira tambien en el original

	var slice1 []int
	fmt.Printf("tamaño Slice1: %d, valor: %v\n", len(slice1), slice1)
	fmt.Println()	
	slice1= append(slice1, 10,20,30,40,50)
	fmt.Printf("tamaño Slice1: %d, valor: %v\n", len(slice1), slice1)

	slice2:= slice1[2:4]
	fmt.Printf("tamaño Slice2: %d, valor: %v\n", len(slice2), slice2)

	slice3:= miArreglo[1:4]     // ojo empieza en cero y hasta el 4-1= 3 elementos
	fmt.Printf("tamaño Slice3: %d, valor: %v\n", len(slice3), slice3)

	fmt.Println(&miArreglo[2], "valor: &d", miArreglo[2] )  //direccion de memoria
	fmt.Println(&slice3[1], "valor: &d", slice3[1])         //direccion de memoria

	fmt.Println("miarreglo: ", miArreglo)
	slice3[0]= 80
	slice3[1]=100

	fmt.Println("Ahora miarreglo: ",miArreglo)  // lo modifico el slice

	miArreglo[5]= 58
	fmt.Println("otra vez miarreglo: ",miArreglo)  // modificado a mano

	fmt.Println("miarreglo parcial: ", miArreglo[2:])  // desde 2 a donde acabe

*/	

}