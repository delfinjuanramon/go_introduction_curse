//   Structs

package main

import (
	"fmt"
	"github.com/delfinjuanramon/go_introduction_curse/08-Structs/strucks"
)

func main() {
	var p1 structs.Product

	p1.ID= 12
	p1.Name= "Prueba"
	fmt.Println(p1)

	p2:= structs.Product {}
	p2.ID= 25
	p2.Name= "Juan"
	fmt.Println(p2)

              // no es una buena practica, pero si se usa, hay que poner todos los campos
	p3:= structs.Product {6, "es el p3", "A",  45.50}
	fmt.Println(p3)

	p4:= structs.Product {  // esta es la forma recomendada, y no es necesario poner todos los campos
		ID: 78,
		Name: "Torcuata",
		Type: "B",              
	}
	fmt.Println(p4)

	np:= structs.NewProduct{
		Name: "Heladeras marca Saras",
		Type: structs.Tipo {
			Code: "A",
			Description: "Electrodomestico",
		},

		Tags: []string{"heladera", "sarasa", "freezer", "refrigerador" },
	}
	fmt.Println(np)
}
