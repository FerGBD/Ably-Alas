BINARY_NAME=tempo-real-futebol
MAIN_PATH=cmd/football/main.go

.PHONY: all setup run build clean

all: setup run

setup:
	go mod tidy
	go get -u github.com/ably/ably-go/ably
	go get github.com/joho/godotenv

run:
	go run $(MAIN_PATH)

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

clean:
	rm -f $(BINARY_NAME)