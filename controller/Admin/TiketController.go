package Admin

import (
	"fmt"
	"laraska/model"
	"laraska/utils"
	"strconv"
	"strings"
	"time"
)

func TiketController() {
	var pilihan int

	utils.PrintHead("Menu Tiket Kereta Api")

	fmt.Println("[1] Tambah Tiket")
	fmt.Println("[2] Cek Tiket")
	fmt.Println("[3] Hapus Tiket")
	fmt.Println("[0] Kembali ke Menu Utama")

	utils.Divider("-")
	InputPilihan := utils.Input("Masukkan Pilihan: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Pilihan tidak boleh kosong"
		}

		if !utils.IsNumeric(input) {
			return false, "Pilihan harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "3", "0"}) {
			return false, "Pilihan tidak valid"
		}

		return true, ""
	})

	pilihan, _ = strconv.Atoi(InputPilihan)

	switch pilihan {
	case 1:

	}
}

func TiketStasiun(Label string, StasiunSebelumnya model.Stasiun) model.Stasiun {
	utils.ClearScreen()
	SearchOn := ""
	Stop := false
	var Stasiun model.Stasiun

	for !Stop {
		utils.ClearScreen()
		var Data []model.Stasiun
		for _, stasiun := range model.ListStasiun {
			if SearchOn == "" || strings.Contains(strings.ToLower(stasiun.Nama), strings.ToLower(SearchOn)) ||
				strings.Contains(strings.ToLower(stasiun.Kota), strings.ToLower(SearchOn)) {
				Data = append(Data, stasiun)
			}
		}

		TableStasiunWithData(Data, fmt.Sprintf("%s Stasiun", Label))

		fmt.Println(utils.ColorText("[1] Pencarian", 90, 49, false))
		fmt.Println(utils.ColorText("[2] Pilih Kereta", 90, 49, false))

		utils.Divider("-")

		InputSelect := utils.Input("Pilih menu: ", func(input string) (bool, string) {
			if input == "" {
				return false, "Input tidak boleh kosong"
			}
			if !utils.IsNumeric(input) {
				return false, "Input harus berupa angka"
			}

			if !utils.IsIn(input, []string{"1", "2"}) {
				return false, "Input tidak valid, silakan pilih menu yang tersedia"
			}

			return true, ""
		})

		Select, _ := strconv.Atoi(InputSelect)

		switch Select {
		case 1:
			SearchOn = utils.Input("Masukkan nama stasiun atau kota: ", func(input string) (bool, string) {
				if input == "" {
					return false, "Input tidak boleh kosong"
				}
				return true, ""
			})
		case 2:
			var Have bool
			utils.Input("Masukkan ID Stasiun yang ingin dipilih: ", func(input string) (bool, string) {
				if input == "" {
					return false, "Input tidak boleh kosong"
				}

				Sorting := utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool {
					return a.IDStasiun < b.IDStasiun
				})

				if StasiunSebelumnya.IDStasiun != "" && StasiunSebelumnya.IDStasiun == input {
					return false, "Stasiun awal dan tujuan tidak boleh sama"
				}

				Stasiun, Have, _ = utils.BinaryFindOne(Sorting, model.Stasiun{IDStasiun: input}, func(a, b model.Stasiun) int {
					if a.IDStasiun < b.IDStasiun {
						return -1
					} else if a.IDStasiun > b.IDStasiun {
						return 1
					}
					return 0
				})

				if !Have {
					return false, "Stasiun tidak ditemukan, silakan coba lagi"
				}

				if Have {
					Stop = true
				}

				return true, ""
			})
		}

	}

	utils.ClearScreen()
	return Stasiun

}

func AddTiket() {
	StasiunAwal := TiketStasiun("Stasiun Awal", model.Stasiun{})
	StasiunTujuan := TiketStasiun("Stasiun Tujuan", StasiunAwal)

	utils.PrintHead("Tambah Tiket")
	fmt.Println("Stasiun Awal:", StasiunAwal.Nama, "-", StasiunAwal.Kota)
	fmt.Println("Stasiun Tujuan:", StasiunTujuan.Nama, "-", StasiunTujuan.Kota)

	TanggalKeberangkatan := utils.Input("Masukkan Tanggal Keberangkatan (DD-MM-YYYY): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Tanggal keberangkatan tidak boleh kosong"
		}

		if _, err := time.Parse("02-01-2006", input); err != nil {
			return false, "Format tanggal tidak valid, gunakan DD-MM-YYYY"
		}

		return true, ""
	})

	InputJumlahPenumpang := utils.Input("Masukkan Jumlah Penumpang: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Jumlah penumpang tidak boleh kosong"
		}

		if !utils.IsNumeric(input) {
			return false, "Jumlah penumpang harus berupa angka"
		}

		jumlah, _ := strconv.Atoi(input)
		if jumlah <= 0 {
			return false, "Jumlah penumpang harus lebih dari 0"
		}

		return true, ""
	})

	JumlahPenumpang, _ := strconv.Atoi(InputJumlahPenumpang)

	fmt.Println(JumlahPenumpang, TanggalKeberangkatan)

}
