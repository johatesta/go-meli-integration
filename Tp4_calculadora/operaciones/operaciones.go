package operaciones

import (
	"errors"
	"fmt"
	"os"
)

type operandos struct {
	suma           float64
	resta          float64
	división       float64
	multiplicacion float64
}

func sum(a, b float64) float64 {
	return a + b
}

func rest(a, b float64) float64 {
	return a - b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("No se puede realizar una divisón por 0")
	}
	return a / b, nil

}

func mul(a, b float64) float64 {
	return a * b
}

func Operaciones(a, b float64) operandos {
	DIVI, err := div(a, b)
	if err != nil {
		fmt.Printf("Error: ", err.Error())
		os.Exit(1)
	}
	return operandos
	{
		suma := sum(a, b),
		resta := rest(a, b),
		division := DIVI,
		multiplicacion := mul(a, b),
	}
}

func desvia(o operandos, a float64) operandos {
	o.suma += a
	o.resta += a
	o.division += a
	o.multiplicacion += a

	return o
}
