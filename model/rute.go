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
		Kode:         "A305",
		Nama:         "Jalur A",
		Harga:        10000,
		HargaTetap:   true,
		Gerbong:      5,
		Kereta:       ListKereta[1],
		StasiunAwal:  ListStasiun[8],
		StasiunAkhir: ListStasiun[0],
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
	{
		Kode:         "B3J9",
		Nama:         "Jalur B",
		Harga:        5000,
		HargaTetap:   true,
		Gerbong:      4,
		Kereta:       ListKereta[6],
		StasiunAwal:  ListStasiun[8],
		StasiunAkhir: ListStasiun[19],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 20, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[8],
				StasiunAkhir: ListStasiun[10],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 22, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 40, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[10],
				StasiunAkhir: ListStasiun[13],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 42, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 50, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[13],
				StasiunAkhir: ListStasiun[16],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 52, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 10, 5, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[16],
				StasiunAkhir: ListStasiun[19],
			},
		},
	},
	{
		Kode:         "C4L0",
		Nama:         "Jalur C",
		Harga:        12500,
		HargaTetap:   true,
		Gerbong:      6,
		Kereta:       ListKereta[0],
		StasiunAwal:  ListStasiun[14],
		StasiunAkhir: ListStasiun[31],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 30, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[14],
				StasiunAkhir: ListStasiun[24],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 10, 10, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 10, 23, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[24],
				StasiunAkhir: ListStasiun[26],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 10, 30, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 11, 9, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[26],
				StasiunAkhir: ListStasiun[31],
			},
		},
	},
	{
		Kode:         "D5R7",
		Nama:         "Jalur D",
		Harga:        7500,
		HargaTetap:   true,
		Gerbong:      4,
		Kereta:       ListKereta[4],
		StasiunAwal:  ListStasiun[24],
		StasiunAkhir: ListStasiun[33],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 10, 15, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 10, 25, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[24],
				StasiunAkhir: ListStasiun[26],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 10, 37, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 10, 55, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[26],
				StasiunAkhir: ListStasiun[29],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 10, 57, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 11, 20, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[29],
				StasiunAkhir: ListStasiun[33],
			},
		},
	},
	{
		Kode:         "E84T",
		Nama:         "Jalur E",
		Harga:        21000,
		HargaTetap:   true,
		Gerbong:      5,
		Kereta:       ListKereta[7],
		StasiunAwal:  ListStasiun[33],
		StasiunAkhir: ListStasiun[6],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 7, 15, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 7, 42, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[33],
				StasiunAkhir: ListStasiun[31],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 7, 50, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 8, 33, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[31],
				StasiunAkhir: ListStasiun[28],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 8, 35, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 8, 51, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[28],
				StasiunAkhir: ListStasiun[18],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 1, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 14, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[18],
				StasiunAkhir: ListStasiun[6],
			},
		},
	},
	{
		Kode:         "J305",
		Nama:         "Jalur J",
		Harga:        10000,
		HargaTetap:   true,
		Gerbong:      5,
		Kereta:       ListKereta[1],
		StasiunAwal:  ListStasiun[8],
		StasiunAkhir: ListStasiun[0],
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