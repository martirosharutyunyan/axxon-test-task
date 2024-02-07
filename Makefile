include .env
default: run

build:
	go build -C ./cmd/main-service -o main-service
	du -sh .

start-prod:
	cd cmd/main-service && ./main-service

run:
	cd cmd/main-service && go run .

dev:
	make swag
	cd cmd/main-service && go run .


install:
	cd cmd/main-service && go install

update:
	go get -u ./...

clean:
	cd cmd/main-service && go clean

swag:
	cd cmd/main-service && swag init -g ./modules/controllers/server.go -o ../../docs -d ../../pkg --pd

swag-fmt:
	swag fmt

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

lint-clean:
	golangci-lint cache clean

tidy:
	go mod tidy

install-devtools:
	sh install-dev-dependencies.sh

pre-commit:
	pre-commit run --all-files

up-migrations:
	cd ./pkg/migrations && goose postgres "user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable host=${DB_HOST} port=${DB_PORT}"  up

down-migration:
	cd ./pkg/migrations && goose postgres "user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable host=${DB_HOST} port=${DB_PORT}"  down


reset-migrations:
	cd ./pkg/migrations && goose postgres "user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable host=${DB_HOST} port=${DB_PORT}" reset

generate-migration:
	cd ./pkg/migrations && goose postgres "user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable host=${DB_HOST} port=${DB_PORT}" create ${name} sql

fieldalignment:
	fieldalignment -fix ./...

sqlc-generate:
	sqlc generate

docker-build:
	docker buildx build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} \
                  --build-arg GITHUB_USERNAME=${GITHUB_USERNAME} \
                  --build-arg GOPRIVATE=github.com/go-boilerplate/* \
                  -t backend-workspace-main-service:latest .