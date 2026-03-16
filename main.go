package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Entre com a operação completa nesse formato -> 2*2")
	var input string
	fmt.Scan(&input)

	operacao := strings.Split(input, "")
	fmt.Println(operacao)

	resultado := getResult(operacao)
	fmt.Println("Resultado de", operacao[0], operacao[1], operacao[2], "=", resultado)
}

func getResult(operacao []string) int {
	num1, _ := strconv.Atoi(operacao[0])
	num2, _ := strconv.Atoi(operacao[2])

	switch operacao[1] {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("Operador não válido. (Operadores válidos -> | + | - | * | / |)")
	}
}

// Operações Soma (+), Subtração (-), Multiplicação (*), Divisão (/)
