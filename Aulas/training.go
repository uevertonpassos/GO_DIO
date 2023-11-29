package main

import (
	"fmt"
	"os"
	"strconv"
)

// Interface Calculator define as operações suportadas.
type Calculator interface {
	Add(float64, float64) float64
	Subtract(float64, float64) float64
	Multiply(float64, float64) float64
	Divide(float64, float64) (float64, error)
}

// BasicCalculator é uma implementação simples da interface Calculator.
type BasicCalculator struct{}

func (c BasicCalculator) Add(a, b float64) float64 {
	return a + b
}

func (c BasicCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c BasicCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c BasicCalculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero não é permitida")
	}
	return a / b, nil
}

func main() {
	var a, b float64
	var operator string

	fmt.Println("Calculadora em Go")
	fmt.Println("Operações disponíveis: +, -, *, /")
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)
	fmt.Print("Digite o operador: ")
	fmt.Scan(&operator)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	var calc Calculator
	calc = BasicCalculator{}

	var result float64
	var err error

	switch operator {
	case "+":
		result = calc.Add(a, b)
	case "-":
		result = calc.Subtract(a, b)
	case "*":
		result = calc.Multiply(a, b)
	case "/":
		result, err = calc.Divide(a, b)
		if err != nil {
			fmt.Println("Erro:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Operador inválido")
		os.Exit(1)
	}

	fmt.Printf("Resultado: %.2f\n", result)
}


//mudando
