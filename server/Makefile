postgresinit:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_SSLMODE=require -d postgres:alpine3.19

postgres:
	docker exec -it postgres psql

createdb:
	docker exec -it postgres createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres dropdb go-chat

migrateup:
	migrate --path db/migrations -database "postgres://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate --path db/migrations -database "postgres://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown