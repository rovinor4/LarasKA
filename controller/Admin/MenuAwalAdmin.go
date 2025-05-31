package Admin

import (
	"fmt"
	"laraska/controller"
	"laraska/utils"
	"os"
	"strconv"
	"time"
)

func MenuAwalAdmin() {

	now := time.Now()

	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		fmt.Sprintf("Hai, %s", controller.AuthData.Admin),
		now.Format("01-02-2006 15:04 WIB"),
	})

	fmt.Println("[1] Menu Stasiun")
	fmt.Println("[2] Menu Kereta Api")
	fmt.Println("[3] Menu Rute")
	fmt.Println("[4] Menu User")
	fmt.Println("[5] Menu Reservasi Tiket Kereta")
	fmt.Println("[6] Keluar Akun")
	fmt.Println("[7] Tutup Program")

	InputSelect := utils.Input("Pilih menu: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "3", "4", "5", "6", "7"}) {
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
	case 5:
	case 6:
		controller.AuthController()
	case 7:
		os.Exit(1)
		return
	}

}
