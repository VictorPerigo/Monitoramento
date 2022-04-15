package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibirMenu()
}

func exibirMenu() {
	for {
		fmt.Println("\nOque deseja fazer?")
		fmt.Println("1 - Iniciar monitoramento")
		fmt.Println("2 - Exibir os logs")
		fmt.Println("0 - Sair do programa:")
		menuFuncionamento()
	}
}

func menuFuncionamento() {
	switch inputOpcao() {
	case 1: // Iniciar monitoramento
		monitoramento()
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

func inputOpcao() int {
	opcao := 5
	// variavel opcao recebe o valor digitado pelo usuario, sem expecificar o tipo do input, o proprio go ira inferir o tipo apartir da variavel opcao
	fmt.Scan(&opcao)

	return opcao
}

func monitoramento() {
	fmt.Print("\033[H\033[2J")
	urls := []string{
		"https://www.google.com.br/jonas",
		"https://www.youtube.com",
	}

	for _, url := range urls {
		response, err := http.Get(url)

		if err != nil {
			fmt.Println("Erro ao acessar a url:", err)
		} else if response.StatusCode == 200 {
			fmt.Println("status code:", response.StatusCode, "OK  Site:", url)
			defer response.Body.Close()
		} else {
			fmt.Println("status code:", response.StatusCode, "ERR Site:", url)
			defer response.Body.Close()
		}
	}

}
