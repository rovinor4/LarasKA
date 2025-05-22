package controller

import (
	"bufio"
	"fmt"
	"laraska/model"
	"laraska/utils"
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

	utils.PrintHead("Menu Kereta")

	for index, menu := range menuList {
		fmt.Printf("[%d] %s\n", index+1, menu)
	}
	utils.Divider("-")

	fmt.Print("Pilih menu: ")
	_, err := fmt.Scan(&choice)
	if err != nil || !utils.IsNumeric(choice) {
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
		MenuKereta()
		return
	}

	switch {
	case choice == "1":
		TampilkanKereta()
	case choice == "2":
		TambahKereta()
	case choice == "3":
		EditKereta()
	case choice == "4":
		HapusKereta()
	case choice == "5":
		MenuAwalAdmin()
	default:
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
		MenuKereta()
	}
}

func TableKereta(Title string) {
	var mapped []map[string]string
	for _, dt := range model.ListKereta {
		mapped = append(mapped, map[string]string{
			"Kode":  strconv.Itoa(dt.Kode),
			"Nama":  dt.Nama,
			"Kelas": dt.Kelas,
		})
	}

	utils.PrintTable(
		[]string{"Kode", "Nama", "Kelas"},
		mapped,
		[]string{"Kode", "Nama", "Kelas"},
		4,
		Title,
	)
}

func TampilkanKereta() {

	var searchOn, sortOn, sortType string

	utils.ClearScreen()
	reader := bufio.NewReader(os.Stdin)
Step1:
	var menu string
	var mapped []map[string]string
	Data := model.ListKereta

	switch {
	case sortOn == "" && sortType == "":
		Data = utils.InsertionSort(Data, func(a, b model.Kereta) bool { return a.Kode < b.Kode })
	case sortOn == "Kode" && sortType == "asc":
		Data = utils.InsertionSort(Data, func(a, b model.Kereta) bool { return a.Kode < b.Kode })
	case sortOn == "Kode" && sortType == "desc":
		Data = utils.InsertionSort(Data, func(a, b model.Kereta) bool { return a.Kode > b.Kode })
	case sortOn == "Nama" && sortType == "asc":
		Data = utils.InsertionSort(Data, func(a, b model.Kereta) bool { return strings.ToLower(a.Nama) < strings.ToLower(b.Nama) })
	case sortOn == "Nama" && sortType == "desc":
		Data = utils.InsertionSort(Data, func(a, b model.Kereta) bool { return strings.ToLower(a.Nama) > strings.ToLower(b.Nama) })
	case sortOn == "Kelas" && sortType == "asc":
		Data = utils.InsertionSort(Data, func(a, b model.Kereta) bool { return strings.ToLower(a.Kelas) < strings.ToLower(b.Kelas) })
	case sortOn == "Kelas" && sortType == "desc":
		Data = utils.InsertionSort(Data, func(a, b model.Kereta) bool { return strings.ToLower(a.Kelas) > strings.ToLower(b.Kelas) })
	}

	for _, dt := range Data {
		if searchOn == "" || strings.Contains(strings.ToLower(fmt.Sprintf("%s %s %d", dt.Nama, dt.Kelas, dt.Kode)), strings.ToLower(searchOn)) {
			mapped = append(mapped, map[string]string{
				"Kode":  strconv.Itoa(dt.Kode),
				"Nama":  dt.Nama,
				"Kelas": dt.Kelas,
			})
		}

	}

	Title := "Data Kereta"
	if searchOn != "" {
		Title = fmt.Sprintf("Hasil Pencarian Data Kereta : %s", searchOn)
	}

	utils.PrintTable(
		[]string{"Kode", "Nama", "Kelas"},
		mapped,
		[]string{"Kode", "Nama", "Kelas"},
		4,
		Title,
	)

	fmt.Println(utils.ColorText("[1] Pencarian", 90, 49, false))
	fmt.Println(utils.ColorText("[2] Tampilkan Seluruh Data", 90, 49, false))
	fmt.Println(utils.ColorText("[3] Sorting Data", 90, 49, false))
	fmt.Println(utils.ColorText("[0] Kembali", 90, 49, false))

	utils.Divider("-")
Step2:
	fmt.Print("Masukan nomor menu : ")
	_, err := fmt.Scan(&menu)
	if err != nil || !utils.IsNumeric(menu) {
		utils.ClearScreen()
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
	}

	switch menu {
	case "1":
	Step2Search:
		fmt.Print("Masukan keyword pencarian : ")
		search, _ := reader.ReadString('\n')
		searchOn = strings.TrimSpace(search)

		if searchOn == "" {
			utils.PrintMessage("Pencarian Tidak Boleh Kosong", "error")
			goto Step2Search
		}

		utils.ClearScreen()
		goto Step1

	case "2":
		searchOn = ""
		sortOn = ""
		sortType = ""
		utils.ClearScreen()
		goto Step1
	case "3":
	Step2Sort:
		fmt.Print("Pilih kolom untuk sort (Kode/Nama/Kelas) : ")
		fmt.Scan(&sortOn)
		if sortOn != "Kode" && sortOn != "Nama" && sortOn != "Kelas" {
			utils.PrintMessage("Pastikan kamu memilih (Kode/Nama/Kelas)", "error")
			goto Step2Sort
		}
	Step2SortSelect:
		fmt.Print("Pilih jenis sort (asc/desc) : ")
		fmt.Scan(&sortType)

		if sortType != "asc" && sortType != "desc" {
			utils.PrintMessage("Pastikan kamu memilih (asc/desc)", "error")
			goto Step2SortSelect
		}

		utils.ClearScreen()
		goto Step1
	case "4":
		utils.ClearScreen()
		MenuKereta()
	default:
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
		goto Step2
	}

}

func TambahKereta() {
	reader := bufio.NewReader(os.Stdin)
	var kode int
	utils.PrintHead("Tambah Data Kereta")
	fmt.Print("Masukan Kode Kereta: ")
	_, err := fmt.Scan(&kode)

	if err != nil {
		utils.PrintMessage("Hanya di izinakan berupa angka", "error")
		TambahKereta()
	}

	fmt.Print("Masukan Nama Kereta: ")
	nama, _ := reader.ReadString('\n')
	fmt.Print("Masukan Kelas Kereta: ")
	kelas, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	kelas = strings.TrimSpace(kelas)

	if nama == "" || kelas == "" {
		utils.PrintMessage("Nama kereta atau kelas tidak boleh kosong", "error")
		TambahKereta()
		return
	}

	baru := model.Kereta{Kode: kode, Nama: nama, Kelas: kelas}
	model.ListKereta = append(model.ListKereta, baru)
	utils.ClearScreen()
	fmt.Println(utils.ColorText("Kereta berhasil ditambahkan.", 30, 42, false))
	MenuKereta()

}

func EditKereta() {
	reader := bufio.NewReader(os.Stdin)
	var kode int
	TableKereta("Edit Data Kereta")
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

			utils.ClearScreen()
			fmt.Println(utils.ColorText("Kereta berhasil diubah.", 30, 42, false))
			MenuKereta()
			return
		}
	}
	utils.PrintMessage("Kode kereta tidak ditemukan.", "error")
	MenuKereta()
}

func HapusKereta() {
	var kode int
	TableKereta("Hapus Data Kereta")
Input:
	fmt.Print("Masukan Kode Kereta untuk dihapus: ")
	fmt.Scan(&kode)
	for i, kt := range model.ListKereta {
		if kt.Kode == kode {
			model.ListKereta = append(model.ListKereta[:i], model.ListKereta[i+1:]...)
			utils.ClearScreen()
			fmt.Println("Kereta berhasil dihapus.")
			MenuKereta()
			return
		}
	}

	utils.PrintMessage("Kode kereta tidak ditemukan.", "error")
	goto Input

}
