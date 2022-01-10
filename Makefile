postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

redis:
	docker run --name redis-instance -p 6378:6379 -d redis

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root product-api

dropdb:
	docker exec -it postgres12 dropdb product-api

migrate-up:
	migrate -path database/migrations -database postgres://root:secret@localhost:5433/product-api?sslmode=disable up

migrate-down:
	migrate -path database/migrations -database postgres://root:secret@localhost:5433/product-api?sslmode=disable down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

run:
	go run cmd/api/main.go cmd/api/app.go