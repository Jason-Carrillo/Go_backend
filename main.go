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

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM employee ORDER by id DESC")
	if err != nil {
		panic(err.Error())
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
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)

	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
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
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("employee_id")
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
	tmpl.ExecuteTemplate(w, "Edit", emp)
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
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)
}
