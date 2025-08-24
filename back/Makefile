.PHONY: docker-run docker-run-dev run build clear test

run:
	@cd entrypoint
	go run ./entrypoint/main.go

build:
	go build -o bin/ecoply ./entrypoint

clear:
	rm -rf bin/**

docker-run:
	docker compose up --build -d

docker-run-dev:
	docker compose -f compose.yaml -f compose.dev.yaml up --build -d


