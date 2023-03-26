package userDB

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Get a database handle.
const dbName = "mysql"
const dbUrl = "root:Chandu@9@tcp(localhost:3306)/users"

func Dbconnection() {
	db, err := sql.Open(dbName, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db, nil
}
