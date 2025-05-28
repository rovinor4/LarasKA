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
		Kode:         "F305",
		Nama:         "Jalur F",
		Harga:        12000,
		HargaTetap:   true,
		Gerbong:      6,
		Kereta:       ListKereta[1],
		StasiunAwal:  ListStasiun[31],
		StasiunAkhir: ListStasiun[22],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 20, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[31],
				StasiunAkhir: ListStasiun[28],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 23, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 34, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[28],
				StasiunAkhir: ListStasiun[25],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 9, 36, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 9, 48, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[25],
				StasiunAkhir: ListStasiun[22],
			},
		},
	},
	{
		Kode:         "G8W9",
		Nama:         "Jalur G",
		Harga:        10000,
		HargaTetap:   true,
		Gerbong:      5,
		Kereta:       ListKereta[2],
		StasiunAwal:  ListStasiun[16],
		StasiunAkhir: ListStasiun[3],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 10, 34, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[16],
				StasiunAkhir: ListStasiun[11],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 10, 36, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 10, 49, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[11],
				StasiunAkhir: ListStasiun[9],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 10, 52, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 11, 10, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[9],
				StasiunAkhir: ListStasiun[8],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 11, 13, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 11, 27, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[8],
				StasiunAkhir: ListStasiun[3],
			},
		},
	},
	{
		Kode:         "H2UI",
		Nama:         "Jalur H",
		Harga:        7000,
		HargaTetap:   true,
		Gerbong:      5,
		Kereta:       ListKereta[3],
		StasiunAwal:  ListStasiun[25],
		StasiunAkhir: ListStasiun[8],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 11, 0, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 11, 13, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[25],
				StasiunAkhir: ListStasiun[22],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 11, 17, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 11, 32, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[22],
				StasiunAkhir: ListStasiun[15],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 11, 35, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 11, 46, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[15],
				StasiunAkhir: ListStasiun[12],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 11, 49, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[12],
				StasiunAkhir: ListStasiun[8],
			},
		},
	},
	{
		Kode:         "IP99",
		Nama:         "Jalur I",
		Harga:        5000,
		HargaTetap:   true,
		Gerbong:      5,
		Kereta:       ListKereta[8],
		StasiunAwal:  ListStasiun[27],
		StasiunAkhir: ListStasiun[15],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 11, 53, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 12, 5, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[27],
				StasiunAkhir: ListStasiun[23],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 12, 7, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 12, 20, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[23],
				StasiunAkhir: ListStasiun[20],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 12, 22, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 12, 32, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[20],
				StasiunAkhir: ListStasiun[15],
			},
		},
	},
	{
		Kode:         "J83I",
		Nama:         "Jalur J",
		Harga:        15000,
		HargaTetap:   true,
		Gerbong:      8,
		Kereta:       ListKereta[7],
		StasiunAwal:  ListStasiun[33],
		StasiunAkhir: ListStasiun[1],
		RuteBerhenti: []RuteBerhenti{
			{
				Berangkat:    time.Date(2023, 10, 1, 12, 43, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 12, 58, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[33],
				StasiunAkhir: ListStasiun[29],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 13, 1, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 13, 25, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[29],
				StasiunAkhir: ListStasiun[21],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 13, 27, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 13, 36, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[21],
				StasiunAkhir: ListStasiun[13],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 13, 39, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 13, 50, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[13],
				StasiunAkhir: ListStasiun[8],
			},
			{
				Berangkat:    time.Date(2023, 10, 1, 13, 52, 0, 0, time.UTC),
				Tiba:         time.Date(2023, 10, 1, 14, 10, 0, 0, time.UTC),
				StasiunAwal:  ListStasiun[13],
				StasiunAkhir: ListStasiun[1],
			},
		},
	},
}