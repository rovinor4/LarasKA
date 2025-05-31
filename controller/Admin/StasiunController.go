package Admin

import (
	"fmt"
	"laraska/model"
	"laraska/utils"
	"strconv"
	"strings"
)

func StasiunController() {
	utils.PrintHead("Menu Stasiun")

	fmt.Println("[1] Lihat Stasiun")
	fmt.Println("[2] Tambah Stasiun")
	fmt.Println("[3] Edit Stasiun")
	fmt.Println("[4] Hapus Stasiun")
	fmt.Println("[5] Kembali ke Menu Awal")
	utils.Divider("-")

	InputSelect := utils.Input("Pilih menu: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "3", "4", "5"}) {
			return false, "Input tidak valid, silakan pilih menu yang tersedia"
		}

		return true, ""
	})

	Select, _ := strconv.Atoi(InputSelect)

	switch Select {
	case 1:
		ShowStasiun()
	case 2:
		AddStasiun()
	case 3:
		EditStasiun()
	}
}

func ShowStasiun() {
	utils.ClearScreen()

	var SearchOn, SortBy, SortCol string
	Stop := false

	for !Stop {
		var Data []model.Stasiun = model.ListStasiun

		if SortBy != "" && SortCol != "" {
			Data = utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool {
				switch SortCol {
				case "IDStasiun":
					if SortBy == "asc" {
						return a.IDStasiun < b.IDStasiun
					} else {
						return a.IDStasiun > b.IDStasiun
					}
				case "Nama":
					if SortBy == "asc" {
						return strings.ToLower(a.Nama) < strings.ToLower(b.Nama)
					} else {
						return strings.ToLower(a.Nama) > strings.ToLower(b.Nama)
					}
				case "Kota":
					if SortBy == "asc" {
						return strings.ToLower(a.Kota) < strings.ToLower(b.Kota)
					} else {
						return strings.ToLower(a.Kota) > strings.ToLower(b.Kota)
					}
				}
				return false
			})
		}

		dataAny := make([]map[string]string, 0, len(Data))
		for _, dt := range Data {
			if SearchOn == "" || strings.Contains(strings.ToLower(fmt.Sprintf("%s %s %s", dt.IDStasiun, dt.Nama, dt.Kota)), strings.ToLower(SearchOn)) {
				dataAny = append(dataAny, map[string]string{
					"IDStasiun": dt.IDStasiun,
					"Nama":      dt.Nama,
					"Kota":      dt.Kota,
				})
			}
		}

		utils.PrintTable(
			[]string{"ID Stasiun", "Nama Stasiun", "Kota"},
			dataAny,
			[]string{"IDStasiun", "Nama", "Kota"},
			4,
			"Data Stasiun",
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

		utils.Divider("-")
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
			SortCol = utils.Input("Masukkan kolom yang ingin diurutkan (IDStasiun, Nama, Kota): ", func(value string) (bool, string) {
				if value == "" {
					return false, "Kolom tidak boleh kosong"
				}
				if !utils.IsIn(value, []string{"IDStasiun", "Nama", "Kota"}) {
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
			StasiunController()

		}
	}

}

func AddStasiun() {
	var Stasiun model.Stasiun
	utils.ClearScreen()
	utils.PrintHead("Tambah Stasiun")

	Stasiun.IDStasiun = utils.Input("Masukkan ID Stasiun: ", func(value string) (bool, string) {
		if value == "" {
			return false, "ID Stasiun tidak boleh kosong"
		}
		if len(value) > 3 {
			return false, "ID Stasiun maksimal 3 karakter"
		}

		_, Have, _ := utils.FindOne(model.ListStasiun, model.Stasiun{IDStasiun: value}, func(a, b model.Stasiun) int {
			if a.IDStasiun < b.IDStasiun {
				return -1
			} else if a.IDStasiun > b.IDStasiun {
				return 1
			}
			return 0
		})

		if Have {
			return false, "ID Stasiun sudah ada, silakan gunakan ID yang lain"
		}

		return true, ""
	})

	Stasiun.Nama = utils.Input("Masukkan Nama Stasiun: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Nama Stasiun tidak boleh kosong"
		}
		return true, ""
	})

	Stasiun.Kota = utils.Input("Masukkan Kota Stasiun: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Kota Stasiun tidak boleh kosong"
		}
		return true, ""
	})

	Next := utils.Input("Apakah Anda yakin ingin menambahkan stasiun ini? (y/n): ", func(value string) (bool, string) {
		if value == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsIn(value, []string{"y", "n"}) {
			return false, "Input tidak valid, silakan coba lagi."
		}
		return true, ""
	})

	utils.ClearScreen()
	if Next == "y" {
		model.ListStasiun = append(model.ListStasiun, Stasiun)
		utils.PrintMessage("Stasiun berhasil ditambahkan", "success")
	} else {
		utils.PrintMessage("Stasiun tidak ditambahkan", "warning")
	}

	StasiunController()
}

func EditStasiun() {
	utils.ClearScreen()
	utils.PrintHead("Edit Stasiun")

	var Index int
	var Stasiun model.Stasiun

	utils.Input("Masukkan ID Stasiun yang ingin diedit: ", func(value string) (bool, string) {
		if value == "" {
			return false, "ID Stasiun tidak boleh kosong"
		}
		if len(value) > 3 {
			return false, "ID Stasiun maksimal 3 karakter"
		}

		Have := false
		Stasiun, Have, Index = utils.FindOne(model.ListStasiun, model.Stasiun{IDStasiun: value}, func(a, b model.Stasiun) int {
			if a.IDStasiun < b.IDStasiun {
				return -1
			} else if a.IDStasiun > b.IDStasiun {
				return 1
			}
			return 0
		})

		if !Have {
			return false, "ID Stasiun tidak ditemukan, silakan coba lagi."
		}
	
		return true, ""
	})

	utils.ClearScreen()
	utils.PrintHead("Edit Stasiun")
	utils.PrintBoxLeft(60, []string{
		fmt.Sprintf("ID Stasiun: %s", Stasiun.IDStasiun),
		fmt.Sprintf("Nama Stasiun: %s", Stasiun.Nama),
		fmt.Sprintf("Kota Stasiun: %s", Stasiun.Kota),
	})

	NamaBaru := utils.Input("Masukkan Nama Stasiun baru (kosongkan untuk tidak mengubah): ", func(value string) (bool, string) {
		if value == "" {
			return true, ""
		}
		return true, ""
	})

	KotaBaru := utils.Input("Masukkan Kota Stasiun baru (kosongkan untuk tidak mengubah): ", func(value string) (bool, string) {
		if value == "" {
			return true, ""
		}
		return true, ""
	})

	if NamaBaru != "" {
		Stasiun.Nama = NamaBaru
	}

	if KotaBaru != "" {
		Stasiun.Kota = KotaBaru
	}

	Next := utils.Input("Apakah Anda yakin ingin mengedit stasiun ini? (y/n): ", func(value string) (bool, string) {
		if value == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsIn(value, []string{"y", "n"}) {
			return false, "Input tidak valid, silakan coba lagi."
		}
		return true, ""
	})

	if Next == "y" {
		model.ListStasiun[Index] = Stasiun
		utils.PrintMessage("Stasiun berhasil diedit", "success")
	} else {
		utils.PrintMessage("Stasiun tidak diedit", "warning")
	}
	utils.ClearScreen()
	StasiunController()
}
