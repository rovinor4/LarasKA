package User

import (
	"fmt"
	"laraska/model"
	"laraska/utils"
	"strconv"
	"time"
)

func Login() {
	var User model.User
	utils.PrintBoxWithText(60, []string{
		"Login Sebagai Pengguna",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

	utils.Input("Masukan username (email/nomor hp): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Username wajib di isi"
		}

		Have := false
		Stop := false
		for i := 0; i < len(model.ListUser) && !Stop; i++ {
			if model.ListUser[i].Email == input || model.ListUser[i].NoHP == input {
				Stop = true
				Have = true
				User = model.ListUser[i]
			}
		}

		if !Have {
			return false, "Username tidak ditemukan"
		}

		return true, ""
	})

	utils.Input("Masukan password: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Password wajib di isi"
		}

		if User.Password != input {
			return false, "Password salah"
		}

		return true, ""
	})

	utils.ClearScreen()
	MenuAwalUser(User)
}

func Register() {
	utils.PrintBoxWithText(60, []string{
		"Daftar Pengguna Baru",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

	var User model.User

	NamaLengkap := utils.Input("Masukan nama lengkap: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Nama lengkap wajib diisi"
		}
		return true, ""
	})

	NIK := utils.Input("Masukan NIK (16 digit): ", func(input string) (bool, string) {
		if input == "" {
			return false, "NIK wajib diisi"
		}
		if len(input) != 16 {
			return false, "NIK harus 16 digit"
		}

		for _, user := range model.ListUser {
			if user.NIK == input {
				return false, "NIK sudah digunakan"
			}
		}

		return true, ""
	})

	JenisKelamin := utils.Input("Masukan Jenis Kelamin (l/p): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Jenis kelamin wajib diisi"
		}
		if input != "l" && input != "p" {
			return false, "Pilihan hanya l (Laki-Laki) dan p (Perempuan)"
		}
		return true, ""
	})

	Alamat := utils.Input("Masukan Alamat: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Alamat wajib diisi"
		}
		return true, ""
	})

	TglLahir := utils.Input("Masukan Tanggal Lahir (DD/MM/YYYY): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Tanggal lahir wajib diisi"
		}
		if _, err := time.Parse("02/01/2006", input); err != nil {
			return false, "Format tanggal lahir salah, gunakan DD/MM/YYYY"
		}
		return true, ""
	})

	Pekerjaan := utils.Input("Masukan Pekerjaan: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Pekerjaan wajib diisi"
		}
		return true, ""
	})

	NoHP := utils.Input("Masukan No HP (08XXXXXXXXXX): ", func(input string) (bool, string) {
		if input == "" {
			return false, "No HP wajib diisi"
		}
		if len(input) < 10 || len(input) > 13 || input[:2] != "08" {
			return false, "Format No HP salah, harus diawali 08 dan 10-13 digit"
		}

		SortingData := utils.InsertionSort(model.ListUser, func(a, b model.User) bool {
			return a.NoHP < b.NoHP
		})

		_, found, _ := utils.BinaryFindOne(SortingData, model.User{NoHP: input}, func(a, b model.User) int {
			if a.NoHP < b.NoHP {
				return -1
			} else if a.NoHP > b.NoHP {
				return 1
			}
			return 0
		})
		if found {
			return false, "No HP sudah digunakan"
		}

		return true, ""
	})

	Email := utils.Input("Masukan Email: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Email wajib diisi"
		}
		if !utils.IsValidEmail(input) {
			return false, "Format Email salah"
		}

		SortingData := utils.InsertionSort(model.ListUser, func(a, b model.User) bool {
			return a.Email < b.Email
		})
		_, found, _ := utils.BinaryFindOne(SortingData, model.User{Email: input}, func(a, b model.User) int {
			if a.Email < b.Email {
				return -1
			} else if a.Email > b.Email {
				return 1
			}
			return 0
		})
		if found {
			return false, "Email sudah digunakan"
		}

		return true, ""
	})

	Password := utils.Input("Masukan Password (min 6 karakter): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Password wajib diisi"
		}
		if len(input) < 6 {
			return false, "Password minimal 6 karakter"
		}
		return true, ""
	})

	User = model.User{
		NamaLengkap:  NamaLengkap,
		NIK:          NIK,
		JenisKelamin: JenisKelamin,
		Alamat:       Alamat,
		TglLahir:     TglLahir,
		Pekerjaan:    Pekerjaan,
		NoHP:         NoHP,
		Email:        Email,
		Password:     Password,
	}
	model.ListUser = append(model.ListUser, User)
	utils.ClearScreen()
	utils.PrintMessage("User Berhasil Terdaftar", "success")
	Login()
}

func MenuAwalUser(authUser model.User) {

	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		fmt.Sprintf("Hai, %s", authUser.NamaLengkap),
	})

	fmt.Println("[1] Pesan Tiket")
	fmt.Println("[2] History Pemesanan Tiket")
	fmt.Println("[3] Edit Akun")
	fmt.Println("[4] Log out")

	utils.Divider("-")

	InputSelect := utils.Input("Pilih menu: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "3", "4"}) {
			return false, "Pilihan menu tidak tersedia"
		}

		return true, ""
	})

	input, _ := strconv.Atoi(InputSelect)

	switch input {
	case 1:
		utils.ClearScreen()
		PesanTiket(authUser)
	case 2:
		utils.ClearScreen()
		ShowHistoryTiket(authUser)
	case 3:
		utils.ClearScreen()
		EditAkunUser(&authUser)
	case 4:
		authUser = model.User{}
		utils.ClearScreen()
		Login()
	}

}
