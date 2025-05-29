package user

import (
	"bufio"
	"fmt"
	"laraska/controller"
	"laraska/utils"
	"os"
	"strings"
	"time"
)

func MenuAwalUser() {
	reader := bufio.NewReader(os.Stdin)

	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		fmt.Sprintf("Hai, %s", controller.AuthData.User.NamaLengkap),
		time.Now().Format("01-02-2006 15:04 WIB"),
	})

UserMenu:
	fmt.Println("[1] Pesan Tiket")
	fmt.Println("[2] History Pemesanan Tiket")
	fmt.Println("[3] Log out")

	utils.Divider("-")

	fmt.Print("Pilih menu: ")
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err != nil {
		utils.PrintMessage("Terjadi kesalahan saat membaca input", "error")
		return
	}
	if input == "" {
		utils.PrintMessage("Input tidak boleh kosong", "error")
		goto UserMenu
	}

	switch input {
	case "1":
		utils.ClearScreen()
		PesanTiket()
	case "2":
		utils.ClearScreen()
		ShowHistoryTiket()
	case "3":
		utils.ClearScreen()
		controller.AuthController()
	default:
		utils.ClearScreen()
		utils.PrintMessage("Menu tidak ada", "error")
		goto UserMenu

	}

}
