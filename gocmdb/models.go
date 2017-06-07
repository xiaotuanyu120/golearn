package main

import "time"

type server struct {
	UUID         string
	SN           string
	IP           string
	CPU          string
	Memory       string
	Disktype     string
	Disksize     string
	NIC          string
	Manufacturer string
	Model        string
	Expiredate   time.Time
	IDC          string
	Comment      string
}
