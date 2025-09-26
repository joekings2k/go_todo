package main

import (
	"database/sql"
	"log"

	"github.com/go_todos/api"
	db "github.com/go_todos/db/sqlc"
	"github.com/go_todos/util"
	_ "github.com/lib/pq"
)



func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load files")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil{
		log.Fatal("cannot connect to db:", err)
		
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server :", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server :", err)
	}
}