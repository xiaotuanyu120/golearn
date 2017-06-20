package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type server struct {
	UUID         string `json:"uuid,omitempty" sql:"uuid"`
	SN           string `json:"sn,omitempty" sql:"sn"`
	IP           string `json:"ip,omitempty" sql:"ip"`
	CPU          string `json:"cpu,omitempty" sql:"cpu"`
	Memory       string `json:"memory,omitempty" sql:"memory"`
	Disktype     string `json:"disktype,omitempty" sql:"disktype"`
	Disksize     string `json:"disksize,omitempty" sql:"disksize"`
	NIC          string `json:"nic,omitempty" sql:"nic"`
	Manufacturer string `json:"manufacturer,omitempty" sql:"manufacturer"`
	Model        string `json:"model,omitempty" sql:"model"`
	Expiredate   string `json:"expiredate,omitempty" sql:"expiredate"`
	IDC          string `json:"idc,omitempty" sql:"idc"`
	Comment      string `json:"comment,omitempty" sql:"comment"`
}

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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/server", serverListHandler(db)).Methods("GET")
	router.HandleFunc("/api/server", serverCreateHandler(db)).Methods("POST")
	router.HandleFunc("/api/server/{uuid}", serverDetailHandler(db)).Methods("GET", "PUT", "DELETE", "POST")
	log.Print("Restful API mux: 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
