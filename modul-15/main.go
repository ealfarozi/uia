package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	type Institution struct {
		ID   int64  `json:"id,omitempty"`
		Code string `json:"code" validate:"required"`
		Name string `json:"name" validate:"required"`
	}

	var ins []Institution
	var insSing Institution

	// Open up our database connection.
	db, err := sql.Open("mysql", "admin:password@(127.0.0.1:3306)/zulu")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	defer db.Close()

	sqlQuery := "select id, code, name from institutions"
	res, err2 := db.Query(sqlQuery)

	for res.Next() {
		res.Scan(&insSing.ID, &insSing.Code, &insSing.Name)
		ins = append(ins, insSing)
	}

	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Println("Struct Print: ", &ins)

	b, err := json.Marshal(ins)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json Print: ", string(b))

}
