package Admin

import (
	"fmt"
	"laraska/utils"
	"strconv"
)

func TiketController() {
	var pilihan int

	utils.PrintHead("Menu Tiket Kereta Api")

	fmt.Println("[1] Tambah Tiket")
	fmt.Println("[2] Cek Tiket")
	fmt.Println("[3] Hapus Tiket")
	fmt.Println("[0] Kembali ke Menu Utama")

	utils.Divider("-")
	InputPilihan := utils.Input("Masukkan Pilihan: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Pilihan tidak boleh kosong"
		}

		if !utils.IsNumeric(input) {
			return false, "Pilihan harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "3", "0"}) {
			return false, "Pilihan tidak valid"
		}

		return true, ""
	})

	pilihan, _ = strconv.Atoi(InputPilihan)

	switch pilihan {
	case 1:

	}
}

func AddTiket() {

}
