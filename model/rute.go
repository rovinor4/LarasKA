package model

import "time"

type Rute struct {
	Harga         int
	Kapasitas     int
	Keberangkatan time.Time
	Kedatangan    time.Time
	Kereta        Kereta
	StasiunAwal   Stasiun
	StasiunTujuan Stasiun
}
