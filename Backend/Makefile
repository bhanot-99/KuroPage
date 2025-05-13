.PHONY: proto build user manga order gateway run-deps test

proto:
	./scripts/generate_proto.sh

build: proto
	go build -o bin/user-service github.com/bhanot-99/MangaVerse/user-service/cmd
	go build -o bin/manga-service github.com/bhanot-99/MangaVerse/manga-service/cmd
	go build -o bin/order-service github.com/bhanot-99/MangaVerse/order-service/cmd
	go build -o bin/api-gateway github.com/bhanot-99/MangaVerse/api-gateway/cmd

user:
	go run user-service/cmd/main.go

manga:
	go run manga-service/cmd/main.go

order:
	go run order-service/cmd/main.go

gateway:
	go run api-gateway/cmd/main.go

run-deps:
	docker-compose up -d

test:
	go test ./...