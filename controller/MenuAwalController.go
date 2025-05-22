package controller

import (
	"fmt"
	"laraska/utils"
	"os"
	"time"
)

func MenuAwalAdmin() {

Step1:
	var input string
	now := time.Now()
	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		fmt.Sprintf("Hai, %s", AuthData.admin.Name),
		now.Format("01-02-2006 15:04 WIB"),
	})

	if AuthData.admin.Role == 1 {
		fmt.Println("[1] Data Stasiun")
		fmt.Println("[2] Data Kereta")
		fmt.Println("[3] Data Rute")
		fmt.Println("[4] Data User")
		fmt.Println("[5] Data Tiket")
		fmt.Println("[6] Keluar Akun")
		fmt.Println("[7] Tutup Program")

		utils.Divider("-")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&input)

		switch input {
		case "1":
			utils.ClearScreen()
			MenuStasiun()
		case "2":
			utils.ClearScreen()
			MenuKereta()
		case "3":
			utils.ClearScreen()
			RuteController()
		case "4":
			utils.ClearScreen()
		case "5":
			utils.ClearScreen()
		case "6":
			utils.ClearScreen()
			AuthController()
		case "7":
			os.Exit(1)
			return
		default:
			utils.ClearScreen()
			utils.PrintMessage("Menu tidak ada","error")
			goto Step1
		}
	}

}
