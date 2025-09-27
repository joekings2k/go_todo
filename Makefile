-include local.app.env
ifneq ("$(wildcard local.app.env)","")
    export $(shell sed 's/=.*//' local.app.env)
endif
postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -d postgres:12-alpine
createdb: 
	docker exec -it postgres12 createdb --username=root --owner=root go_todos
dropdb: 
	docker exec -it postgres12 dropdb go_todos
migrateup:
	migrate -path db/migration -database "postgresql://root:mypassword@localhost:5432/go_todos?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:mypassword@localhost:5432/go_todos?sslmode=disable" -verbose down
sqlc:
	sqlc generate
create migration:
	migrate create -ext sql -dir db/migration -seq init_schema
test:
	go test -v -cover ./...
server:
	go run $(MAIN_GO)

.PHONY: postgres createdb dropdb migrateup migratedown test server
