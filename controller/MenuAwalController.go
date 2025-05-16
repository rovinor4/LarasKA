package controller

import (
	"fmt"
	"time"
)

func MenuAwal() {
	var input string
	now := time.Now()

	PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		now.Format("01-02-2006 15:04 WIB"),
	})
	fmt.Println("1. Data Stasiun")
	fmt.Println("2. Data Penumpang")
	fmt.Println("3. Data Rute")
	fmt.Println("4. Data Kereta")
	fmt.Println("5. Data Tiket")
	fmt.Println("6. Data Admin")
	fmt.Println("7. Keluar")
	Pembatas("-")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&input)

	switch input {
	case "1":
		ClearScreen()
		MenuStasiun()
	default:
		PrintError("Menu tidak ada")
		MenuAwal()
	}

}
