package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	fmt.Println("\n╔══════════════════════════════════════════════════╗")
	fmt.Println("║     SISTEMA DE MENSAGERIA - TEMPO REAL          ║")
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Println("║  Status: CONECTADO                               ║")
	fmt.Println("║  Canal:  flamengo-vs-vasco                       ║")
	fmt.Println("╚══════════════════════════════════════════════════╝\n")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite o lance (ou 'sair'): ")

	for scanner.Scan() {
		input := scanner.Text()

		if strings.ToLower(input) == "sair" {
			fmt.Println("Encerrando...")
			break
		}

		if input != "" {
			// Parse do input: TIPO | MINUTO | DESCRIÇÃO
			parts := strings.Split(input, "|")

			var tipo, descricao string
			var minuto int

			if len(parts) == 3 {
				// Formato completo: TIPO | MINUTO | DESCRIÇÃO
				tipo = strings.TrimSpace(parts[0])
				minutoStr := strings.TrimSpace(parts[1])
				descricao = strings.TrimSpace(parts[2])

				minuto, err = strconv.Atoi(minutoStr)
				if err != nil {
					minuto = 0
				}
			} else if len(parts) == 2 {
				// Formato: TIPO | DESCRIÇÃO (minuto = 0)
				tipo = strings.TrimSpace(parts[0])
				descricao = strings.TrimSpace(parts[1])
				minuto = 0
			} else {
				// Formato simples: apenas descrição
				tipo = "lance"
				descricao = input
				minuto = 0
			}

			// Publicacao via servico
			err := fbService.PublishLance(tipo, descricao, minuto)
			if err != nil {
				fmt.Printf("Falha ao publicar: %v\n", err)
			} else {
				fmt.Printf("Lance publicado: [%s] %d' - %s\n", tipo, minuto, descricao)
			}
		}
		fmt.Print("\nDigite o lance (ou 'sair'): ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro de leitura:", err)
	}
}
