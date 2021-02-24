package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	employee_id int
	name        string
	dateCreated string
	dateUpdated string
}

func dbConn() (db *sql.db) {
	dbDriver := "mysql"
	dbUser := "root"
	dbpass := "codeup"
	dbName := "go_chal_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbpass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var temp = template.Must(template.ParseGlob("form/*"))

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM employee ORDER by id DESC")
	if err != nil {
		panic("Index Error: " + err.Error())
	}
	emp := employee{}
	res := []employee{}
	for selDB.Next() {
		var employee_id int
		var name string
		var dateCreated string
		var dateUpdated string

		err = selDB.Scan(&employee_id, &name, &dateCreated, &dateUpdated)
		if err != nil {
			panic(err.Error())
		}
		emp.employee_id = employee_id
		emp.name = name
		emp.dateCreated = dateCreated
		emp.dateUpdated = dateUpdated
	}
}

func main() {

}
