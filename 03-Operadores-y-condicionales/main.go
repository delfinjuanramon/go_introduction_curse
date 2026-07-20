// Programa para resolver tarea: Bases de Golang

package main

import (
	"fmt"

)

func main(){
	licencia:= true
	edad:= 18

	if edad == 18 && licencia {
		fmt.Println("Puede seguir avanzando")
	} else if edad <= 15 || !licencia {
		fmt.Println("No puede seguir circulando")
	}


}