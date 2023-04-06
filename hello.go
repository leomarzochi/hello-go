package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 10
const delay = 2

func main() {
	showIntroduction()
	for {
		showMenu()

		comando := readCommand()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}

func showMenu() {
	fmt.Println(" ")
	fmt.Println("1. Iniciar monitoramento")
	fmt.Println("2. Exibir logs")
	fmt.Println("0. Sair")
}

func showIntroduction() {
	versao := 1.1
	fmt.Println("--- Bem vindo ao monitorador de sites")
	fmt.Println("--- Rodando na versão: ", versao)
}

func readCommand() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando selecionado: ", comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	sites := readFromFile()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site: ", i)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println(site, "está com problemas")
		registraLog(site, false)
	}
}

func readFromFile() []string {
	sites := []string{}
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro ao abrir arquivo", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
