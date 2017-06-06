package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type server struct {
	uuid         string
	sn           string
	ip           string
	cpu          string
	memory       string
	disktype     string
	disksize     string
	nic          string
	manufacturer string
	model        string
	expiredate   time.Time
	idc          string
	comment      string
}

type servers []server

func apiHandler(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error parsing url %v", err), 500)
		}

		switch r.Method {
		case "GET":
			stmt, err := db.Prepare("SELECT * FROM assets.server;")
			if err != nil {
				log.Fatal(err)
			}
			// defer stmt.Close()

			rows, err := stmt.Query()
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			var result servers

			for rows.Next() {
				var uuid string
				var sn string
				var ip string
				var cpu string
				var memory string
				var disktype string
				var disksize string
				var nic string
				var manufacturer string
				var model string
				var expiredate time.Time
				var idc string
				var comment string

				err = rows.Scan(
					&uuid,
					&sn,
					&ip,
					&cpu,
					&memory,
					&disktype,
					&disksize,
					&nic,
					&manufacturer,
					&model,
					&expiredate,
					&idc,
					&comment,
				)
				if err != nil {
					log.Fatal(err)
				}

				ser := server{
					uuid,
					sn,
					ip,
					cpu,
					memory,
					disktype,
					disksize,
					nic,
					manufacturer,
					model,
					expiredate,
					idc,
					comment,
				}
				result = append(result, ser)
			}

		default:
		}

		json, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(todos)
	}

	return fn
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:sudomysql@tcp(127.0.0.1:3306)/assets")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected")
	}

}
