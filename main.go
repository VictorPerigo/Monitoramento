package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	exibirMenu()
}

var urls []string = getUrlsFromTxt()

// funcionamento dos menus:

func menuFuncionamento() {
	switch inputOpcao() {
	case 1: // Iniciar monitoramento
		monitoramento()
	case 2: // Exibir os logs
		fmt.Println("Exibindo os logs...")
	case 3: // Urls
		menuFuncionamentoUrls()
	case 0: // Sair do programa
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default: // Caso o usuario digite uma opcao invalida
		fmt.Println("Opção invalida")
		os.Exit(-1)
	}
}

func menuFuncionamentoUrls() {
	exibirUrls()
	exibirMenuUrl()

	switch inputOpcao() {
	case 1: // Adicionar url
		_ = selecionarUrl()
	case 2: // Remover url
		removerUrl()
	case 0: // Sair do programa
		fmt.Println("Saindo do programa...")
		fmt.Println("\033[H\033[2J")
	default: // Caso o usuario digite uma opcao invalida
		fmt.Println("Opção invalida")
	}
}

// exibicoes:

func exibirUrls() {
	urls = getUrlsFromTxt()
	limparTerminal()
	fmt.Println("Urls:")
	if urls == nil {
		fmt.Println("Nenhuma url cadastrada")
	}
	for index, url := range urls {
		fmt.Println("index:", index, ", url:", url)
	}
}

func exibirMenuUrl() {
	fmt.Println("\nOque deseja fazer?")
	fmt.Println("1 - Adicionar url")
	fmt.Println("2 - Remover url")
	fmt.Println("0 - Voltar ao menu:")
}

func exibirMenu() {
	for {
		fmt.Println("\nOque deseja fazer?")
		fmt.Println("1 - Iniciar monitoramento")
		fmt.Println("2 - Exibir os logs")
		fmt.Println("3 - Urls")
		fmt.Println("0 - Sair do programa:")
		menuFuncionamento()
	}
}

// funcionamento das urls:

func adicionarUrl(url string) error {
	file, err := os.OpenFile("urls.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return err
	}
	defer limparTerminal()
	defer file.Close()

	if _, err = file.WriteString(url + "\n"); err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return err
	}
	urls = append(urls, url)
	return nil
}

func selecionarUrl() error {
	var url string

	fmt.Println("Digite a url:")
	fmt.Scan(&url)

	if err := adicionarUrl(url); err != nil {
		return err
	}

	return nil
}

func removerUrl() {
	defer limparTerminal()
	exibirUrls()
	fmt.Println("\nDigite o index da a url que deseja remover:")
	var indexEscolhido int
	fmt.Scan(&indexEscolhido)
	if err := os.Truncate("urls.txt", 0); err != nil {
		fmt.Println("erro ao recriar arquivo", err)
	}
	urls = append(urls[:indexEscolhido], urls[indexEscolhido+1:]...)
	for _, url := range urls {
		adicionarUrl(url)
	}
}

func getUrlsFromTxt() []string {
	var sites []string
	file, err := os.Open("urls.txt")

	if err != nil {
		os.Create("urls.txt")
	} else {
		leitor := bufio.NewReader(file)
		defer file.Close()
		for {
			linha, err := leitor.ReadString('\n')
			linha = strings.TrimSpace(linha)
			if linha != "" {
				sites = append(sites, linha)
			}
			if err == io.EOF {
				break
			}

		}
	}

	return sites
}

// funcionamento das opcoes:

func monitoramento() {
	limparTerminal()
	if urls == nil {
		fmt.Println("Nenhuma url cadastrada")
		return
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

// ferramentas:

func inputOpcao() int {
	opcao := 999
	// variavel opcao recebe o valor digitado pelo usuario, sem expecificar o tipo do input, o proprio go ira inferir o tipo apartir da variavel opcao
	fmt.Scan(&opcao)

	return opcao
}

func limparTerminal() {
	fmt.Print("\033[H\033[2J")
}
