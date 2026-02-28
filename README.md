#  Ably-Alas Football Match - Transmissão de Lances em Tempo Real

Sistema de transmissão de lances de futebol em tempo real utilizando **Ably** para mensageria pub/sub e **Go** no backend.

## Sobre o Projeto

Este projeto demonstra o uso da plataforma **Ably** para criar um sistema de transmissão de eventos de partidas de futebol em tempo real. O sistema permite que jornalistas publiquem lances (gols, cartões, substituições, etc.) que são instantaneamente transmitidos para todos os espectadores conectados via web.

### Funcionalidades

- Transmissão em tempo real usando Ably Realtime
- Interface web responsiva para visualização dos lances
- Suporte a diferentes tipos de eventos (gol, cartão, substituição, etc.)
- Registro de minuto e autor de cada lance


##  Arquitetura

```
.
├── cmd/
│   ├── football/       # CLI para publicar lances
│   └── web/           # Servidor web HTTP
├── client/
│   ├── ably/          # Cliente Ably
│   ├── models/        # Modelos de dados
│   └── service/       # Lógica de negócio
├── web/
│   ├── static/        # Arquivos estáticos (CSS, JS)
│   └── templates/     # Templates HTML
├── .gitignore
├── go.mod
├── go.sum
└── Makefile
```


### Dependências

- `github.com/ably/ably-go` - SDK oficial do Ably para Go
- `github.com/joho/godotenv` - Gerenciamento de variáveis de ambiente

## Instalação

### Pré-requisitos

- Go 1.23.4 ou superior
- Conta no [Ably](https://ably.com/)


### Passo a Passo

1. **Clone o repositório**
```bash
git clone <url-do-repositorio>
cd Ably-Alas
```

2. **Configure as variáveis de ambiente**

Crie um arquivo `.env` na raiz do projeto:

```env
ABLY_API_KEY=your_ably_api_key_here
```

> 💡 Para obter sua API Key, acesse [Ably Dashboard](https://ably.com/dashboard) → Crie um app → Copie a API Key

3. **Instale as dependências**
```bash
go mod download
```

4. **Build do projeto**
```bash
# Usando Make
make build

# Ou manualmente
go build -o bin/football cmd/football/main.go
go build -o bin/web cmd/web/main.go
```

## Como Usar

### 1. Inicie o Servidor Web

```bash
# Usando Make
make run-web

# Ou manualmente
go run cmd/web/main.go
```

Acesse: http://localhost:8080

### 2. Publique Lances via CLI

Em outro terminal:

```bash
# Usando Make
make run-football

# Ou manualmente
go run cmd/football/main.go
```

### 3. Formatos de Input

O sistema aceita três formatos de entrada:

**Formato completo:**
```
TIPO | MINUTO | DESCRIÇÃO
```
Exemplo:
```
gol | 23 | Gabigol abre o placar para o Flamengo!
```

**Formato sem minuto:**
```
TIPO | DESCRIÇÃO
```
Exemplo:
```
cartao-amarelo | Everton Ribeiro recebe amarelo
```

**Formato simples:**
```
DESCRIÇÃO
```
Exemplo:
```
Bola na trave do Vasco!
```

### Tipos de Lances Suportados

- `gol` - Gol marcado
- `cartao-amarelo` - Cartão amarelo
- `cartao-vermelho` - Cartão vermelho
- `substituicao` - Substituição de jogador
- `impedimento` - Lance impedido
- `lance` - Lance genérico

## Comandos Make Disponíveis

```bash
make build          # Compila os binários
make run-web        # Executa o servidor web
make run-football   # Executa o CLI de publicação
make clean          # Remove os binários
make test           # Executa os testes
```


## API Ably

O projeto utiliza o canal: `flamengo-vs-vasco`

**Publicação:**
```go
channel.Publish(context.Background(), "lance", jsonData)
```

**Subscrição:**
```javascript
channel.subscribe('lance', (message) => {
    // Processa o lance recebido
});
```

##  Exemplos de Uso

### Transmitindo uma Partida Completa

```bash
# Terminal 1: Servidor Web
make run-web

# Terminal 2: CLI Jornalista
make run-football

# Publicando lances
> gol | 12 | Pedro abre o placar de cabeça!
> cartao-amarelo | 18 | Falta forte de Gabriel Pec
> substituicao | 45 | Sai Everton Ribeiro, entra Lorran
> gol | 67 | Arrascaeta amplia após contra-ataque!
> impedimento | 72 | Vegetti em posição irregular
> cartao-vermelho | 85 | Expulsão após falta violenta
> gol | 90 | Gabigol fecha a conta nos acréscimos!
```

## Autores

- **Jornalista-Go** - Sistema de publicação

## Links Úteis

- [Documentação Ably](https://ably.com/docs)
- [Ably Go SDK](https://github.com/ably/ably-go)