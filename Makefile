postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Hoanglong2502 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root musicafy

dropdb:
	docker exec -it postgres dropdb musicafy

migrateup:
	migrate -path db/migration -database "postgresql://root:localhost1234@localhost:5432/musicafy?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:localhost1234@localhost:5432/musicafy?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:localhost1234@localhost:5432/musicafy?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:localhost1234@localhost:5432/musicafy?sslmode=disable" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build docs/db.dbml

db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

new_service:
	mkdir $(name)
	cd $(name)
	mkdir proto
	mkdir pb
	mkdir db
	mkdir docs
	mkdir gapi
	touch main.go

docker_build:
	docker build -t pharmago_be .

docker_run:
	docker run --name pharmago_be --network pharmago-network -p 8080:8080 -p 9090:9090 -e DB_SOURCE="postgresql://root:Hoanglong2502@pharmago_be-postgres-1:5432/simple_bank?sslmode=disable" -e REDIS_ADDRESS="redis:6379" pharmago_be

# sqlc:
# 	bash sqlc.sh

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 new_migration server sqlc proto evans new_service