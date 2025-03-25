package main

import (
	"database/sql"
	"fmt"
)

func insert(data string, table string, column string, db *sql.DB) {
	//insert := fmt.sprintf("INSERT INTO %s  %s\n", column,data)
	result, err := db.Exec(`INSERT INTO $1($2) VALUES ($3)`, table, column, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
