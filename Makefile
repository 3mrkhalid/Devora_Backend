.PHONY: dev prod

dev:
	APP_ENV=development go run ./cmd/api

prod:
	mkdir -p bin
	go build -o bin/devora ./cmd/api
	APP_ENV=production ./bin/devora
