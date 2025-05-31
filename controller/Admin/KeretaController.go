package Admin

import (
	"fmt"
	"laraska/model"
	"laraska/utils"
	"strconv"
	"strings"
)

func KeretaController() {
	utils.PrintHead("Menu Kereta")

	fmt.Println("[1] Tampilkan Data Kereta")
	fmt.Println("[2] Tambah Data Kereta")
	fmt.Println("[3] Edit Data Kereta")
	fmt.Println("[4] Hapus Data Kereta")
	fmt.Println("[5] Kembali Ke Menu Awal")
	utils.Divider("-")

	Input := utils.Input("Masukan nomor menu : ", func(value string) (bool, string) {
		if value == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsNumeric(value) {
			return false, "Input harus berupa angka"
		}

		if !utils.IsIn(value, []string{"1", "2", "3", "4", "5"}) {
			return false, "Pilihan tidak valid, silakan coba lagi."
		}

		return true, ""
	})

	Select, _ := strconv.Atoi(Input)

	switch Select {
	case 1:
		showKereta()
	case 2:
		addKereta()
	case 3:
		updateKereta()
	case 4:
		deleteKereta()
	case 5:
		utils.ClearScreen()
		MenuAwalAdmin()
	}

}

func showKereta() {
	utils.ClearScreen()

	var SearchOn, SortBy, SortCol string

	Stop := false

	for !Stop {
		var Data []model.Kereta = model.ListKereta

		if SortBy != "" && SortCol != "" {
			Data = utils.InsertionSort(model.ListKereta, func(a, b model.Kereta) bool {
				switch SortCol {
				case "Kode":
					if SortBy == "asc" {
						return a.Kode < b.Kode
					} else {
						return a.Kode > b.Kode
					}
				case "Nama":
					if SortBy == "asc" {
						return strings.ToLower(a.Nama) < strings.ToLower(b.Nama)
					} else {
						return strings.ToLower(a.Nama) > strings.ToLower(b.Nama)
					}
				case "Kelas":
					if SortBy == "asc" {
						return strings.ToLower(a.Kelas) < strings.ToLower(b.Kelas)
					} else {
						return strings.ToLower(a.Kelas) > strings.ToLower(b.Kelas)
					}
				}
				return false
			})
		}

		dataAny := make([]map[string]string, 0, len(Data))
		for _, dt := range Data {
			if SearchOn == "" || strings.Contains(strings.ToLower(fmt.Sprintf("%s %s %d", dt.Nama, dt.Kelas, dt.Kode)), strings.ToLower(SearchOn)) {
				dataAny = append(dataAny, map[string]string{
					"Kode":  strconv.Itoa(dt.Kode),
					"Nama":  dt.Nama,
					"Kelas": dt.Kelas,
				})
			}
		}

		utils.PrintTable(
			[]string{"Kode", "Nama", "Kelas"},
			dataAny,
			[]string{"Kode", "Nama", "Kelas"},
			4,
			"Data Kereta",
		)

		fmt.Println(utils.ColorText("[1] Pencarian", 90, 49, false))
		fmt.Println(utils.ColorText("[2] Sorting Data", 90, 49, false))
		fmt.Println(utils.ColorText("[3] Tampilkan Seluruh Data", 90, 49, false))
		fmt.Println(utils.ColorText("[0] Kembali", 90, 49, false))
		utils.Divider("-")

		Input := utils.Input("Masukan nomor menu : ", func(value string) (bool, string) {
			if value == "" {
				return false, "Input tidak boleh kosong"
			}

			if !utils.IsNumeric(value) {
				return false, "Input harus berupa angka"
			}

			if !utils.IsIn(value, []string{"0", "1", "2", "3"}) {
				return false, "Pilihan tidak valid, silakan coba lagi."
			}

			return true, ""
		})

		Select, _ := strconv.Atoi(Input)
		switch Select {
		case 1:
			SearchOn = utils.Input("Masukkan keyword pencarian: ", func(value string) (bool, string) {
				if value == "" {
					return false, "Pencarian tidak boleh kosong"
				}
				return true, ""
			})
			utils.ClearScreen()
		case 2:
			SortCol = utils.Input("Masukkan kolom yang ingin diurutkan (Kode, Nama, Kelas): ", func(value string) (bool, string) {
				if value == "" {
					return false, "Kolom tidak boleh kosong"
				}
				if !utils.IsIn(value, []string{"Kode", "Nama", "Kelas"}) {
					return false, "Kolom tidak valid, silakan coba lagi."
				}
				return true, ""
			})
			SortBy = utils.Input("Masukkan urutan (asc/desc): ", func(value string) (bool, string) {
				if value == "" {
					return false, "Urutan tidak boleh kosong"
				}
				if !utils.IsIn(strings.ToLower(value), []string{"asc", "desc"}) {
					return false, "Urutan tidak valid, silakan coba lagi."
				}
				return true, ""
			})
			utils.ClearScreen()
		case 3:
			SearchOn = ""
			SortBy = ""
			SortCol = ""
			utils.ClearScreen()
		case 0:
			utils.ClearScreen()
			Stop = true
			KeretaController()
		}
	}

}

func addKereta() {
	utils.ClearScreen()
	var Kereta model.Kereta

	utils.PrintHead("Tambah Data Kereta")
	InputKode := utils.Input("Masukan Kode Kereta : ", func(input string) (bool, string) {

		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		SortingData := utils.InsertionSort(model.ListKereta, func(a, b model.Kereta) bool {
			return a.Kode < b.Kode
		})

		Kode, _ := strconv.Atoi(input)
		_, Have, _ := utils.BinaryFindOne(SortingData, model.Kereta{Kode: Kode}, func(a, b model.Kereta) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if Have {
			return false, "Kode kereta sudah ada, silakan masukkan kode yang berbeda."
		}

		return true, ""
	})

	Kereta.Kode, _ = strconv.Atoi(InputKode)

	Kereta.Nama = utils.Input("Masukan Nama Kereta : ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		return true, ""
	})

	Kereta.Kelas = utils.Input("Masukan Kelas Kereta (Ekonomi, Bisnis, Eksekutif) : ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsIn(strings.ToLower(input), []string{"ekonomi", "bisnis", "eksekutif"}) {
			return false, "Input hanya Ekonomi, Bisnis, Eksekutif"
		}

		return true, ""
	})

	Approve := utils.Input("Apakah data kereta sudah benar? (y/n): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsIn(strings.ToLower(input), []string{"y", "n"}) {
			return false, "Input hanya y atau n"
		}
		return true, ""
	})

	utils.ClearScreen()
	Approve = strings.ToLower(Approve)
	if Approve == "n" {
		utils.PrintMessage("Tambah kereta dibatalkan.", "warning")
	} else {
		model.ListKereta = append(model.ListKereta, Kereta)
		utils.PrintMessage("Tambah Kereta berhasil ditambahkan.", "success")
	}

	KeretaController()
}

func updateKereta() {
	utils.ClearScreen()
	utils.PrintHead("Edit Data Kereta")

	var Index int
	var Data model.Kereta

	utils.Input("Masukan Kode Kereta yang ingin diubah : ", func(input string) (bool, string) {
		var Have bool
		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		SortingData := utils.InsertionSort(model.ListKereta, func(a, b model.Kereta) bool {
			return a.Kode < b.Kode
		})

		Kode, _ := strconv.Atoi(input)
		Data, Have, Index = utils.BinaryFindOne(SortingData, model.Kereta{Kode: Kode}, func(a, b model.Kereta) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if !Have {
			return false, "Kode kereta tidak ditemukan, silakan masukkan kode yang valid."
		}

		return true, ""
	})

	utils.ClearScreen()
	utils.PrintHead("Edit Data Kereta")
	utils.PrintBoxLeft(60, []string{
		fmt.Sprintf("Kode Kereta: %d", Data.Kode),
		fmt.Sprintf("Nama Kereta: %s", Data.Nama),
		fmt.Sprintf("Kelas Kereta: %s", Data.Kelas),
	})

	Data.Nama = utils.Input("Masukan Nama Kereta : ", func(input string) (bool, string) {
		return true, ""
	})

	if Data.Nama == "" {
		Data.Nama = model.ListKereta[Index].Nama
	}

	Data.Kelas = utils.Input("Masukan Kelas Kereta (Ekonomi, Bisnis, Eksekutif) : ", func(input string) (bool, string) {
		if input != "" && !utils.IsIn(strings.ToLower(input), []string{"ekonomi", "bisnis", "eksekutif"}) {
			return false, "Input hanya Ekonomi, Bisnis, Eksekutif"
		}

		return true, ""
	})

	if Data.Kelas == "" {
		Data.Kelas = model.ListKereta[Index].Kelas
	}

	Approve := utils.Input("Apakah data kereta sudah benar? (y/n): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsIn(strings.ToLower(input), []string{"y", "n"}) {
			return false, "Input hanya y atau n"
		}
		return true, ""
	})

	utils.ClearScreen()
	Approve = strings.ToLower(Approve)
	if Approve == "n" {
		utils.PrintMessage("Edit kereta dibatalkan.", "warning")
	} else {
		model.ListKereta[Index] = Data
		utils.PrintMessage("Edit Kereta berhasil diubah.", "success")
	}
	KeretaController()

}

func deleteKereta() {
	utils.ClearScreen()
	utils.PrintHead("Hapus Data Kereta")

	var Index int
	var Data model.Kereta

	utils.Input("Masukan Kode Kereta yang ingin dihapus : ", func(input string) (bool, string) {
		var Have bool
		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		SortingData := utils.InsertionSort(model.ListKereta, func(a, b model.Kereta) bool {
			return a.Kode < b.Kode
		})

		Kode, _ := strconv.Atoi(input)
		Data, Have, Index = utils.BinaryFindOne(SortingData, model.Kereta{Kode: Kode}, func(a, b model.Kereta) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if !Have {
			return false, "Kode kereta tidak ditemukan, silakan masukkan kode yang valid."
		}

		return true, ""
	})

	utils.ClearScreen()
	utils.PrintHead("Hapus Data Kereta")
	utils.PrintBoxLeft(60, []string{
		fmt.Sprintf("Kode Kereta: %d", Data.Kode),
		fmt.Sprintf("Nama Kereta: %s", Data.Nama),
		fmt.Sprintf("Kelas Kereta: %s", Data.Kelas),
	})

	Approve := utils.Input("Apakah Anda yakin ingin menghapus data ini? (y/n): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsIn(strings.ToLower(input), []string{"y", "n"}) {
			return false, "Input hanya y atau n"
		}
		return true, ""
	})

	utils.ClearScreen()
	if strings.ToLower(Approve) == "n" {
		utils.PrintMessage("Hapus kereta dibatalkan.", "warning")
	} else {
		model.ListKereta = append(model.ListKereta[:Index], model.ListKereta[Index+1:]...)
		utils.PrintMessage("Hapus Kereta berhasil dihapus.", "success")
	}
	KeretaController()
}
