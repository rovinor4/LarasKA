package controller

import (
	"bufio"
	"fmt"
	"laraska/model"
	"laraska/utils"
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

	reader := bufio.NewReader(os.Stdin)
	utils.PrintBoxWithText(60, []string{
		"Selamat Datang",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

	fmt.Println("[1] Login Sebagai Pengguna")
	fmt.Println("[2] Daftar Pengguna Baru")
	fmt.Println("[3] Login Sebagai Admin")

	utils.Divider("-")

Step1:
	fmt.Print("Pilih menu : ")
	pilihan, err := reader.ReadString('\n')
	pilihan = strings.TrimSpace(pilihan)

	if err != nil || !utils.IsNumeric(pilihan) {
		utils.PrintMessage(fmt.Sprintf("Pilihan %s tidak ada", pilihan), "error")
		goto Step1
	}

	switch pilihan {
	case "1":
		utils.ClearScreen()
		LoginForUser()
	case "2":
		utils.ClearScreen()
		RegisterPenumpang()
	case "3":
		utils.ClearScreen()
		LoginForAdmin()
	default:
		utils.PrintMessage(fmt.Sprintf("Pilihan %s tidak ada", pilihan), "error")
		goto Step1
	}
}

func LoginForAdmin() {
	reader := bufio.NewReader(os.Stdin)

	utils.PrintBoxWithText(60, []string{
		"Login Sebagai Admin",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

Form:
	fmt.Print("Masukan username : ")
	username, err := reader.ReadString('\n')
	if username == "" || err != nil {
		utils.PrintMessage("Username tidak boleh kosong", "error")
		goto Form
	}

	fmt.Print("Masukan password : ")
	password, err := reader.ReadString('\n')
	if password == "" || err != nil {
		utils.PrintMessage("password tidak boleh kosong", "error")
		goto Form
	}

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	isNull := true
	for _, admin := range model.ListAdmin {
		if admin.Username == username && admin.Pass == password {
			AuthData = AuthStruct{
				is_admin: true,
				admin:    admin,
			}
			isNull = false
			utils.ClearScreen()
			MenuAwalAdmin()
			return
		}
	}

	if isNull == true {
		utils.PrintMessage("Username dan password salah", "error")
		utils.Divider("-")
		goto Form
	}
}

func RegisterPenumpang() {
	reader := bufio.NewReader(os.Stdin)
	utils.PrintBoxWithText(60, []string{
		"Daftar Pengguna",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

FormNama:
	fmt.Print("Masukan nama lengkap : ")
	nama, err := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama == "" || err != nil {
		utils.PrintMessage("Nama wajib di isi", "error")
		goto FormNama
	}
FormNIK:
	fmt.Print("Masukan NIK : ")
	nik, err := reader.ReadString('\n')
	nik = strings.TrimSpace(nik)
	if nik == "" || err != nil {
		utils.PrintMessage("NIK wajib di isi", "error")
		goto FormNIK
	}

	var nikRegex = regexp.MustCompile(`^\d{16}$`)
	if !nikRegex.MatchString(nik) {
		utils.PrintMessage("Format NIK salah", "error")
		goto FormNIK
	}

	dataNIK := utils.BinaryFindMany(model.ListUser, model.User{NIK: nik}, func(a, b model.User) int {
		if a.NIK < b.NIK {
			return -1
		} else if a.NIK > b.NIK {
			return 1
		}
		return 0
	})

	if len(dataNIK) > 0 {
		utils.PrintMessage("NIK sudah digunakan", "error")
		goto FormNIK
	}
FormKelamin:
	fmt.Print("Masukan Jenis Kelamin (l/p) : ")
	jenisKelamin, err := reader.ReadString('\n')
	jenisKelamin = strings.TrimSpace(jenisKelamin)
	if jenisKelamin == "" || err != nil {
		utils.PrintMessage("Wajib dipilih jenis kelamin", "error")
		goto FormKelamin
	}

	if jenisKelamin != "l" && jenisKelamin != "p" {
		utils.PrintMessage("Pilihan hanya l (Laki-Laki) dan p (Perempuan)", "error")
		goto FormKelamin
	}

FormAlamat:
	fmt.Print("Masukan Alamat : ")
	alamat, err := reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)
	if alamat == "" || err != nil {
		utils.PrintMessage("Alamat Wajib di isi", "error")
		goto FormAlamat
	}

FormTglLahir:
	fmt.Print("Masukan Tanggal Lahir  (DD/MM/YYYY) : ")
	TglLahir, err := reader.ReadString('\n')
	TglLahir = strings.TrimSpace(TglLahir)
	if TglLahir == "" || err != nil {
		utils.PrintMessage("Alamat Wajib di isi", "error")
		goto FormTglLahir
	}

	if !regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/[0-9]{4}$`).MatchString(TglLahir) {
		utils.PrintMessage("Format tanggal lahir salah", "error")
		goto FormTglLahir
	}

FormPekerjaan:
	fmt.Print("Masukan Pekerjaan : ")
	pekerjaan, err := reader.ReadString('\n')
	pekerjaan = strings.TrimSpace(pekerjaan)
	if pekerjaan == "" || err != nil {
		utils.PrintMessage("Pekerjaan wajib diisi", "error")
		goto FormPekerjaan
	}

FormNoHp:
	fmt.Print("Masukan No HP : ")
	noHp, err := reader.ReadString('\n')
	noHp = strings.TrimSpace(noHp)

	if noHp == "" || err != nil {
		utils.PrintMessage("No. Hp wajib diisi", "error")
		goto FormKelamin
	}

	if !regexp.MustCompile(`^08[0-9]{8,11}$`).MatchString(noHp) {
		utils.PrintMessage("Format No HP salah", "error")
		goto FormNoHp
	}
	dataNoHp := utils.BinaryFindMany(model.ListUser, model.User{NoHP: noHp}, func(a, b model.User) int {
		if a.NoHP < b.NoHP {
			return -1
		} else if a.NoHP > b.NoHP {
			return 1
		}
		return 0
	})
	if len(dataNoHp) > 0 {
		utils.PrintMessage("No HP sudah digunakan", "error")
		goto FormNoHp
	}

FormEmail:
	fmt.Print("Masukan Email : ")
	email, err := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	if email == "" || err != nil {
		utils.PrintMessage("Email wajib diisi", "error")
		goto FormKelamin
	}

	if !regexp.MustCompile(`^[\w.+-]+@[\w-]+\.[\w.-]+$`).MatchString(email) {
		utils.PrintMessage("Format Email salah", "error")
		goto FormEmail
	}
	dataEmail := utils.BinaryFindMany(model.ListUser, model.User{Email: email}, func(a, b model.User) int {
		if a.Email < b.Email {
			return -1
		} else if a.Email > b.Email {
			return 1
		}
		return 0
	})
	if len(dataEmail) > 0 {
		utils.PrintMessage("Email sudah digunakan", "error")
		goto FormEmail
	}

FormPassword:
	fmt.Print("Masukan Password : ")
	password, err := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if password == "" || err != nil {
		utils.PrintMessage("Password wajib diisi", "error")
		goto FormKelamin
	}

	if len(password) < 6 {
		utils.PrintMessage("Password minimal 6 karakter", "error")
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
	utils.ClearScreen()
	utils.PrintMessage("User Berhasil Terdaftar", "success")
	AuthController()
}

func LoginForUser() {
	render := bufio.NewReader(os.Stdin)

	utils.PrintBoxWithText(60, []string{
		"Login Sebagai Pengguna",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

FormUsername:
	fmt.Print("Masukan username (email/nomor hp) : ")
	username, err := render.ReadString('\n')
	username = strings.TrimSpace(username)
	if username == "" || err != nil {
		utils.PrintMessage("Username wajib di isi", "error")
		goto FormUsername
	}
FormPassword:
	fmt.Print("Masukan password : ")
	password, err := render.ReadString('\n')
	password = strings.TrimSpace(password)
	if password == "" || err != nil {
		utils.PrintMessage("Password wajib di isi", "error")
		goto FormPassword
	}

	isNull := true
	for _, user := range model.ListUser {
		if (user.Email == username || user.NoHP == username) && user.Password == password {
			AuthData = AuthStruct{
				is_admin: false,
				User:     user,
			}
			isNull = false
			utils.ClearScreen()
			MenuAwalUser()
			return
		}
	}

	if isNull {
		utils.PrintMessage("Username dan password salah", "error")
		utils.Divider("-")
		goto FormUsername
	}

}
