createdb:
	docker exec -it backend-master-class-learn-database-1 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it backend-master-class-learn-database-1 dropdb simple_bank
.PHONY: createdb

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

sqlc:
	sqlc generate