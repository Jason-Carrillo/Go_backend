package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func tables() {
	fmt.Println("Drivers:", sql.Drivers())

	db, err := sql.Open("mysql", "root:codeup@tcp(127.0.0.1:3306)/go_chal_db")

	if err != nil {
		fmt.Println(err)
		db.Close()
	}

	_, err = db.Exec("go_chal_db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	stmt, err := db.Prepare("CREATE Table employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}

}
