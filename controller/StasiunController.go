package controller

import (
	"bufio"
	"fmt"
	"laraska/model"
	"os"
	"strings"
)

func MenuStasiun() {

	var choice string
	var menuList = []string{
		"Tampilkan Stasiun",
		"Tambah Stasiun",
		"Edit Stasiun",
		"Hapus Stasiun",
		"Kembali ke Menu Awal",
	}

	PrintHead("Menu Stasiun")
	for index, menu := range menuList {
		fmt.Printf("[%d] %s\n", index+1, menu)
	}
	Divider("-")

	fmt.Print("Pilih menu: ")
	_, err := fmt.Scan(&choice)
	if err != nil || !IsNumeric(choice) {
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		MenuStasiun()
		return
	}

	switch choice {
	case "1":
		ClearScreen()
		TampilkanStasiun()
	case "2":
		ClearScreen()
		TambahStasiun()
	case "3":
		ClearScreen()
		EditStasiun()
	case "4":
		ClearScreen()
		HapusStasiun()
	case "5":
		ClearScreen()
		MenuAwalAdmin()
	default:
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		MenuStasiun()
	}

}

func TambahStasiun() {
	reader := bufio.NewReader(os.Stdin)
	var statiun model.Stasiun

	PrintHead("Tambah Stasiun")
	fmt.Print("Id Stasiun: ")
	fmt.Scan(&statiun.IDStasiun)

	for _, st := range model.ListStasiun {
		if st.IDStasiun == statiun.IDStasiun {
			PrintError("ID Stasiun sudah ada, silakan masukkan ID yang berbeda.")
			TambahStasiun()
			return
		}
	}

	fmt.Print("Nama Stasiun: ")
	inputName, err := reader.ReadString('\n')
	if err != nil {
		PrintError("Error reading input")
		return
	}
	statiun.Nama = strings.TrimSpace(inputName)

	fmt.Print("Kota : ")
	fmt.Scan(&statiun.Kota)

	model.ListStasiun = append(model.ListStasiun, statiun)
	ClearScreen()
	fmt.Println(ColorText("Stasiun berhasil ditambahkan.", 30, 42, false))
	MenuStasiun()
}

func TampilkanStasiun(search ...string) {
	var searchQuery, choice string

	if len(search) > 0 {
		searchQuery = search[0]
	}

	var mapped []map[string]string
	for _, st := range model.ListStasiun {
		if searchQuery == "" || strings.Contains(strings.ToLower(fmt.Sprintf("%s %s %s", st.Nama, st.IDStasiun, st.Kota)), strings.ToLower(searchQuery)) {
			mapped = append(mapped, map[string]string{
				"IDStasiun": st.IDStasiun,
				"Nama":      st.Nama,
				"Kota":      st.Kota,
			})
		}
	}

	PrintTable(
		[]string{"ID", "Nama", "Kota"},
		mapped,
		[]string{"IDStasiun", "Nama", "Kota"},
		4,
		"Data Stasiun",
	)

	fmt.Println(ColorText("[1] Pencarian", 90, 49, false))
	fmt.Println(ColorText("[2] Tampilkan Semua", 90, 49, false))
	fmt.Println(ColorText("[3] Kembali Ke Menu Stasiun", 90, 49, false))

	fmt.Print("Pilih menu: ")
	_, err := fmt.Scan(&choice)

	if err != nil || !IsNumeric(choice) {
		PrintError("Pilihan tidak valid, silakan coba lagi.\n")
		TampilkanStasiun()
	}

	switch choice {
	case "1":
		fmt.Print("Masukkan nama stasiun yang dicari: ")
		fmt.Scan(&searchQuery)
		ClearScreen()
		TampilkanStasiun(searchQuery)
	case "2":
		ClearScreen()
		TampilkanStasiun()
	case "3":
		ClearScreen()
		MenuStasiun()
	default:
		PrintError("Pilihan tidak valid, silakan coba lagi.\n")
		TampilkanStasiun()
	}

	ClearScreen()
	MenuStasiun()
}

func ShowListStasiun() {
	for index, st := range model.ListStasiun {
		fmt.Printf("%d. %s (%s) - Kota %s \n", index+1, st.Nama, st.IDStasiun, st.Kota)
	}
}

func HapusStasiun() {
	PrintHead("Hapus Stasiun")
	ShowListStasiun()
	Divider("-")

	var choice int
	fmt.Print("Pilih nomor stasiun yang ingin dihapus: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(model.ListStasiun) {
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		HapusStasiun()
		return
	}

	var confirm string
	fmt.Print("Apakah Anda yakin ingin menghapus stasiun ini? (y/n): ")
	fmt.Scan(&confirm)

	if strings.ToLower(confirm) != "y" {
		ClearScreen()
		fmt.Println(ColorText("Penghapusan dibatalkan.", 30, 41, false))
		MenuStasiun()
		return
	}

	model.ListStasiun = append(model.ListStasiun[:choice-1], model.ListStasiun[choice:]...)
	fmt.Println("Stasiun berhasil dihapus!")

	ClearScreen()
	fmt.Println(ColorText("Stasiun berhasil dihapus.", 30, 42, false))
	MenuStasiun()
}

func EditStasiun() {
	var pilihIndex int

	PrintHead("Edit Stasiun")
	ShowListStasiun()
	Divider("-")

	fmt.Print("Pilih nomor stasiun yang ingin di edit: ")
	fmt.Scan(&pilihIndex)

	if pilihIndex < 1 || pilihIndex > len(model.ListStasiun) {
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		EditStasiun()
		return
	}

	stasiunDipilih := &model.ListStasiun[pilihIndex-1]

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Nama Stasiun (%s): ", stasiunDipilih.Nama)
	inputNama, _ := reader.ReadString('\n')
	inputNama = strings.TrimSpace(inputNama)
	if inputNama != "" {
		stasiunDipilih.Nama = inputNama
	}

	fmt.Printf("Kota (%s): ", stasiunDipilih.Kota)
	inputKota, _ := reader.ReadString('\n')
	inputKota = strings.TrimSpace(inputKota)
	if inputKota != "" {
		stasiunDipilih.Nama = inputKota
	}

	ClearScreen()
	fmt.Println(ColorText("Stasiun berhasil diperbarui.", 30, 42, false))
	MenuStasiun()
}
