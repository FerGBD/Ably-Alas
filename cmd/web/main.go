package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type PageData struct {
	AblyAPIKey string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado.")
	}

	apiKey := os.Getenv("ABLY_API_KEY")
	if apiKey == "" {
		log.Fatal("ABLY_API_KEY não encontrada no arquivo .env")
	}

	// Servir arquivos estáticos (CSS, JS, imagens, etc)
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rota principal - renderiza o template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplPath := filepath.Join("web", "templates", "index.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
			log.Printf("Erro ao carregar template: %v", err)
			return
		}

		data := PageData{
			AblyAPIKey: apiKey,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
			log.Printf("Erro ao renderizar template: %v", err)
		}
	})

	port := "8080"
	log.Printf("Servidor rodando em http://localhost:%s", port)
	log.Printf("Servindo arquivos estáticos de: web/static")
	log.Printf("Template: web/templates/index.html")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
