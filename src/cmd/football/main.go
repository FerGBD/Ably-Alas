package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"ably-go-quickstart/client/ably"
	"ably-go-quickstart/client/service"

	"github.com/joho/godotenv"
)

func main() {
	// Carregamento das variaveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env nao encontrado.")
	}

	// Inicializacao do cliente Ably isolado
	client, err := ably.NewClient()
	if err != nil {
		log.Fatalf("Erro ao inicializar cliente: %v", err)
	}
	defer client.Close()

	// Inicializacao do servico de dominio
	fbService := &service.FootballService{
		Channel: client.Channels.Get("flamengo-vs-vasco"),
	}

	// Inscricao no canal para receber mensagens em background
	fbService.SubscribeLances()

	fmt.Println("--------------------------------------------------")
	fmt.Println("SISTEMA DE MENSAGERIA - TEMPO REAL")
	fmt.Println("Status: CONECTADO")
	fmt.Println("Canal:  flamengo-vs-vasco")
	fmt.Println("--------------------------------------------------")

	// Loop de entrada para publicacao de mensagens (Jornalista)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite o lance (ou 'sair'): ")

	for scanner.Scan() {
		input := scanner.Text()

		if strings.ToLower(input) == "sair" {
			fmt.Println("Encerrando...")
			break
		}

		if input != "" {
			// Publicacao via servico (tipo, descricao, minuto)
			err := fbService.PublishLance("DADO_BRUTO", input, 0)
			if err != nil {
				fmt.Printf("Falha ao publicar: %v\n", err)
			} else {
				fmt.Println("Mensagem enviada.")
			}
		}
		fmt.Print("Digite o lance (ou 'sair'): ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro de leitura:", err)
	}
}
