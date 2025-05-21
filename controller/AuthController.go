package controller

import (
	"bufio"
	"fmt"
	"laraska/model"
	"os"
	"regexp"
	"strings"
)

type AuthStruct struct {
	is_admin bool
	admin    model.Admin
	User     model.User
}

var AuthData AuthStruct

func AuthController() {
Step1:
	PrintBoxWithText(60, []string{
		"Selamat Datang",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

	var pilihan string
	fmt.Println("[1] Login Sebagai Pengguna")
	fmt.Println("[2] Daftar Pengguna Baru")
	fmt.Println("[3] Login Sebagai Admin")

	fmt.Print("Pilih menu : ")
	_, err := fmt.Scan(&pilihan)

	if err != nil || !IsNumeric(pilihan) {
		ClearScreen()
		PrintError(fmt.Sprintf("Pilihan %s tidak ada", pilihan))
		goto Step1
	}

	switch pilihan {
	case "2":
		RegisterPenumpang()
	case "3":
		ClearScreen()
		LoginForAdmin()
	default:
		ClearScreen()
		PrintError(fmt.Sprintf("Pilihan %s tidak ada", pilihan))
		goto Step1
	}
}

func LoginForAdmin() {
	var username, password string

	PrintBoxWithText(60, []string{
		"Login Sebagai Admin",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

Form:
	fmt.Print("Masukan username : ")
	fmt.Scan(&username)

	fmt.Print("Masukan password : ")
	fmt.Scan(&password)

	isNull := true
	for _, admin := range model.ListAdmin {
		if admin.Username == username && admin.Pass == password {
			AuthData = AuthStruct{
				is_admin: true,
				admin:    admin,
			}
			isNull = false
			ClearScreen()
			MenuAwalAdmin()
			return
		}
	}

	if isNull == true {
		PrintError("Username dan password salah")
		Divider("-")
		goto Form
	}
}

func RegisterPenumpang() {
	reader := bufio.NewReader(os.Stdin)
	ClearScreen()
	PrintBoxWithText(60, []string{
		"Daftar Pengguna",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

FormNama:
	fmt.Print("Masukan nama lengkap : ")
	nama, err := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama == "" || err != nil {
		PrintError("Nama wajib di isi")
		goto FormNama
	}
FormNIK:
	fmt.Print("Masukan NIK : ")
	nik, err := reader.ReadString('\n')
	nik = strings.TrimSpace(nik)
	if nik == "" || err != nil {
		PrintError("NIK wajib di isi")
		goto FormNIK
	}

	var nikRegex = regexp.MustCompile(`^\d{16}$`)
	if !nikRegex.MatchString(nik) {
		PrintError("Format NIK salah")
		goto FormNIK
	}

	dataNIK := BinaryFindMany(model.ListUser, model.User{NIK: nik}, func(a, b model.User) int {
		if a.NIK < b.NIK {
			return -1
		} else if a.NIK > b.NIK {
			return 1
		}
		return 0
	})

	if len(dataNIK) > 0 {
		PrintError("NIK sudah digunakan")
		goto FormNIK
	}
FormKelamin:
	fmt.Print("Masukan Jenis Kelamin (l/p) : ")
	jenisKelamin, err := reader.ReadString('\n')
	jenisKelamin = strings.TrimSpace(jenisKelamin)
	if jenisKelamin == "" || err != nil {
		PrintError("Wajib dipilih jenis kelamin")
		goto FormKelamin
	}

	if jenisKelamin != "l" && jenisKelamin != "p" {
		PrintError("Pilihan hanya l (Laki-Laki) dan p (Perempuan)")
		goto FormKelamin
	}

FormAlamat:
	fmt.Print("Masukan Alamat : ")
	alamat, err := reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)
	if alamat == "" || err != nil {
		PrintError("Alamat Wajib di isi")
		goto FormAlamat
	}

FormTglLahir:
	fmt.Print("Masukan Tanggal Lahir  (DD/MM/YYYY) : ")
	TglLahir, err := reader.ReadString('\n')
	TglLahir = strings.TrimSpace(TglLahir)
	if TglLahir == "" || err != nil {
		PrintError("Alamat Wajib di isi")
		goto FormTglLahir
	}

	if !regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/[0-9]{4}$`).MatchString(TglLahir) {
		PrintError("Format tanggal lahir salah")
		goto FormTglLahir
	}

FormPekerjaan:
	fmt.Print("Masukan Pekerjaan : ")
	pekerjaan, err := reader.ReadString('\n')
	pekerjaan = strings.TrimSpace(pekerjaan)
	if pekerjaan == "" || err != nil {
		PrintError("Pekerjaan wajib diisi")
		goto FormPekerjaan
	}

FormNoHp:
	fmt.Print("Masukan No HP : ")
	noHp, err := reader.ReadString('\n')
	noHp = strings.TrimSpace(noHp)
	if !regexp.MustCompile(`^08[0-9]{8,11}$`).MatchString(noHp) {
		PrintError("Format No HP salah")
		goto FormNoHp
	}
	dataNoHp := BinaryFindMany(model.ListUser, model.User{NoHP: noHp}, func(a, b model.User) int {
		if a.NoHP < b.NoHP {
			return -1
		} else if a.NoHP > b.NoHP {
			return 1
		}
		return 0
	})
	if len(dataNoHp) > 0 {
		PrintError("No HP sudah digunakan")
		goto FormNoHp
	}

FormEmail:
	fmt.Print("Masukan Email : ")
	email, err := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if !regexp.MustCompile(`^[\w.+-]+@[\w-]+\.[\w.-]+$`).MatchString(email) {
		PrintError("Format Email salah")
		goto FormEmail
	}
	dataEmail := BinaryFindMany(model.ListUser, model.User{Email: email}, func(a, b model.User) int {
		if a.Email < b.Email {
			return -1
		} else if a.Email > b.Email {
			return 1
		}
		return 0
	})
	if len(dataEmail) > 0 {
		PrintError("Email sudah digunakan")
		goto FormEmail
	}

FormPassword:
	fmt.Print("Masukan Password : ")
	password, err := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	if len(password) < 6 {
		PrintError("Password minimal 6 karakter")
		goto FormPassword
	}

	User := model.User{
		NamaLengkap:  nama,
		NIK:          nik,
		JenisKelamin: jenisKelamin,
		Alamat:       alamat,
		TglLahir:     TglLahir,
		Pekerjaan:    pekerjaan,
		NoHP:         noHp,
		Email:        email,
		Password:     password,
	}
	model.ListUser = append(model.ListUser, User)
	ClearScreen()
	fmt.Println(ColorText("User Berhasil Terdaftar", 30, 42, false))
	AuthController()
}
