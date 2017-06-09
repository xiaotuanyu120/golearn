package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func serverListHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error parsing url %v", err), 500)
		}

		var result []string

		var stmt *sql.Stmt
		stmt, err = db.Prepare("SELECT * FROM server")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		var rows *sql.Rows
		rows, err = stmt.Query()
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
			var expiredate string
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

		// write result(JSON) to ResponseWriter
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func serverCreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		decoder := json.NewDecoder(r.Body)
		var ser server
		err := decoder.Decode(&ser)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		uuid := ser.UUID
		sn := ser.SN
		ip := ser.IP
		cpu := ser.CPU
		memory := ser.Memory
		disktype := ser.Disktype
		disksize := ser.Disksize
		nic := ser.NIC
		manufacturer := ser.Manufacturer
		model := ser.Model
		expiredate := ser.Expiredate
		idc := ser.IDC
		comment := ser.Comment

		var stmt *sql.Stmt
		stmt, err = db.Prepare("INSERT INTO server VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(uuid, sn, ip, cpu, memory, disktype, disksize, nic, manufacturer, model, expiredate, idc, comment)
		if err != nil {
			log.Fatal(err)
		}

		// write result(JSON) to ResponseWriter
		err = json.NewEncoder(w).Encode(ser)
		if err != nil {
			log.Println(ser)
			return
		}
	}
}

func serverDetailHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		var result []byte

		switch r.Method {
		case "GET":
			UUID := mux.Vars(r)["uuid"]

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
			var expiredate string
			var idc string
			var comment string

			stmt, err := db.Prepare("SELECT * FROM server WHERE uuid = ?")
			if err != nil {
				log.Fatal(err)
			}

			err = stmt.QueryRow(UUID).Scan(
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
			switch {
			case err == sql.ErrNoRows:
				// result = fmt.Sprintf("NO SUCH UUID: %v.", UUID)
				result = []byte{}
				// log.Printf(result)
			case err != nil:
				log.Fatal(err)
			default:
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

				result, err = json.Marshal(ser)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

		case "PUT":
			UUID := mux.Vars(r)["uuid"]

			decoder := json.NewDecoder(r.Body)
			var ser server
			err := decoder.Decode(&ser)
			if err != nil {
				panic(err)
			}
			defer r.Body.Close()

			//test
			//test end

			sn := ser.SN
			ip := ser.IP
			cpu := ser.CPU
			memory := ser.Memory
			disktype := ser.Disktype
			disksize := ser.Disksize
			nic := ser.NIC
			manufacturer := ser.Manufacturer
			model := ser.Model
			expiredate := ser.Expiredate
			idc := ser.IDC
			comment := ser.Comment

			var stmt *sql.Stmt
			stmt, err = db.Prepare("UPDATE server SET sn=?, ip=?, cpu=?, memory=?, disktype=?, disksize=?, nic=?, manufacturer=?, model=?, expiredate=?, idc=?, comment=? WHERE uuid=?")
			if err != nil {
				log.Fatal(err)
			}

			res, err := stmt.Exec(sn, ip, cpu, memory, disktype, disksize, nic, manufacturer, model, expiredate, idc, comment, UUID)
			if err != nil {
				log.Fatal(err)
			}

			if res != nil {
				result, err = json.Marshal(ser)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			// case "DELETE":

			// default:
			// 	result = fmt.Sprintf("Don't support Method %v.", r.Method)
		}

		// write result(JSON) to ResponseWriter
		w.Write(result)
		return
	}
}
