.PHONY: docker-dev docker-dev-build docker-prod docker-prod-build run build clear test

run:
	@cd entrypoint
	APP_ENV=dev go run ./entrypoint/main.go

build:
	go build -o bin/ecoply ./entrypoint

clear:
	rm -rf bin/**

docker-dev:
	docker compose -f compose.dev.yaml up database -d
	docker compose -f compose.dev.yaml up go

docker-dev-build:
	docker compose -f compose.dev.yaml up --build database -d
	docker compose -f compose.dev.yaml up --build go

docker-prod:
	docker compose -f compose.prod.yaml up database -d
	docker compose -f compose.prod.yaml up go

docker-prod-build:
	docker compose -f compose.prod.yaml up --build database -d
	docker compose -f compose.prod.yaml up --build go
