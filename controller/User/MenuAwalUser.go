package User

import (
	"fmt"
	"laraska/controller"
	"laraska/model"
	"laraska/utils"
	"strconv"
)

func MenuAwalUser(authUser model.User) {

	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		fmt.Sprintf("Hai, %s", controller.AuthData.User.NamaLengkap),
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
		EditAkunUser(authUser)
		//TODO: cycle import
		// case 4:
		// 	utils.ClearScreen()
		// 	controller.AuthController()

	}

}
