package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Drivers:", sql.Drivers())

	db, err := sql.Open("mysql", "root:codeup@tcp(127.0.0.1:3306)/go_chal_db")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	stmt, err := db.Query("DROP TABLE employee")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	fmt.Println("Done")

}
