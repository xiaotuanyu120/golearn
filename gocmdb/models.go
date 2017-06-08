package main

import "time"

type server struct {
	UUID         string    `json:"uuid"`
	SN           string    `json:"sn"`
	IP           string    `json:"ip"`
	CPU          string    `json:"cpu"`
	Memory       string    `json:"memory"`
	Disktype     string    `json:"disktype"`
	Disksize     string    `json:"disksize"`
	NIC          string    `json:"nic"`
	Manufacturer string    `json:"manufacturer"`
	Model        string    `json:"model"`
	Expiredate   time.Time `json:"expiredate"`
	IDC          string    `json:"idc"`
	Comment      string    `json:"comment"`
}
