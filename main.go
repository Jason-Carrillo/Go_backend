package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int    `db:"age"`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "codeup"
	dbName := "go_chal_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic("dbConn:" + err.Error())
	}
	fmt.Println("TEST")
	return db
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM person")
	ErrorCheck(err)
	per := Person{}
	var res []Person
	for selDB.Next() {
		var ID int
		var FirstName, LastName string
		var Age int
		err = selDB.Scan(&ID, &FirstName, &LastName, &Age)
		ErrorCheck(err)
		per.ID = ID
		per.FirstName = FirstName
		per.LastName = LastName
		per.Age = Age
		res = append(res, per)
		fmt.Println(res)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

// func Show(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	nId := r.URL.Query().Get("id")
// 	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	emp := Employee{}
// 	for selDB.Next() {
// 		var id int
// 		var name string
// 		err = selDB.Scan(&id, &name)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		emp.Id = id
// 		emp.Name = name
// 	}
// 	tmpl.ExecuteTemplate(w, "Show", emp)
// 	defer db.Close()
// }

// func New(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "New", nil)
// }

// func Edit(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	nId := r.URL.Query().Get("id")
// 	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	emp := Employee{}
// 	for selDB.Next() {
// 		var id int
// 		var name string
// 		err = selDB.Scan(&id, &name)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		emp.Id = id
// 		emp.Name = name
// 	}
// 	tmpl.ExecuteTemplate(w, "Edit", emp)
// 	defer db.Close()
// }

// func Insert(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		insForm, err := db.Prepare("INSERT INTO Employee(name) VALUES(?,?)")
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		insForm.Exec(name)
// 		log.Println("INSERT: Name: " + name)
// 	}
// 	defer db.Close()
// 	http.Redirect(w, r, "/", 301)
// }

// func Update(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		id := r.FormValue("uid")
// 		insForm, err := db.Prepare("UPDATE Employee SET name=? WHERE id=?")
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		insForm.Exec(name, id)
// 		log.Println("UPDATE: Name: " + name)
// 	}
// 	defer db.Close()
// 	http.Redirect(w, r, "/", 301)
// }

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	emp := r.URL.Query().Get("id")
// 	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	delForm.Exec(emp)
// 	log.Println("DELETE")
// 	defer db.Close()
// 	http.Redirect(w, r, "/", 301)
// }

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	// http.HandleFunc("/show", Show)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/update", Update)
	// http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
