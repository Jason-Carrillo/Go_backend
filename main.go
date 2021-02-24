package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	employeeID  int
	name        string
	dateCreated string
	dateUpdated string
}

func dbConn() (db *sql.DB) {
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
		var employeeID int
		var name string
		var dateCreated string
		var dateUpdated string

		err = selDB.Scan(&employeeID, &name, &dateCreated, &dateUpdated)
		if err != nil {
			panic(err.Error())
		}
		emp.employeeID = employeeID
		emp.name = name
		emp.dateCreated = dateCreated
		emp.dateUpdated = dateUpdated
	}
	temp.ExecuteTemplate(w, "Index", res)

	defer db.Close()
}

func show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("employeeID")
	selDB, err := db.Query("SELECT * FROM Employee where id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	emp := employee{}
	for selDB.Next() {
		var employeeID int
		var name string
		var dateCreated string
		var dateUpdated string
		err = selDB.Scan(&employeeID, &name, &dateCreated, &dateUpdated)
		if err != nil {
			panic(err.Error())
		}
		emp.employeeID = employeeID
		emp.name = name
		emp.dateCreated = dateCreated
		emp.dateUpdated = dateUpdated
	}
	temp.ExecuteTemplate(w, "show", emp)
	defer db.Close()
}

func new(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("employeeID")
	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nID)
}

func main() {

}
