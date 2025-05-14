package model

type Stasiun struct {
	IDStasiun string
	Nama      string
	Alamat    string
}

type Penumpang struct {
	IDPenumpang string
	NamaLengkap string
	NIK         string
	Kelamin     string
	Alamat      string
	TglLahir    string
	Pekerjaan   string
	NoHP        string
	Email       string
	Pass        string
}

type Rute struct {
	IDRute        string
	StasiunAwal   string
	Tujuan        string
	Harga         int
	Kapasitas     int
	JamPerjalanan string
	IDKereta      string
}

type Kereta struct {
	NamaKereta string
	Kelas      string
}

type Tiket struct {
	Tanggal         string
	Jam             string
	IDPenumpang     string
	Harga           int
	IDRute          string
	NomorTempatDuduk string
}

type Admin struct {
	Username string
	Pass     string
}
