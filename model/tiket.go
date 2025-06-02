package model

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
	StasiunAwal  Stasiun
	StasiunAkhir Stasiun
}

var ListTiket []Tiket = []Tiket{
	{
		Kode:  "TK-X8F3L",
		Rute:  ListRute[0],
		Price: ListRute[0].Harga,
		User:  User{},
		Penumpang: []Penumpang{
			{
				Kode:        "PN-G7R4M",
				Nama:        "Penumpang 1",
				NIK:         "3510000001",
				Gerbong:     ListRute[0].Gerbong,
				TempatDuduk: "A1",
			},
			{
				Kode:        "PN-G824M",
				Nama:        "Penumpang 2",
				NIK:         "3510000001",
				Gerbong:     ListRute[0].Gerbong,
				TempatDuduk: "A2",
			},
		},
		StasiunAwal:  ListRute[0].RuteBerhenti[0].StasiunAwal,
		StasiunAkhir: ListRute[0].RuteBerhenti[0].StasiunAkhir,
	},
	{
		Kode:  "TK-K2Z9B",
		Rute:  ListRute[1],
		Price: ListRute[1].Harga,
		User:  User{},
		Penumpang: []Penumpang{
			{
				Kode:        "PN-V4D8Q",
				Nama:        "Penumpang 2",
				NIK:         "3510000002",
				Gerbong:     ListRute[1].Gerbong,
				TempatDuduk: "A2",
			},
		},
		StasiunAwal:  ListRute[1].RuteBerhenti[0].StasiunAwal,
		StasiunAkhir: ListRute[1].RuteBerhenti[0].StasiunAkhir,
	},
}
