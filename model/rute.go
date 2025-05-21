package model

import "time"

type Rute struct {
	Kode         string
	Nama         string
	HargaTetap   bool
	Harga        int
	Gerbong      int
	Kereta       Kereta
	StasiunAwal  Stasiun
	StasiunAkhir Stasiun
	RuteBerhenti []RuteBerhenti
}

type RuteBerhenti struct {
	Berangkat    time.Time
	Tiba         time.Time
	StasiunAwal  Stasiun
	StasiunAkhir Stasiun
}

var RuteList []Rute = []Rute{
	{
		Kode:         "J305",
		Nama:         "Jalur J",
		Harga:        10000,
		HargaTetap:   true,
		Gerbong:      5,
		Kereta:       ListKereta[1],
		StasiunAwal:  ListStasiun[8],
		StasiunAkhir: ListStasiun[1],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 8, 30, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[8],
				StasiunAkhir: ListStasiun[3],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 8, 30, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[3],
				StasiunAkhir: ListStasiun[2],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 30, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[2],
				StasiunAkhir: ListStasiun[1],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 30, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[1],
				StasiunAkhir: ListStasiun[0],
			},
		},
	},
}
