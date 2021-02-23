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

	// crt, err := db.Query("CREATE TABLE employee(id int AUTO INCREMENT, first_name VARCHAR(255))")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// defer crt.Close()

	// inst, err := db.Query("INSERT INTO employee(first_name) VALUES ('Jason')")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// defer inst.Close()

	sel, err := db.Query("SELECT * FROM employee")

	if err != nil {
		fmt.Println(err.Error())
	}

	for sel.Next() {
		var (
			id   int
			name string
		)

		err = sel.Scan(&id, &name)

		fmt.Println(id, name)
	}

	defer sel.Close()

}
