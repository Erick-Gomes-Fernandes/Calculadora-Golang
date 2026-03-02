package main

import "fmt"

var resultado int

func main() {
	fmt.Println("Resultado:", resultado)
	resultado = soma(5)
	fmt.Println("Resultado:", resultado)
	resultado = subtrai(3)
	fmt.Println("Resultado:", resultado)
	resultado = multiplica(4)
	fmt.Println("Resultado:", resultado)
	resultado = divide(2)
	fmt.Println("Resultado:", resultado)
}

func soma(x int) int {
	return (resultado + x)
}

func subtrai(x int) int {
	return (resultado - x)
}

func multiplica(x int) int {
	return (resultado * x)
}

func divide(x int) int {
	return (resultado / x)
}

// Operações Soma (+), Subtração (-), Multiplicação (*), Divisão (/)
// Rodar enquanto quiser rodar, dar opção de finalizar o programa ao usuário
