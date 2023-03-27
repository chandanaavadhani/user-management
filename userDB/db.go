package userDB

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Get a database handle.
const dbName = "mysql"
const dbUrl = "root:Chandu@9@tcp(localhost:3306)/users"

func Dbconnection() (*sql.DB, error) {
	db, err := sql.Open(dbName, dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	// defer db.Close()
	return db, nil
}
