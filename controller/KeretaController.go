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

func KeretaController() {
	var choice string

	utils.PrintHead("Menu Kereta")

	fmt.Println("[1] Tampilkan Data Kereta")
	fmt.Println("[2] Tambah Data Kereta")
	fmt.Println("[3] Edit Data Kereta")
	fmt.Println("[4] Hapus Data Kereta")
	fmt.Println("[5] Kembali Ke Menu Awal")
	utils.Divider("-")

	fmt.Print("Pilih menu: ")
	_, err := fmt.Scan(&choice)
	if err != nil || !utils.IsNumeric(choice) {
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
		KeretaController()
		return
	}

	switch {
	case choice == "1":
		showKereta()
	case choice == "2":
		addKereta()
	case choice == "3":
		updateKereta()
	case choice == "4":
		deleteKereta()
	case choice == "5":
		MenuAwalAdmin()
	default:
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
		KeretaController()
	}
}

func showKereta() {

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

	utils.PrintTable(
		[]string{"Kode", "Nama", "Kelas"},
		mapped,
		[]string{"Kode", "Nama", "Kelas"},
		4,
		"Data Kereta",
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

Step3:
	switch menu {
	case "1":
		fmt.Print("Masukan keyword pencarian : ")
		search, _ := reader.ReadString('\n')
		searchOn = strings.TrimSpace(search)

		if searchOn == "" {
			utils.PrintMessage("Pencarian Tidak Boleh Kosong", "error")
			goto Step3
		}

		utils.ClearScreen()
		goto Step1

	case "2":
		searchOn = ""
		sortOn = ""
		sortType = ""
		utils.ClearScreen()
		goto Step3
	case "3":
		fmt.Print("Pilih kolom untuk sort (Kode/Nama/Kelas) : ")
		fmt.Scan(&sortOn)
		if sortOn != "Kode" && sortOn != "Nama" && sortOn != "Kelas" {
			utils.PrintMessage("Pastikan kamu memilih (Kode/Nama/Kelas)", "error")
			goto Step3
		}
		fmt.Print("Pilih jenis sort (asc/desc) : ")
		fmt.Scan(&sortType)

		if sortType != "asc" && sortType != "desc" {
			utils.PrintMessage("Pastikan kamu memilih (asc/desc)", "error")
			goto Step3
		}

		utils.ClearScreen()
		goto Step1
	case "4":
		utils.ClearScreen()
		KeretaController()
	default:
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
		goto Step2
	}

}

func formKereta(Kereta model.Kereta, Update bool) model.Kereta {
	var kode int

	render := bufio.NewReader(os.Stdin)
	utils.PrintHead("Form Kereta")

	format := "Masukkan %s %s: "
	if Update {
		format = "Masukkan %s (%s): "
	}

Step1:
	fmt.Printf(format, "Kode Kereta", Kereta.Kode)
	kodeInput, err := render.ReadString('\n')
	kodeInput = strings.TrimSpace(kodeInput)

	if !Update && (err != nil || kodeInput == "") {
		utils.PrintMessage("Input tidak boleh kosong", "error")
		goto Step1
	}

	kode, err = strconv.Atoi(kodeInput)
	if !Update && (err != nil || kode <= 0) {
		utils.PrintMessage("Kode harus berupa angka positif", "error")
		goto Step1
	}

	if !Update && kodeInput != "" {
		_, uniq := utils.FindOne(model.ListKereta, model.Kereta{Kode: kode}, func(a, b model.Kereta) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})
		Kereta.Kode = kode
		if uniq {
			utils.PrintMessage("Kode Kereta sudah ada, silakan coba lagi.", "error")
			goto Step1
		}
	}

Step2:
	fmt.Printf(format, "Nama Kereta", Kereta.Nama)
	namaInput, err := render.ReadString('\n')
	namaInput = strings.TrimSpace(namaInput)
	if !Update && (err != nil || namaInput == "") {
		utils.PrintMessage("Input tidak boleh kosong", "error")
		goto Step2
	}
	if !Update && namaInput != "" {
		Kereta.Nama = namaInput
	}
Step3:
	fmt.Printf(format, "Kelas Kereta", Kereta.Kelas)
	kelasInput, err := render.ReadString('\n')
	kelasInput = strings.TrimSpace(kelasInput)
	if !Update && (err != nil || kelasInput == "") {
		utils.PrintMessage("Input tidak boleh kosong", "error")
		goto Step3
	}
	if !Update && kelasInput != "" {
		Kereta.Kelas = kelasInput
	}

	return model.Kereta{
		Kode:  Kereta.Kode,
		Nama:  Kereta.Nama,
		Kelas: Kereta.Kelas,
	}

}

func addKereta() {
	utils.PrintHead("Tambah Data Kereta")
	kereta := formKereta(model.Kereta{}, false)
	model.ListKereta = append(model.ListKereta, kereta)
	utils.ClearScreen()
	utils.PrintMessage("Kereta berhasil ditambahkan.", "success")
	KeretaController()
}

func updateKereta() {
	var kode int
	utils.PrintHead("Edit Data Kereta")
Step1:
	fmt.Print("Masukan Kode Kereta yang ingin diubah: ")
	_, err := fmt.Scan(&kode)
	if err != nil || kode <= 0 {
		utils.PrintMessage("Kode tidak valid, silakan coba lagi.", "error")
		goto Step1
	}
	data, has := utils.FindOne(model.ListKereta, model.Kereta{Kode: kode}, func(a, b model.Kereta) int {
		if a.Kode < b.Kode {
			return -1
		} else if a.Kode > b.Kode {
			return 1
		}
		return 0
	})

	if !has {
		utils.PrintMessage("Kode kereta tidak ditemukan, silakan coba lagi.", "error")
		goto Step1
	}

	kereta := formKereta(data, true)
	for i, kt := range model.ListKereta {
		if kt.Kode == kode {
			model.ListKereta[i] = kereta
			utils.ClearScreen()
			fmt.Println("Kereta berhasil diubah.")
			KeretaController()
			return
		}
	}
	utils.PrintMessage("Terjadi kesalahan saat mengubah data kereta.", "error")
	utils.ClearScreen()
	updateKereta()
	return
}

func deleteKereta() {
	var kode int
	utils.PrintHead("Hapus Data Kereta")
Step1:
	fmt.Print("Masukan Kode Kereta yang ingin dihapus: ")
	_, err := fmt.Scan(&kode)
	if err != nil || kode <= 0 {
		utils.PrintMessage("Kode tidak valid, silakan coba lagi.", "error")
		goto Step1
	}
	_, has := utils.FindOne(model.ListKereta, model.Kereta{Kode: kode}, func(a, b model.Kereta) int {
		if a.Kode < b.Kode {
			return -1
		} else if a.Kode > b.Kode {
			return 1
		}
		return 0
	})
	if !has {
		utils.PrintMessage("Kode kereta tidak ditemukan, silakan coba lagi.", "error")
		goto Step1
	}
	for i, kt := range model.ListKereta {
		if kt.Kode == kode {
			model.ListKereta = append(model.ListKereta[:i], model.ListKereta[i+1:]...)
			utils.ClearScreen()
			utils.PrintMessage("Kereta berhasil dihapus.", "success")
			KeretaController()
			return
		}
	}
	utils.PrintMessage("Terjadi kesalahan saat menghapus data", "error")

}
