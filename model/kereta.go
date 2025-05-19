package model

type Kereta struct {
	Kode  int
	Nama  string
	Kelas string
}

var ListKereta = []Kereta{
	{Kode: 873, Nama: "Commuterline Jenggala", Kelas: "Ekonomi"},
	{Kode: 219, Nama: "Argo Bromo Anggrek", Kelas: "Eksekutif"},
	{Kode: 614, Nama: "Gajayana", Kelas: "Eksekutif"},
	{Kode: 475, Nama: "Taksaka", Kelas: "Eksekutif"},
	{Kode: 342, Nama: "Lodaya", Kelas: "Bisnis"},
	{Kode: 981, Nama: "Mutiara Selatan", Kelas: "Bisnis"},
	{Kode: 130, Nama: "Majapahit", Kelas: "Ekonomi"},
	{Kode: 728, Nama: "Jayabaya", Kelas: "Ekonomi"},
	{Kode: 599, Nama: "Kertajaya", Kelas: "Ekonomi"},
	{Kode: 407, Nama: "Pasundan", Kelas: "Ekonomi"},
	{Kode: 362, Nama: "Commuterline Dhoho", Kelas: "Ekonomi"},
	{Kode: 290, Nama: "Commuterline Penataran", Kelas: "Ekonomi"},
}
