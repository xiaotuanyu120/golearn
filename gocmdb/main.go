package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/assets?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB connected")
	}

	http.HandleFunc("/api/server", serverlistHandler(db))
	http.HandleFunc("/api/server/{uuid}", serverdetailHandler(db))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
