package controller

import (
	"fmt"
	"os"
	"time"
)

func MenuAwalAdmin() {

Step1:
	var input string
	now := time.Now()
	PrintBoxWithText(60, []string{
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

		Pembatas("-")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&input)

		switch input {
		case "1":
			ClearScreen()
			MenuStasiun()
		case "2":
			ClearScreen()
			MenuKereta()
		case "3":
			ClearScreen()
			MenuRute()
		case "4":
			ClearScreen()
		case "5":
			ClearScreen()
		case "6":
			ClearScreen()
			AuthController()
		case "7":
			os.Exit(1)
			return
		default:
			ClearScreen()
			PrintError("Menu tidak ada")
			goto Step1
		}
	}

}
