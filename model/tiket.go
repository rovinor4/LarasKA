package model

import "time"

type Penumpang struct {
	Kode        string
	Nama        string
	NIK         string
	Gerbong     int
	TempatDuduk string
}

type Tiket struct {
	Kode         string
	Rute         Rute
	Price        int
	User         User
	Penumpang    []Penumpang
	CreatedAt    time.Time
	StasiunAwal  Stasiun
	StasiunAkhir Stasiun
}

var ListTiket []Tiket
