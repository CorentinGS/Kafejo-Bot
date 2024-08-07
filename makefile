APP_NAME = kafejoBot

OUTPUT_DIR = bin

ci:
	./scripts/ci.sh

test:
	go test -v ./...

build:
	go build -o $(OUTPUT_DIR)/$(APP_NAME) -v .

run:
	go run ./...

release:
	GOAMD64=v3 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT_DIR)/$(APP_NAME) -v .
	upx -9 --best $(OUTPUT_DIR)/$(APP_NAME)

clean:
	rm -rf bin/*
