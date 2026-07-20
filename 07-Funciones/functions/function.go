package functions

// el nombre con Mayuscula para indicar que es Publica
//     con minuscaula es que el alcance es privado

// OJO  todos los archivos que estan dentro de un mismo package,
//  es como si en realidad fuera un solo archivo. UNa variable local 
// puede leerse en otro archivo de ese package sin ningun problema.

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(veces int, valor string) {
	for i := 1; i < veces; i++ {
		fmt.Print(valor)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case MUL:
		return x * y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("División entre cero")
		}
		return x / y, nil
	}
	return 0, errors.New("Ha habido un error")
}

func Split(v int) (int, int) {
	x := v * 4 / 9
	y := v - x
	return x, y
}

// Para hacer el numero de parametros variables, pero del mismo tipo, se usa ...
// se puede usar varios parametros (str: string, v2...float64).
// los ... siempre van al final

func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}
