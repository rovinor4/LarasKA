package model

type User struct {
	NamaLengkap  string
	NIK          string
	JenisKelamin string
	Alamat       string
	TglLahir     string
	Pekerjaan    string
	NoHP         string
	Email        string
	Password     string
}

var ListUser = []User{
	{
		NamaLengkap:  "Budi Santoso",
		NIK:          "3203123456789012",
		JenisKelamin: "Laki-laki",
		Alamat:       "Jl. Merdeka No.1, Surabaya",
		TglLahir:     "1995-08-17",
		Pekerjaan:    "Mahasiswa",
		NoHP:         "081234567890",
		Email:        "budi@example.com",
		Password:     "password123",
	},
	{
		NamaLengkap:  "Siti Aminah",
		NIK:          "3203123456789013",
		JenisKelamin: "Perempuan",
		Alamat:       "Jl. Pahlawan No.2, Surabaya",
		TglLahir:     "1996-05-05",
		Pekerjaan:    "Karyawan",
		NoHP:         "082345678901",
		Email:        "siti@example.com",
		Password:     "securepass",
	},
}
