.PHONY: docker-run docker-run-dev run run-debug build clear test

run:
	@cd entrypoint
	APP_ENV=dev go run ./entrypoint/main.go

run-debug:
	APP_ENV=dev dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient ./entrypoint/main.go

build:
	go build -o bin/ecoply ./entrypoint

clear:
	rm -rf bin/**

docker-run:
	docker compose up --build database nginx -d
	docker compose up --build go

docker-run-prod:
	docker compose -f compose.yaml -f compose.prod.yaml up --build -d


