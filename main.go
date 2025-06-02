package main

import (
	"fmt"
	"laraska/controller/Admin"
	"laraska/controller/User"
	"laraska/utils"
	"os"
	"strconv"
)

func main() {

	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		"Telkom University Surabaya | Informatika",
	})

	utils.PrintBoxWithText(60, []string{
		"Kelompok 2 : ",
		"Rovino Ramadhani (103072400031)",
		"Setyo Nugroho (103072400045)",
		"Rangga Dani Prasetya (103072400057)",
	})

	var input string
	fmt.Print("Masukan x untuk menjalankan program : ")
	_, err := fmt.Scan(&input)

	if err != nil || input != "x" {
		fmt.Println("Invalid input. Please press x to run the program.")
		os.Exit(0)
	} else {

		utils.ClearScreen()

		utils.PrintBoxWithText(60, []string{
			"Selamat Datang",
			"LarasKA (Layanan Reservasi Kereta Api)",
		})

		fmt.Println("[1] Login Sebagai Pengguna")
		fmt.Println("[2] Daftar Pengguna Baru")
		fmt.Println("[3] Login Sebagai Admin")

		utils.Divider("-")

		InputSelect := utils.Input("Pilih menu: ", func(input string) (bool, string) {
			if input == "" {
				return false, "Pilihan tidak boleh kosong"
			}

			if !utils.IsNumeric(input) {
				return false, "Input harus berupa angka"
			}

			if !utils.IsIn(input, []string{"1", "2", "3"}) {
				return false, "Pilihan menu tidak tersedia"
			}
			return true, ""
		})

		Select, _ := strconv.Atoi(InputSelect)
		utils.ClearScreen()
		switch Select {
		case 1:
			User.Login()
		case 2:
			User.Register()
		case 3:
			Admin.Login()
		}
		utils.ClearScreen()
	}
}
