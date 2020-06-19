package tp6_calculadora_hilos

import (
	"fmt"
	"time"
)

type funciones struct {
	Desvia    bool
	salida    int
	Operandos []float64
	Operacion string
	Resultado float64
}

var c1 chan funciones = make(chan funciones)
var c2 chan funciones = make(chan funciones)
var c3 chan funciones = make(chan funciones)
var c4 chan funciones = make(chan funciones)

func CalculadoraHilos() {

	var num1, num2 float64
	var operaciones []float64
	var respuesta string
	var desvia bool
	var salida int

	fmt.Printf("Ingrese los numeros con los que va a trabajar: \n")
	fmt.Scanln(&num1, &num2)
	operaciones = []float64{num1, num2}
	fmt.Printf("Quiere desviar la calculadora?: (S/N)")
	fmt.Scanln(&respuesta)
	if respuesta == "S" {
		desvia = true
		fmt.Printf("Ingrese el numero de la desviacion:")
		fmt.Scanln(&salida)
	}

	go func() {

		resultado := operaciones[0] + operaciones[1]
		if desvia {
			resultado += float64(des)
		}
		c1 <- funciones{
			Desvia:=    des,
			salida:=   out,
			Operandos:= operandos,
			Operacion:= "suma",
			Resultado:= resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		resultado := operaciones[0] - operaciones[1]
		if desvia {
			resultado += float64(des)
		}
		c2 <- funciones{
			Desvia:=    des,
			salida:=  out,
			Operandos:= operandos,
			Operacion:= "resta",
			Resultado:= resultado,
		}
		time.Sleep(time.Second * 2)
	}()

	go func() {

		resultado := operaciones[0] * operaciones[1]
		if desvia {
			resultado += float64(des)
		}
		c3 <- funciones{
			Desvia:=   des,
			salida:=   out,
			Operandos:= operandos,
			Operacion:="multiplicacion",
			Resultado:= resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {

		resultado := operaciones[0] / operaciones[1]
		if desvia {
			resultado += float64(des)
		}
		c4 <- funciones{
			Desvia:=   des,
			salida:=   out,
			Operandos:= operandos,
			Operacion:= "division",
			Resultado:= resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		for {
			select {
			case suma := <-c1:
				fmt.Println(suma)
			case resta := <-c2:
				fmt.Println(resta)
			case multiplicar := <-c3:
				fmt.Println(multiplicar)
			case dividir := <-c4:
				fmt.Println(dividir)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
