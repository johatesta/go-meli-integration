package main

import (
	"fmt"

	"github.com/johatesta/go-meli-integration/Tp5_Apis"
	"github.com/johatesta/go-meli-integration/Tp6_calculadora_hilos"
	"github.com/johatesta/go-meli-integration/Tp_HolaMundo"

	"github.com/johatesta/go-meli-integration/tp7_calculadora_test"
	"github.com/johatesta/go-meli-integration/tp_8_calculadoraApi"
)

func main() {
	var opcion int64
	fmt.Printf("Ingrese la opcion: \n")
	fmt.Printf("3): Hola Mundo \n")
	fmt.Printf("4): Calculadora \n")
	fmt.Printf("5): APIs en go \n")
	fmt.Printf("6): Calculadora hilos \n")
	fmt.Printf("7): Calculadora Testing \n")
	fmt.Printf("8): Api calcu \n\n ")
	fmt.Scan(&opcion)
	switch numeroTP {
	case 3:
		Tp_HolaMundo.HelloWord()
		break
	case 4:
		Tp4_calculdora.op()
		break
	case 5:
		Tp5_Apis.ConsumoApi()
		break
	case 6:
		Tp6_calculadora_hilos.CalculadoraHilos()
		break
	case 7:
		tp7_calculadora_test.CalculadoraTest()
		break
	case 8:
		tp_8_calculadoraApi.CalculadoraApi()
		break
	}
}
