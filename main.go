package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	type employee struct {
		id        int
		firstName string
	}

	db, err := sql.Open("mysql", "root:codeup@tcp(127.0.0.1:3306)/go_chal_db")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	crt, err := db.Query("CREATE TABLE IF NOT EXISTS employee(employee_id int primary key auto_increment, name VARCHAR(255)), date_joined TIMESTAMP NOT NULL, date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP")

	if err != nil {
		fmt.Println("CREATE TABLE ERROR:", err.Error())
	}

	defer crt.Close()

}
