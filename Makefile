include .env

.PHONY: build
build:
	go build -o bin/$(APP_NAME).exe cmd/app/main.go

.PHONY: air
air: 
	air -c .air.toml

.PHONY: migrate-up
migrate-up: 
	go run cmd/app/main.go migrate -action up

.PHONY: migrate-down
migrate-down: 
	go run cmd/app/main.go migrate -action down

.PHONY: run
run: 
	go run cmd/app/main.go

.PHONY: seed
seed: 
	go run cmd/app/main.go seed

.PHONY: compose-up
compose-up:
	docker-compose up --detach --build

.PHONY: compose-down
compose-down:
	docker-compose down