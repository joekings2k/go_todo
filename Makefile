ifneq ("$(wildcard local.app.env)","")
    include local.app.env
    export $(shell sed 's/=.*//' local.app.env)
else ifneq ("$(wildcard app.env)","")
    include app.env
    export $(shell sed 's/=.*//' app.env)
endif

postgres:
	docker run --name postgres12 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p 5432:5432 -d postgres:12-alpine
createdb: 
	docker exec -it postgres12 createdb --username=$(DB_USER) --owner=$(DB_USER) go_todos
dropdb: 
	docker exec -it postgres12 dropdb go_todos
migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down
sqlc:
	sqlc generate
create migration:
	migrate create -ext sql -dir db/migration -seq init_schema
test:
	go test -v -cover ./...
server:
	go run $(MAIN_GO)
mock:
	mockgen -package=mockdb -destination=db/mock/store.go --build_flags=--mod=mod github.com/go_todos/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown test server
