package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	type employee struct {
		id         int
		first_name string
	}

}
