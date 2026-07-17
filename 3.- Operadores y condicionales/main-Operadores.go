// Programa para estudiar Operadores y Comparaciones

package main

import(
	"fmt"
)

func main(){
	
	yearsOld:= 32

	fmt.Println("Operadores")
	fmt.Println(yearsOld > 30)     // true
	fmt.Println(yearsOld < 32)     // false
	fmt.Println(yearsOld <= 32)    // true
	fmt.Println(yearsOld >= 40)    // false
	fmt.Println(yearsOld == 32)    // true
	fmt.Println(yearsOld == 30)    // false

	fmt.Println()                    // Siempre se evalua primero los And y despues los Or
	fmt.Println(yearsOld == 32 || yearsOld == 30 )   //El operador OR es ||
	fmt.Println(yearsOld == 32 && yearsOld == 30 )   //El operador AND es &&
	
	fmt.Println()
	fmt.Println(true)
	fmt.Println(!true)		//El operador NOT es !

	edad:= 18

	if edad > 18{
		fmt.Println("%d es mayor que 18\n", edad)
	} else if edad < 18 {
		fmt.Println("%d es menor que 18\n", edad)
	} else {
		fmt.Println("%d es Igual que 18\n", edad)
	}

	edad=4
	switch edad {
	case 1:						// en el caso de que edad valga 1
		fmt.Printf("Uno") 
	case 2:
		fmt.Printf("Dos") 
	case 3:
		fmt.Printf("Tres")      // en el caso de que edad valga 3
	case 4:
		fmt.Printf("Cuatro") 
	default:
		fmt.Printf("Otro numero") 
	}


}

