package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func tables() {
	fmt.Println("Drivers:", sql.Drivers())

	//CHOOSE DATABASE TO USE
	db, err := sql.Open("mysql", "root:codeup@tcp(127.0.0.1:3306)/go_chal_db")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//DROP TABLES
	dr, err := db.Query("DROP TABLE IF EXISTS product")

	if err != nil {
		fmt.Println("DROP TABLE ERROR:", err.Error())
	}

	defer dr.Close()

	//CREATE TABLES

	crt, err := db.Query("CREATE TABLE IF NOT EXISTS employee(employee_id int primary key auto_increment, name VARCHAR(255), date_updated TIMESTAMP NOT NULL, date_joined TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")

	if err != nil {
		fmt.Println("CREATE TABLE ERROR:", err.Error())
	}

	defer crt.Close()

	//INSERT USER
	// inst, err := db.Query("INSERT INTO employee(first_name) VALUES ('Caleb')")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// defer inst.Close()

	//DELETE USER
	// inst, err := db.Query("DELETE FROM employee where id = '3'")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// defer inst.Close()

	//SELECT FROM A TABLE
	// sel, err := db.Query("SELECT * FROM employee")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// for sel.Next() {
	// 	var (
	// 		id   int
	// 		name string
	// 	)

	// 	err = sel.Scan(&id, &name)
	// 	if err != nil {
	// 		fmt.Println("No users found")
	// 	}

	// 	fmt.Println(id, name)
	// }

	// defer sel.Close()

}
