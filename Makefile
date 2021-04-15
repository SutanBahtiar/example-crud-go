postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12.6
startdb:
	docker start postgres12
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root example_crud
dropdb:
	docker exec -it postgres12 dropdb example_crud
migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/example_crud?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/example_crud?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: postgres startdb createdb dropdb migrateup migratedown sqlc