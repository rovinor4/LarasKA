package model

import "time"

type Rute struct {
	Kode          string
	Harga         int
	Kapasitas     int
	Gerbong       int
	Keberangkatan time.Time
	Kedatangan    time.Time
	Kereta        Kereta
	StasiunAwal   Stasiun
	StasiunTujuan Stasiun
}

var ListRute = []Rute{
	// Kereta 0, 3 rute
	{
		Kode:          "R001",
		Harga:         100000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 6, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kereta:        ListKereta[0],
		StasiunAwal:   ListStasiun[0],
		StasiunTujuan: ListStasiun[1],
	},
	{
		Kode:          "R002",
		Harga:         110000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kereta:        ListKereta[0],
		StasiunAwal:   ListStasiun[1],
		StasiunTujuan: ListStasiun[2],
	},
	{
		Kode:          "R003",
		Harga:         120000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 12, 0, 0, 0, time.Local),
		Kereta:        ListKereta[0],
		StasiunAwal:   ListStasiun[2],
		StasiunTujuan: ListStasiun[3],
	},
	// Kereta 1, 3 rute
	{
		Kode:          "R004",
		Harga:         100000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 6, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kereta:        ListKereta[1],
		StasiunAwal:   ListStasiun[1],
		StasiunTujuan: ListStasiun[2],
	},
	{
		Kode:          "R005",
		Harga:         110000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kereta:        ListKereta[1],
		StasiunAwal:   ListStasiun[2],
		StasiunTujuan: ListStasiun[3],
	},
	{
		Kode:          "R006",
		Harga:         120000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 12, 0, 0, 0, time.Local),
		Kereta:        ListKereta[1],
		StasiunAwal:   ListStasiun[3],
		StasiunTujuan: ListStasiun[4],
	},
	// Kereta 2, 3 rute
	{
		Kode:          "R007",
		Harga:         100000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 6, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kereta:        ListKereta[2],
		StasiunAwal:   ListStasiun[2],
		StasiunTujuan: ListStasiun[3],
	},
	{
		Kode:          "R008",
		Harga:         110000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kereta:        ListKereta[2],
		StasiunAwal:   ListStasiun[3],
		StasiunTujuan: ListStasiun[4],
	},
	{
		Kode:          "R009",
		Harga:         120000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 12, 0, 0, 0, time.Local),
		Kereta:        ListKereta[2],
		StasiunAwal:   ListStasiun[4],
		StasiunTujuan: ListStasiun[5],
	},
	// Kereta 3, 3 rute
	{
		Kode:          "R010",
		Harga:         100000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 6, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kereta:        ListKereta[3],
		StasiunAwal:   ListStasiun[3],
		StasiunTujuan: ListStasiun[4],
	},
	{
		Kode:          "R011",
		Harga:         110000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kereta:        ListKereta[3],
		StasiunAwal:   ListStasiun[4],
		StasiunTujuan: ListStasiun[5],
	},
	{
		Kode:          "R012",
		Harga:         120000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 12, 0, 0, 0, time.Local),
		Kereta:        ListKereta[3],
		StasiunAwal:   ListStasiun[5],
		StasiunTujuan: ListStasiun[6],
	},
	// Kereta 4, 3 rute
	{
		Kode:          "R013",
		Harga:         100000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 6, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kereta:        ListKereta[4],
		StasiunAwal:   ListStasiun[4],
		StasiunTujuan: ListStasiun[5],
	},
	{
		Kode:          "R014",
		Harga:         110000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 8, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kereta:        ListKereta[4],
		StasiunAwal:   ListStasiun[5],
		StasiunTujuan: ListStasiun[6],
	},
	{
		Kode:          "R015",
		Harga:         120000,
		Kapasitas:     200,
		Gerbong:       5,
		Keberangkatan: time.Date(2025, 5, 20, 10, 0, 0, 0, time.Local),
		Kedatangan:    time.Date(2025, 5, 20, 12, 0, 0, 0, time.Local),
		Kereta:        ListKereta[4],
		StasiunAwal:   ListStasiun[6],
		StasiunTujuan: ListStasiun[7],
	},
}
