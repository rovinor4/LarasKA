package main

import (
	"fmt"
	"laraska/controller"
	"laraska/utils"
	"os"
)

func main() {
 
	utils.PrintMessage("Selamat Datang di LarasKA", "info")
	return
	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		"Telkom University Surabaya | Informatika",
	})

	utils.PrintBoxWithText(60, []string{
		"Kelompok 2 : ",
		"Rovino Ramadhani (103072400031)",
		"Rangga Dani Prasetya (103072400057)",
		"Setyo Nugroho (103072400045)",
	})

	var input string
	fmt.Print("Masukan x untuk menjalankan program : ")
	_, err := fmt.Scan(&input)

	if err != nil || input != "x" {
		fmt.Println("Invalid input. Please press x to run the program.")
		os.Exit(0)
	} else {
		utils.ClearScreen()
		controller.AuthController()
	}
}
