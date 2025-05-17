package controller

import (
	"fmt"
	"time"
)

func MenuAwal() {
	var input string
	now := time.Now()

	var menuList = []string{
		"Data Stasiun",
		"Data Penumpang",
		"Data Rute",
		"Data Kereta",
		"Data Tiket",
		"Keluar Akun",
		"Tutup Program",
	}

	PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		now.Format("01-02-2006 15:04 WIB"),
	})

	for index, menu := range menuList {
		fmt.Printf("[%d] %s\n", index+1, menu)
	}

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
