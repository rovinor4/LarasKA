package Admin

import (
	"fmt"
	"laraska/model"
	"laraska/utils"
	"os"
	"strconv"
	"time"
)

var SessionAdmin model.Admin

func Login() {
	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		"Login Admin",
	})

	var Admin model.Admin

	utils.Input("Username: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Username tidak boleh kosong"
		}

		Sorting := utils.InsertionSort(model.ListAdmin, func(a, b model.Admin) bool {
			return a.Username < b.Username
		})

		Find := false
		Admin, Find, _ = utils.BinaryFindOne(Sorting, model.Admin{Username: input}, func(a, b model.Admin) int {
			if a.Username < b.Username {
				return -1
			} else if a.Username > b.Username {
				return 1
			}
			return 0
		})

		if !Find {
			return false, "Username tidak ditemukan"
		}
		return true, ""
	})

	utils.Input("Password: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Password tidak boleh kosong"
		}

		if Admin.Pass != input {
			return false, "Password salah"
		}

		return true, ""
	})

	SessionAdmin = Admin
	utils.ClearScreen()
	MenuAwalAdmin()
}

func Logout() {
	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		"Anda telah keluar dari akun admin",
	})

	fmt.Println("Tekan enter untuk melanjutkan . . .")
	fmt.Scanln()

	utils.ClearScreen()
	SessionAdmin = model.Admin{}
	Login()
}

func MenuAwalAdmin() {

	now := time.Now()

	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		fmt.Sprintf("Hai, %s", SessionAdmin.Name),
		now.Format("01-02-2006 15:04 WIB"),
	})

	fmt.Println("[1] Menu Stasiun")
	fmt.Println("[2] Menu Kereta Api")
	fmt.Println("[3] Menu Rute")
	fmt.Println("[4] Menu  Tiket Kereta")
	fmt.Println("[5] Keluar Akun")
	fmt.Println("[6] Tutup Program")

	utils.Divider("-")
	InputSelect := utils.Input("Pilih menu: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "3", "4", "5", "6"}) {
			return false, "Input tidak valid, silakan pilih menu yang tersedia"
		}

		return true, ""
	})

	Select, _ := strconv.Atoi(InputSelect)

	utils.ClearScreen()
	switch Select {
	case 1:
		StasiunController()
	case 2:
		KeretaController()
	case 3:
		RuteController()
	case 4:
		TiketController()
	case 5:
		Logout()
	case 6:
		os.Exit(1)
		return
	}

}
