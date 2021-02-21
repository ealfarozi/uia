package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "admin:password@(127.0.0.1:3306)/zulu")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	defer db.Close()

	var cnt = 0

	sqlQuery := "select COUNT(*) from institutions"
	err = db.QueryRow(sqlQuery).Scan(&cnt)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(cnt)

}
