package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "admin:password@(127.0.0.1:3306)/go_coba")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	tx, err := db.Begin()

	insertQuery := "update pet set name = ? where name = ?"
	updName := "piki"
	extName := "piko"

	_, err2 := tx.Exec(insertQuery, &updName, &extName)
	if err2 != nil {
		tx.Rollback()
		panic(err2.Error())
	}

	tx.Commit()

	var cnt = 0

	sqlQuery := "select COUNT(*) from pet"
	err = db.QueryRow(sqlQuery).Scan(&cnt)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(cnt)

}
