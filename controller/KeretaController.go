package controller

import (
	"bufio"
	"fmt"
	"laraska/model"
	"os"
	"strconv"
	"strings"
)

func MenuKereta() {
	var choice string
	var menuList = []string{
		"Tampilkan Kereta",
		"Tambah Kereta",
		"Edit Kereta",
		"Hapus Kereta",
		"Kembali ke Menu Awal",
	}

	PrintJudul("Menu Kereta")

	for index, menu := range menuList {
		fmt.Printf("[%d] %s\n", index+1, menu)
	}
	Pembatas("-")

	fmt.Print("Pilih menu: ")
	_, err := fmt.Scan(&choice)
	if err != nil || !isNumeric(choice) {
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		MenuKereta()
		return
	}

	switch choice {
	case "1":
		ClearScreen()
		TampilkanKereta()
	case "2":
		ClearScreen()
		TambahKereta()
	case "3":
		ClearScreen()
		EditKereta()
	case "4":
		ClearScreen()
		HapusKereta()
	case "5":
		ClearScreen()
		return
	default:
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		MenuKereta()
	}
}

func tableKereta(Title string, searchPrams ...string) {
	var mapped []map[string]string
	var searchQuery string

	if len(searchPrams) > 0 {
		searchQuery = searchPrams[0]
	}

	for _, dt := range model.ListKereta {
		if searchQuery == "" || strings.Contains(strings.ToLower(fmt.Sprintf("%s %s %d", dt.Nama, dt.Kelas, dt.Kode)), strings.ToLower(searchQuery)) {
			mapped = append(mapped, map[string]string{
				"Kode":  strconv.Itoa(dt.Kode),
				"Nama":  dt.Nama,
				"Kelas": dt.Kelas,
			})
		}

	}

	PrintTable(
		[]string{"Kode", "Nama", "Kelas"},
		mapped,
		[]string{"Kode", "Nama", "Kelas"},
		4,
		Title,
	)
}

func TampilkanKereta(search ...string) {

	var menu, searchOn string

	if len(search) > 0 {
		searchOn = search[0]
	}

	reader := bufio.NewReader(os.Stdin)
	tableKereta("Tampilkan Data Kereta", searchOn)
	fmt.Println(ColorText("[1] Pencarian", 90, 49, false))
	fmt.Println(ColorText("[2] Tampilkan Seluruh Data", 90, 49, false))
	fmt.Println(ColorText("[3] Kembali Ke Menu Stasiun", 90, 49, false))
	fmt.Print("Masukan nomor menu : ")
	_, err := fmt.Scan(&menu)

	if err != nil || !isNumeric(menu) {
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		TampilkanKereta("")
	}

	switch menu {
	case "1":
		fmt.Print("Masukan keyword pencarian : ")
		search, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		ClearScreen()
		TampilkanKereta(strings.TrimSpace(search))
	case "2":
		ClearScreen()
		TampilkanKereta()
	case "3":
		ClearScreen()
		MenuKereta()
	default:
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		TampilkanKereta()
	}

}

func TambahKereta() {
	reader := bufio.NewReader(os.Stdin)
	var kode int
	PrintJudul("Tambah Data Kereta")
	fmt.Print("Masukan Kode Kereta: ")
	_, err := fmt.Scan(&kode)

	if err != nil {
		PrintError("Hanya di izinakan berupa angka")
		TambahKereta()
	}

	fmt.Print("Masukan Nama Kereta: ")
	nama, _ := reader.ReadString('\n')
	fmt.Print("Masukan Kelas Kereta: ")
	kelas, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	kelas = strings.TrimSpace(kelas)

	if nama == "" || kelas == "" {
		PrintError("Nama kereta atau kelas tidak boleh kosong")
		TambahKereta()
		return
	}

	baru := model.Kereta{Kode: kode, Nama: nama, Kelas: kelas}
	model.ListKereta = append(model.ListKereta, baru)
	ClearScreen()
	fmt.Println(ColorText("Kereta berhasil ditambahkan.", 30, 42, false))
	MenuKereta()

}

func EditKereta() {
	reader := bufio.NewReader(os.Stdin)
	var kode int
	tableKereta("Edit Data Kereta")
	fmt.Print("Masukan Kode Kereta untuk diubah: ")
	fmt.Scan(&kode)
	for i, kt := range model.ListKereta {
		if kt.Kode == kode {
			fmt.Printf("Masukan Nama Kereta baru (%s): ", kt.Nama)
			nama, _ := reader.ReadString('\n')
			fmt.Printf("Masukan Kelas Kereta baru (%s): ", kt.Kelas)
			kelas, _ := reader.ReadString('\n')
			if nama != "" {
				model.ListKereta[i].Nama = strings.TrimSpace(nama)
			}

			if kelas != "" {
				model.ListKereta[i].Kelas = strings.TrimSpace(kelas)
			}

			ClearScreen()
			fmt.Println(ColorText("Kereta berhasil diubah.", 30, 42, false))
			MenuKereta()
			return
		}
	}
	PrintError("Kode kereta tidak ditemukan.")
	MenuKereta()
}

func HapusKereta() {
	var kode int
	fmt.Print("Masukan Kode Kereta untuk dihapus: ")
	fmt.Scan(&kode)
	for i, kt := range model.ListKereta {
		if kt.Kode == kode {
			model.ListKereta = append(model.ListKereta[:i], model.ListKereta[i+1:]...)
			fmt.Println("Kereta berhasil dihapus.")
			TampilkanKereta()
			return
		}
	}
	PrintError("Kode kereta tidak ditemukan.")
	MenuKereta()
}
