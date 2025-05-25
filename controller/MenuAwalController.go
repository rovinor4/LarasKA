package controller

import (
	"bufio"
	"fmt"
	"laraska/utils"
	"os"
	"strings"
	"time"
)

func MenuAwalAdmin() {

	var input string
	now := time.Now()
	render := bufio.NewReader(os.Stdin)

	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		fmt.Sprintf("Hai, %s", AuthData.admin.Name),
		now.Format("01-02-2006 15:04 WIB"),
	})

	fmt.Println("[1] Menu Stasiun")
	fmt.Println("[2] Menu Kereta Api")
	fmt.Println("[3] Menu Rute")
	fmt.Println("[4] Menu User")
	fmt.Println("[5] Menu Reservasi Tiket Kereta")
	fmt.Println("[6] Keluar Akun")
	fmt.Println("[7] Tutup Program")

PilihMenu:
	utils.Divider("-")
	fmt.Print("Pilih menu: ")
	input, err := render.ReadString('\n')
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan saat membaca input", "error")
		return
	}
	input = strings.TrimSpace(input)
	if input == "" {
		utils.PrintMessage("Input tidak boleh kosong", "error")
		goto PilihMenu
	}

	switch input {
	case "1":
		utils.ClearScreen()
		StasiunController()
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
		utils.PrintMessage("Menu tidak ada", "error")
		goto PilihMenu
	}

}
