package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func serverlistHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		// err := r.ParseForm()
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("error parsing url %v", err), 500)
		// }

		var result []string

		switch r.Method {
		case "GET":
			// var stmt *sql.Stmt
			stmt, err := db.Prepare("SELECT * FROM server")
			if err != nil {
				log.Fatal(err)
			}
			// defer stmt.Close()

			// var rows *sql.Rows
			rows, err := stmt.Query()
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

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

				// var serJSON []byte
				serJSON, err := json.Marshal(ser)
				if err != nil {
					fmt.Println(err)
					return
				}
				result = append(result, string(serJSON))
			}

		case "POST":
			uuid := r.PostFormValue("uuid")
			sn := r.PostFormValue("sn")
			ip := r.PostFormValue("ip")
			cpu := r.PostFormValue("cpu")
			memory := r.PostFormValue("memory")
			disktype := r.PostFormValue("disktype")
			disksize := r.PostFormValue("disksize")
			nic := r.PostFormValue("nic")
			manufacturer := r.PostFormValue("manufacturer")
			model := r.PostFormValue("model")
			expiredate := r.PostFormValue("expiredate")
			idc := r.PostFormValue("idc")
			comment := r.PostFormValue("comment")

			stmt, err := db.Prepare("INSERT INTO server VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)")
			if err != nil {
				log.Fatal(err)
			}

			res, err := stmt.Exec(uuid, sn, ip, cpu, memory, disktype, disksize, nic, manufacturer, model, expiredate, idc, comment)
			if err != nil {
				log.Fatal(err)
			}

			if res != nil {
				result = append(result, "success!")
			}

		default:
			result = []string{fmt.Sprintf("Don't support Method %v, only GET/POST.", r.Method)}
		}

		// write result(JSON) to ResponseWriter
		err := json.NewEncoder(w).Encode(result)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func serverdetailHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error parsing url %v", err), 500)
		}

		var result []string

		switch r.Method {
		case "GET":
			fmt.Println("start get")
			uuid := r.FormValue("uuid")
			fmt.Println(uuid)

			var stmt *sql.Stmt
			stmt, err = db.Prepare("SELECT * FROM server WHERE uuid = ?")
			if err != nil {
				log.Fatal(err)
			}

			var rows *sql.Rows
			rows, err = stmt.Query(uuid)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

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

				var serJSON []byte
				serJSON, err = json.Marshal(ser)
				if err != nil {
					fmt.Println(err)
					return
				}
				result = append(result, string(serJSON))
			}

		case "PUT":
			fmt.Println("put start")
			uuid := r.PostFormValue("uuid")
			sn := r.PostFormValue("sn")
			ip := r.PostFormValue("ip")
			cpu := r.PostFormValue("cpu")
			memory := r.PostFormValue("memory")
			disktype := r.PostFormValue("disktype")
			disksize := r.PostFormValue("disksize")
			nic := r.PostFormValue("nic")
			manufacturer := r.PostFormValue("manufacturer")
			model := r.PostFormValue("model")
			expiredate := r.PostFormValue("expiredate")
			idc := r.PostFormValue("idc")
			comment := r.PostFormValue("comment")
			fmt.Println(sn, ip, cpu, memory, disktype, disksize, nic, manufacturer, model, expiredate, idc, comment)

			var stmt *sql.Stmt
			stmt, err = db.Prepare("UPDATE server SET sn=?, ip=?, cpu=?, memory=?, disktype=?, disksize=?, nic=?, manufacturer=?, model=?, expiredate=?, idc=?, comment=? WHERE id=?")
			if err != nil {
				log.Fatal(err)
			}

			var res sql.Result
			res, err = stmt.Exec(sn, ip, cpu, memory, disktype, disksize, nic, manufacturer, model, expiredate, idc, comment, uuid)
			if err != nil {
				log.Fatal(err)
			}

			if res != nil {
				result = append(result, "success!")
			}

		// case "DELETE":

		default:
			result = []string{fmt.Sprintf("Don't support Method %v.", r.Method)}
		}

		// write result(JSON) to ResponseWriter
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
