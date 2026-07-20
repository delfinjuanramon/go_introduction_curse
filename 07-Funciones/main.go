//  Funciones

package main

// para crear un modulo del proyecto, en la termnal escribimos:
//  \hithub.com\delfinjuanramon\go_introduction_curse\Go mod init <enter>

import (
	"fmt"

	"github.com/delfinjuanramon/go_introduction_curse/07-Funciones/functions"
)

func main() {
	v := functions.Add(3, 4)
	fmt.Println(v)

	functions.RepeatString(10, "as")

	fmt.Println()

	n, err := functions.Calc(functions.MUL, 5, 4)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Resultado: ", n)
	}

	n, err = functions.Calc(functions.DIV, 16, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Resultado: ", n)
	}

	fmt.Println()
	x, y := functions.Split(20)

	fmt.Println("Otro resultado: x=", x, " y=", y)

	vv := functions.MSum(23, 12, 32, 12, 3, 1, 2, 3, 2, 1, 23, 12, 1)
	fmt.Println("Suma con elipsis = ", vv)

	vv = functions.MSum()
	fmt.Println("Nueva suma sin parametros = ", vv)

	fmt.Println()	
	fmt.Println()

	fn:= functions.FactoryOperation(functions.SUB)
	vv= fn(2,3)
	fmt.Println("Resta del factory= ", vv)

	fn= functions.FactoryOperation(functions.MUL)
	vv= fn(2,3)
	fmt.Println("Multiplicacion del factory= ", vv)

}
