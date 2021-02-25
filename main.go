package main

import (
	"database/sql"
	"html/template"
	"log"
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
	selDB, err := db.Query("SELECT * FROM Employee where employee_id=?", nID)
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

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("employeeID")
	selDB, err := db.Query("SELECT * FROM employee WHERE employee_id=?", nID)
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
	temp.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		insForm, err := db.Prepare("INSERT INTO employee(name) VALUES (?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name)
		log.Println("INSERT: Name: " + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		employeeID := r.FormValue("uID")
		insForm, err := db.Prepare("INSERT INTO employee(name) VALUES (?)")
		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, employeeID)
		log.Println("UPDATE: Name: " + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("employeeID")
	delForm, err := db.Prepare("DELETE FROM employee WHERE employee_id=?")
}

func main() {

}
