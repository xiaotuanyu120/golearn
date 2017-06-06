package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// store sql.DB in env
type Env struct {
	db *sql.DB
}

type Panda struct {
	uuid         string
	sn           string
	ip           string
	cpu          string
	memory       string
	disktype     string
	disksize     string
	nic          string
	manufacturer string
	expiredate   string
	idc          string
	comment      string
}

// APIHandler handle api request
func (env *Env) APIHandler(w http.ResponseWriter, r *http.Request) {
	// set mime to JSON
	w.Header().Set("Content-type", "application/json")

	// err := r.ParseForm()
	// 	if err != nil {
	// 		http.Error(w, fmt.Sprintf("error parsing url %v", err), 500)
	// 	}

	//can't define dynamic slice in golang
	// var result [1000]string

	// switch r.Method {
	// case "GET":
	// 	st, err := env.db.Prepare("select * from assets limit 10")
	// 	if err != nil {
	// 		fmt.Print(err)
	// 	}
	// 	rows, err := st.Query()
	// 	if err != nil {
	// 		fmt.Print(err)
	// 	}
	// 	i := 0
	// 	for rows.Next() {
	// 		var name string
	// 		var id int
	// 		err = rows.Scan(&id, &name)
	// 		panda := &Panda{Id: id, Name: name}
	// 		b, err := json.Marshal(panda)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		result[i] = fmt.Sprintf("%s", string(b))
	// 		i++
	// 	}
	// 	result = result[:i]
	//
	// default:
	// }
}

func main() {
	// connect to database
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/assets")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	env := &Env{db: db}

	// check connection to database
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Mysql Connected!")
	}

	http.HandleFunc("/api", env.APIHandler)
	http.ListenAndServe(":8080", nil)
}
