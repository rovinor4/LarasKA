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

	switch {
	case choice == "1":
		ClearScreen()
		TampilkanKereta()
	case choice == "2":
		ClearScreen()
		TambahKereta()
	case choice == "3":
		ClearScreen()
		EditKereta()
	case choice == "4":
		ClearScreen()
		HapusKereta()
	case choice == "5":
		ClearScreen()
		MenuAwalAdmin()
	default:
		PrintError("Pilihan tidak valid, silakan coba lagi.")
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

	PrintTable(
		[]string{"Kode", "Nama", "Kelas"},
		mapped,
		[]string{"Kode", "Nama", "Kelas"},
		4,
		Title,
	)
}

func TampilkanKereta() {

	var searchOn, sortOn, sortType string

	stop := false
	reader := bufio.NewReader(os.Stdin)

	for !stop {
		var menu string
		var mapped []map[string]string
		Data := model.ListKereta

		switch {
		case sortOn == "" && sortType == "":
			Data = InsertionSort(Data, func(a, b model.Kereta) bool { return a.Kode < b.Kode })
		case sortOn == "Kode" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Kereta) bool { return a.Kode < b.Kode })
		case sortOn == "Kode" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Kereta) bool { return a.Kode > b.Kode })
		case sortOn == "Nama" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Kereta) bool { return strings.ToLower(a.Nama) < strings.ToLower(b.Nama) })
		case sortOn == "Nama" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Kereta) bool { return strings.ToLower(a.Nama) > strings.ToLower(b.Nama) })
		case sortOn == "Kelas" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Kereta) bool { return strings.ToLower(a.Kelas) < strings.ToLower(b.Kelas) })
		case sortOn == "Kelas" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Kereta) bool { return strings.ToLower(a.Kelas) > strings.ToLower(b.Kelas) })
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

		PrintTable(
			[]string{"Kode", "Nama", "Kelas"},
			mapped,
			[]string{"Kode", "Nama", "Kelas"},
			4,
			Title,
		)

		fmt.Println(ColorText("[1] Pencarian", 90, 49, false))
		fmt.Println(ColorText("[2] Tampilkan Seluruh Data", 90, 49, false))
		fmt.Println(ColorText("[3] Sorting Data", 90, 49, false))
		fmt.Println(ColorText("[4] Kembali Ke Menu Stasiun", 90, 49, false))

		Pembatas("-")

		fmt.Print("Masukan nomor menu : ")
		_, err := fmt.Scan(&menu)
		if err != nil || !isNumeric(menu) {
			ClearScreen()
			PrintError("Pilihan tidak valid, silakan coba lagi.")
		}

		switch menu {
		case "1":
			fmt.Print("Masukan keyword pencarian : ")
			search, _ := reader.ReadString('\n')
			searchOn = strings.TrimSpace(search)
			ClearScreen()
		case "2":
			searchOn = ""
			sortOn = ""
			sortType = ""
			ClearScreen()
		case "3":
			fmt.Print("Pilih kolom untuk sort : ")
			fmt.Scan(&sortOn)
			fmt.Print("Pilih jenis sort (asc/desc) : ")
			fmt.Scan(&sortType)
			ClearScreen()
		case "4":
			stop = true
			ClearScreen()
			MenuKereta()
		default:
			ClearScreen()
			PrintError("Pilihan tidak valid, silakan coba lagi.")
		}

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
	TableKereta("Hapus Data Kereta")
Input:
	fmt.Print("Masukan Kode Kereta untuk dihapus: ")
	fmt.Scan(&kode)
	for i, kt := range model.ListKereta {
		if kt.Kode == kode {
			model.ListKereta = append(model.ListKereta[:i], model.ListKereta[i+1:]...)
			ClearScreen()
			fmt.Println("Kereta berhasil dihapus.")
			MenuKereta()
			return
		}
	}

	PrintError("Kode kereta tidak ditemukan.")
	goto Input

}
