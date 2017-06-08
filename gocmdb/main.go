package main

import (
	"database/sql"
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
		log.Print("database connected")
	}

	mux := http.NewServeMux()
	mux.Handle("/api/server", serverlistHandler(db))
	mux.Handle("/api/server/", serverdetailHandler(db))
	log.Print("Restful API mux")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
