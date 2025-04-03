package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b float64
	var operacion string
	var nombre string

	fmt.Println("Calculadora básica en Go")
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)

	// Verificar si se proporcionaron argumentos
	if len(os.Args) == 4 {
		// Intentar parsear argumentos
		a, _ = strconv.ParseFloat(os.Args[1], 64)
		operacion = os.Args[2]
		b, _ = strconv.ParseFloat(os.Args[3], 64)
	} else {
		fmt.Println("Operaciones disponibles: + (suma), - (resta), * (multiplicación), / (división)")

		fmt.Print("Ingrese primer número: ")
		fmt.Scanln(&a)

		fmt.Print("Ingrese operación (+, -, *, /): ")
		fmt.Scanln(&operacion)

		fmt.Print("Ingrese segundo número: ")
		fmt.Scanln(&b)
	}

	resultado, err := calcular(a, b, operacion)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("%s, el resultado es: %.2f\n", nombre, resultado)
}

func calcular(a, b float64, operacion string) (float64, error) {
	switch operacion {
	case "+":
		return sumar(a, b), nil
	case "-":
		return restar(a, b), nil
	case "*":
		return multiplicar(a, b), nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("división por cero")
		}
		return dividir(a, b), nil
	default:
		return 0, fmt.Errorf("operación no válida")
	}
}

func sumar(a, b float64) float64 {
	return a + b
}

func restar(a, b float64) float64 {
	return a - b
}

func multiplicar(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	return a / b
}
