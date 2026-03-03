package main

import "fmt"

var resultado, numero int

func main() {
	var opcao int

	for {
		fmt.Println("Número atual:", resultado)
		fmt.Println("Escolha um número:")
		_, err := fmt.Scanln(&numero)
		if err != nil {
			fmt.Println("Entrada inválida! Digite apenas números inteiros.")
			continue
		}

		fmt.Println("Escolha uma operação pelo indíce:")
		fmt.Println("1 => Soma")
		fmt.Println("2 => Subtração")
		fmt.Println("3 => Multiplicação")
		fmt.Println("4 => Divisão")
		fmt.Println("0 => Encerrar programa")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			resultado = soma(numero)
		case 2:
			resultado = subtrai(numero)
		case 3:
			resultado = multiplica(numero)
		case 4:
			resultado = divide(numero)
		case 0:
			fmt.Println("Resultado:", resultado)
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
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
