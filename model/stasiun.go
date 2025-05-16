package model

type Stasiun struct {
	IDStasiun string
	Nama      string
	Kota      string
}

var ListStasiun = []Stasiun{
	{IDStasiun: "Mjk", Nama: "Stasiun Mojokerto", Kota: "Mojokerto"},
	{IDStasiun: "Sby", Nama: "Stasiun Gubeng", Kota: "Surabaya"},
	{IDStasiun: "Bli", Nama: "Stasiun Blitar", Kota: "Blitar"},
	{IDStasiun: "Kdr", Nama: "Stasiun Kediri", Kota: "Kediri"},
	{IDStasiun: "Mlg", Nama: "Stasiun Malang", Kota: "Malang"},
	{IDStasiun: "Jbr", Nama: "Stasiun Jember", Kota: "Jember"},
	{IDStasiun: "Pnk", Nama: "Stasiun Pasuruan", Kota: "Pasuruan"},
	{IDStasiun: "Boj", Nama: "Stasiun Bojonegoro", Kota: "Bojonegoro"},
	{IDStasiun: "Ngw", Nama: "Stasiun Nganjuk", Kota: "Nganjuk"},
	{IDStasiun: "Trl", Nama: "Stasiun Tulungagung", Kota: "Tulungagung"},
}