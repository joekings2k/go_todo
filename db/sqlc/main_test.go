package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go_todos/util"
	_ "github.com/lib/pq"
)



var testQueries *Queries


func TestMain (m *testing.M){
	config,err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot get env variables")
	}
	fmt.Print("config files",config.DBDriver,config.DBSource)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}