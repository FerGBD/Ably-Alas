BINARY_NAME=tempo-real-futebol
MAIN_PATH=cmd/football/main.go
WEB_PATH=cmd/web/main.go

.PHONY: all setup run build clean web publisher

all: setup web

setup:
	go mod tidy
	go get -u github.com/ably/ably-go/ably
	go get github.com/joho/godotenv

# Inicia o servidor web (para espectadores)
web:
	@echo "Iniciando servidor web em http://localhost:8080"
	@go run $(WEB_PATH)

# Inicia o publicador de lances (jornalista)
publisher:
	@echo "Iniciando publicador de lances..."
	@go run $(MAIN_PATH)

# Roda ambos em terminais separados
run-both:
	@echo "Abra outro terminal e execute: make publisher"
	@make web

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)
	go build -o web-server $(WEB_PATH)

clean:
	rm -f $(BINARY_NAME) web-server