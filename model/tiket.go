package model

import "time"

type Penumpang struct {
	Kode string
	Nama string
	NIK  string
}

type Tiket struct {
	Kode        string
	Rute        Rute
	Price       int
	Penumpang   []Penumpang
	User        User
	CreatedAt   time.Time
	Gerbong     int
	TempatDuduk string
}

var ListTiket []Tiket
