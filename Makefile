postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

removedb:
	docker exec -it postgres12 dropdb simple_bank

start:
	docker start postgres12

stop:
	docker stop postgres12

migrateup:
		migrate -path db/migrations -database "postgresql://root:admin@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:admin@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

.PHONY:	postgres createdb dropdb start migratedown migrateup sqlc
	