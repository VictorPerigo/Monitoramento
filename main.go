package main

import (
	"fmt"
	"os"
)

func main() {
	exibirMenu()
	menuFuncionamento()
}

func inputOpcao() int {
	opcao := 5
	// variavel opcao recebe o valor digitado pelo usuario, sem expecificar o tipo do input, o proprio go ira inferir o tipo apartir da variavel opcao
	fmt.Scan(&opcao)

	return opcao
}

func exibirMenu() {
	fmt.Println("Olá, que deseja fazer?")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa:")
}

func menuFuncionamento() {
	switch inputOpcao() {
	case 1: // Iniciar monitoramento
		fmt.Println("Monitorando...")
	case 2: // Exibir os logs
		fmt.Println("Exibindo os logs...")
	case 0: // Sair do programa
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default: // Caso o usuario digite uma opcao invalida
		fmt.Println("Opção invalida")
		os.Exit(-1)
	}
}
