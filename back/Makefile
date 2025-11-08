.PHONY: docker-run run build clear test

run:
	@cd entrypoint
	APP_ENV=dev go run ./entrypoint/main.go

build:
	go build -o bin/ecoply ./entrypoint

clear:
	rm -rf bin/**

docker-run:
	docker compose up --build database -d
	docker compose up --build go
