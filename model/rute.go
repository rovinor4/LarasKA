package model

import "time"

type Rute struct {
	IDRute        string
	Harga         int
	Kapasitas     int
	IDKereta      string
	Keberangkatan time.Time
	Kedatangan    time.Time
	StasiunAwal   Stasiun
	StasiunTujuan Stasiun
}
