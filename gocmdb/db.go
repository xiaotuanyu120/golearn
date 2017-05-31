package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:sudomysql@tcp(127.0.0.1:3306)/assets")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connected")
	}
}
