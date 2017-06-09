package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type server struct {
	UUID         string `json:"uuid,omitempty"`
	SN           string `json:"sn,omitempty"`
	IP           string `json:"ip,omitempty"`
	CPU          string `json:"cpu,omitempty"`
	Memory       string `json:"memory,omitempty"`
	Disktype     string `json:"disktype,omitempty"`
	Disksize     string `json:"disksize,omitempty"`
	NIC          string `json:"nic,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
	Expiredate   string `json:"expiredate,omitempty"`
	IDC          string `json:"idc,omitempty"`
	Comment      string `json:"comment,omitempty"`
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
	router.HandleFunc("/api/server/{uuid}", serverDetailHandler(db)).Methods("GET", "PUT", "DELETE")
	log.Print("Restful API mux: 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
