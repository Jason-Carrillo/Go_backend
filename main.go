package main

import (
	"database/sql"
	"fmt"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	type employee struct {
		id        int
		firstName string
	}

	func dbConn() (db *sql.db) {
		dbDriver := "mysql"
		dbUser := "root"
		dbpass := "codeup"
		dbName := "go_chal_db"
		db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbname )
		if err != nil {
			panic(err.Error())
		}
		return db
	}


	var temp = template.Must(template.ParseGlob("form/*"))

	func index(w http.ResponseWriter, r *http.Request) {
		db := dbConn()
	}

}
