package Admin

import (
	"fmt"
	"laraska/model"
	"laraska/utils"
	"strconv"
	"strings"
	"time"
)

func RuteController() {
	utils.PrintHead("Menu Rute Kereta Api")

	fmt.Println("[1] Lihat Rute")
	fmt.Println("[2] Tambah Rute")
	fmt.Println("[3] Hapus Rute")
	fmt.Println("[0] Kembali ke Menu Awal")
	utils.Divider("-")

	InputSelect := utils.Input("Pilih menu: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "3", "4", "0"}) {
			return false, "Input tidak valid, silakan pilih menu yang tersedia"
		}

		return true, ""
	})

	Select, _ := strconv.Atoi(InputSelect)

	switch Select {
	case 1:
		ShowRute()
	case 2:
		AddRute()
	case 3:
		DeleteRute()
	case 0:
		utils.ClearScreen()
		MenuAwalAdmin()
	}

}

func ShowRute() {
	utils.ClearScreen()

	Data := model.ListRute

	var SearchOn, SortBy, SortCol string
	Stop := false

	for !Stop {
		if SortBy != "" && SortCol != "" {
			Data = utils.InsertionSort(model.ListRute, func(a, b model.Rute) bool {
				switch SortCol {
				case "Kode":
					if SortBy == "asc" {
						return a.Kode < b.Kode
					}
					return a.Kode > b.Kode
				case "Nama":
					if SortBy == "asc" {
						return a.Nama < b.Nama
					}
					return a.Nama > b.Nama
				case "StasiunAsal":
					if SortBy == "asc" {
						return a.StasiunAwal.Nama < b.StasiunAwal.Nama
					}
					return a.StasiunAwal.Nama > b.StasiunAwal.Nama
				case "StasiunTujuan":
					if SortBy == "asc" {
						return a.StasiunAkhir.Nama < b.StasiunAkhir.Nama
					}
					return a.StasiunAkhir.Nama > b.StasiunAkhir.Nama
				case "Harga":
					if SortBy == "asc" {
						return a.Harga < b.Harga
					}
					return a.Harga > b.Harga
				case "Gerbong":
					if SortBy == "asc" {
						return a.Gerbong < b.Gerbong
					}
					return a.Gerbong > b.Gerbong
				case "Kereta":
					if SortBy == "asc" {
						return a.Kereta.Nama < b.Kereta.Nama
					}
					return a.Kereta.Nama > b.Kereta.Nama
				case "RuteBehenti":
					if SortBy == "asc" {
						return len(a.RuteBerhenti) < len(b.RuteBerhenti)
					}
					return len(a.RuteBerhenti) > len(b.RuteBerhenti)
				case "Berangkat":
					if SortBy == "asc" {
						return a.RuteBerhenti[0].Berangkat.Before(b.RuteBerhenti[0].Berangkat)
					}
					return a.RuteBerhenti[0].Berangkat.After(b.RuteBerhenti[0].Berangkat)
				case "Tiba":
					if SortBy == "asc" {
						return a.RuteBerhenti[len(a.RuteBerhenti)-1].Tiba.Before(b.RuteBerhenti[len(b.RuteBerhenti)-1].Tiba)
					}
					return a.RuteBerhenti[len(a.RuteBerhenti)-1].Tiba.After(b.RuteBerhenti[len(b.RuteBerhenti)-1].Tiba)
				}
				return false
			})
		}

		dataAny := make([]map[string]string, 0, len(Data))
		SearchOn = strings.ToLower(SearchOn)
		for _, rute := range Data {
			DataSearch := strings.ToLower(fmt.Sprintf("%s %s %s %s %d %d %s %d %s %s",
				rute.Kode,
				rute.Nama,
				rute.StasiunAwal.Nama,
				rute.StasiunAkhir.Nama,
				rute.Harga,
				rute.Gerbong,
				rute.Kereta.Nama,
				len(rute.RuteBerhenti),
				rute.RuteBerhenti[0].Berangkat.Format("2006-01-02 15:04:05"),
				rute.RuteBerhenti[len(rute.RuteBerhenti)-1].Tiba.Format("2006-01-02 15:04:05"),
			))

			if SearchOn == "" || strings.Contains(DataSearch, SearchOn) {
				dataAny = append(dataAny, map[string]string{
					"Kode":          rute.Kode,
					"Nama":          rute.Nama,
					"StasiunAsal":   rute.StasiunAwal.Nama,
					"StasiunTujuan": rute.StasiunAkhir.Nama,
					"Harga":         utils.RupiahFormat(rute.Harga),
					"Gerbong":       strconv.Itoa(rute.Gerbong),
					"Kereta":        rute.Kereta.Nama,
					"RuteBehenti":   strconv.Itoa(len(rute.RuteBerhenti)),
					"Berangkat":     rute.RuteBerhenti[0].Berangkat.Format("2006-01-02 15:04:05"),
					"Tiba":          rute.RuteBerhenti[len(rute.RuteBerhenti)-1].Tiba.Format("2006-01-02 15:04:05"),
				})
			}
		}

		utils.PrintTable(
			[]string{"Kode", "Nama", "Stasiun Asal", "Stasiun Tujuan", "Harga", "Gerbong", "Kereta", "Titik Rute Berhenti", "Berangkat", "Tiba"},
			dataAny,
			[]string{"Kode", "Nama", "StasiunAsal", "StasiunTujuan", "Harga", "Gerbong", "Kereta", "RuteBehenti", "Berangkat", "Tiba"},
			4,
			"Table Rute Kereta Api",
		)

		fmt.Println(utils.ColorText("[1] Pencarian", 90, 49, false))
		fmt.Println(utils.ColorText("[2] Sorting Data", 90, 49, false))
		fmt.Println(utils.ColorText("[3] Tampilkan Seluruh Data", 90, 49, false))
		fmt.Println(utils.ColorText("[4] Detail Rute", 90, 49, false))
		fmt.Println(utils.ColorText("[0] Kembali", 90, 49, false))
		utils.Divider("-")

		Input := utils.Input("Masukan nomor menu : ", func(value string) (bool, string) {
			if value == "" {
				return false, "Input tidak boleh kosong"
			}

			if !utils.IsNumeric(value) {
				return false, "Input harus berupa angka"
			}

			if !utils.IsIn(value, []string{"0", "1", "2", "3", "4"}) {
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
			SortCol = utils.Input("Masukkan kolom yang ingin diurutkan (Kode/Nama/Stasiun Asal/Stasiun Tujuan/Harga/Gerbong/Kereta/Titik Rute Berhenti/Berangkat/Tiba): ", func(value string) (bool, string) {
				if value == "" {
					return false, "Kolom tidak boleh kosong"
				}
				if !utils.IsIn(value, []string{"Kode", "Nama", "StasiunAsal", "StasiunTujuan", "Harga", "Gerbong", "Kereta", "RuteBehenti", "Berangkat", "Tiba"}) {
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
		case 4:
			DetailRute()
		case 0:
			utils.ClearScreen()
			Stop = true
			RuteController()

		}

	}
}

func DetailRute() {
	var Data model.Rute
	var Have bool

	utils.Input("Masukkan Kode Rute: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Kode Rute tidak boleh kosong"
		}

		SortData := utils.InsertionSort(model.ListRute, func(a, b model.Rute) bool {
			return a.Kode < b.Kode
		})

		Data, Have, _ = utils.BinaryFindOne(SortData, model.Rute{Kode: value}, func(a model.Rute, b model.Rute) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if !Have {
			return false, fmt.Sprintf("Rute dengan Kode %s tidak ditemukan", value)
		}

		return true, ""
	})

	utils.ClearScreen()
	utils.PrintHead("Detail Rute Kereta Api")

	utils.PrintBoxLeft(60, []string{
		fmt.Sprintf("Kode Rute: %s", Data.Kode),
		fmt.Sprintf("Nama Rute: %s", Data.Nama),
		fmt.Sprintf("Stasiun Asal: %s", Data.StasiunAwal.Nama),
		fmt.Sprintf("Stasiun Tujuan: %s", Data.StasiunAkhir.Nama),
		fmt.Sprintf("Harga: %s", utils.RupiahFormat(Data.Harga)),
		fmt.Sprintf("Gerbong: %d", Data.Gerbong),
		fmt.Sprintf("Kereta: %s", Data.Kereta.Nama),
		fmt.Sprintf("Titik Rute Berhenti: %d", len(Data.RuteBerhenti)),
		fmt.Sprintf("Berangkat: %s", Data.RuteBerhenti[0].Berangkat.Format("2006-01-02 15:04:05")),
		fmt.Sprintf("Tiba: %s", Data.RuteBerhenti[len(Data.RuteBerhenti)-1].Tiba.Format("2006-01-02 15:04:05")),
	})

	RuteMap := make([]map[string]string, 0, len(Data.RuteBerhenti))

	for _, rute := range Data.RuteBerhenti {
		durasi := rute.Tiba.Sub(rute.Berangkat)
		RuteMap = append(RuteMap, map[string]string{
			"StasiunAwal":  rute.StasiunAwal.Nama,
			"StasiunAkhir": rute.StasiunAkhir.Nama,
			"Berangkat":    rute.Berangkat.Format("2006-01-02 15:04:05"),
			"Tiba":         rute.Tiba.Format("2006-01-02 15:04:05"),
			"Durasi":       durasi.String(),
		})
	}

	utils.PrintTable(
		[]string{"Stasiun Awal", "Stasiun Akhir", "Berangkat", "Tiba", "Durasi"},
		RuteMap,
		[]string{"StasiunAwal", "StasiunAkhir", "Berangkat", "Tiba", "Durasi"},
		4,
		"Rute Berhenti",
	)

	fmt.Println(utils.ColorText("[0] Kembali", 90, 49, false))

	utils.Divider("-")
	Input := utils.Input("Masukkan nomor menu: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsNumeric(value) {
			return false, "Input harus berupa angka"
		}

		if value != "0" {
			return false, "Pilihan tidak valid, silakan coba lagi."
		}

		return true, ""
	})

	if Input == "0" {
		utils.ClearScreen()
		ShowRute()
	}

}

func AddRute() {
	var Rute model.Rute
	var JumlahRuteBerhenti int

	utils.ClearScreen()
	utils.PrintHead("Tambah Rute Kereta Api")

	Rute.Kode = utils.Input("Masukkan Kode Rute: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Kode Rute tidak boleh kosong"
		}

		if len(value) < 3 || len(value) > 10 {
			return false, "Kode Rute harus antara 3 hingga 10 karakter"
		}

		Sorting := utils.InsertionSort(model.ListRute, func(a, b model.Rute) bool {
			return a.Kode < b.Kode
		})

		_, Have, _ := utils.BinaryFindOne(Sorting, model.Rute{Kode: value}, func(a model.Rute, b model.Rute) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if Have {
			return false, fmt.Sprintf("Rute dengan Kode %s sudah ada", value)
		}

		return true, ""
	})

	Rute.Nama = utils.Input("Masukkan Nama Rute: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Nama Rute tidak boleh kosong"
		}

		return true, ""
	})

	utils.Input("Masukkan Harga Rute: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Harga Rute tidak boleh kosong"
		}

		if !utils.IsNumeric(value) {
			return false, "Harga Rute harus berupa angka"
		}

		harga, _ := strconv.Atoi(value)
		if harga <= 0 {
			return false, "Harga Rute harus lebih besar dari 0"
		}

		Rute.Harga = harga
		return true, ""
	})

	utils.Input("Masukkan Jumlah Gerbong: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Jumlah Gerbong tidak boleh kosong"
		}

		if !utils.IsNumeric(value) {
			return false, "Jumlah Gerbong harus berupa angka"
		}

		gerbong, _ := strconv.Atoi(value)
		if gerbong <= 0 {
			return false, "Jumlah Gerbong harus lebih besar dari 0"
		}

		Rute.Gerbong = gerbong
		return true, ""
	})

	TableKereta("Pilih Kereta untuk Rute ini")

	utils.Input("Masukkan Kode Kereta: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Kode Kereta tidak boleh kosong"
		}

		Sorting := utils.InsertionSort(model.ListKereta, func(a, b model.Kereta) bool {
			return a.Kode < b.Kode
		})

		Kode, err := strconv.Atoi(value)
		if err != nil {
			return false, "Kode Kereta harus berupa angka"
		}
		Data, Have, _ := utils.BinaryFindOne(Sorting, model.Kereta{Kode: Kode}, func(a model.Kereta, b model.Kereta) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if !Have {
			return false, fmt.Sprintf("Kereta dengan Kode %s tidak ditemukan", value)
		}

		Rute.Kereta = Data
		return true, ""
	})

	utils.Input("Masukkan Jumlah Rute Berhenti: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Jumlah Rute Berhenti tidak boleh kosong"
		}

		if !utils.IsNumeric(value) {
			return false, "Jumlah Rute Berhenti harus berupa angka"
		}

		jumlah, _ := strconv.Atoi(value)
		if jumlah <= 0 {
			return false, "Jumlah Rute Berhenti harus lebih besar dari 0"
		}

		JumlahRuteBerhenti = jumlah
		return true, ""
	})

	TableStasiun("Data Stasiun")

	Rute.RuteBerhenti = make([]model.RuteBerhenti, JumlahRuteBerhenti)

	for i := 0; i < JumlahRuteBerhenti; i++ {
		if i > 0 {
			utils.Divider("-")
		}
		fmt.Printf("Rute Berhenti %d:\n", i+1)

		RuteBerhenti := model.RuteBerhenti{}

		if i == 0 {
			utils.Input("Masukkan IDStasiun Stasiun Awal: ", func(value string) (bool, string) {
				if value == "" {
					return false, "IDStasiun Stasiun Awal tidak boleh kosong"
				}

				Sorting := utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool {
					return a.IDStasiun < b.IDStasiun
				})

				Data, Have, _ := utils.BinaryFindOne(Sorting, model.Stasiun{IDStasiun: value}, func(a model.Stasiun, b model.Stasiun) int {
					if a.IDStasiun < b.IDStasiun {
						return -1
					} else if a.IDStasiun > b.IDStasiun {
						return 1
					}
					return 0
				})

				if !Have {
					return false, fmt.Sprintf("Stasiun dengan IDStasiun %s tidak ditemukan", value)
				}

				RuteBerhenti.StasiunAwal = Data
				return true, ""
			})
		} else {
			RuteBerhenti.StasiunAwal = Rute.RuteBerhenti[i-1].StasiunAkhir
			fmt.Printf("Masukan IDStasiun Stasiun Awal: %s\n", RuteBerhenti.StasiunAwal.IDStasiun)
		}

		utils.Input("Masukkan IDStasiun Stasiun Akhir: ", func(value string) (bool, string) {
			if value == "" {
				return false, "IDStasiun Stasiun Akhir tidak boleh kosong"
			}

			Sorting := utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool {
				return a.IDStasiun < b.IDStasiun
			})

			Data, Have, _ := utils.BinaryFindOne(Sorting, model.Stasiun{IDStasiun: value}, func(a model.Stasiun, b model.Stasiun) int {
				if a.IDStasiun < b.IDStasiun {
					return -1
				} else if a.IDStasiun > b.IDStasiun {
					return 1
				}
				return 0
			})

			if !Have {
				return false, fmt.Sprintf("Stasiun dengan IDStasiun %s tidak ditemukan", value)
			}

			RuteBerhenti.StasiunAkhir = Data
			return true, ""
		})

		utils.Input("Masukkan Waktu Berangkat (YYYY-MM-DD HH:MM:SS): ", func(value string) (bool, string) {
			if value == "" {
				return false, "Waktu Berangkat tidak boleh kosong"
			}
			waktuBerangkat, err := time.Parse("2006-01-02 15:04:05", value)
			if err != nil {
				return false, "Format waktu berangkat tidak valid, gunakan format YYYY-MM-DD HH:MM:SS"
			}
			if i > 0 {
				prevTiba := Rute.RuteBerhenti[i-1].Tiba
				if !waktuBerangkat.After(prevTiba) {
					return false, fmt.Sprintf("Waktu Berangkat harus setelah Waktu Tiba pada rute berhenti sebelumnya (%s)", prevTiba.Format("2006-01-02 15:04:05"))
				}
			}
			RuteBerhenti.Berangkat = waktuBerangkat
			return true, ""
		})

		utils.Input("Masukkan Waktu Tiba (YYYY-MM-DD HH:MM:SS): ", func(value string) (bool, string) {
			if value == "" {
				return false, "Waktu Tiba tidak boleh kosong"
			}

			waktuTiba, err := time.Parse("2006-01-02 15:04:05", value)
			if err != nil {
				return false, "Format waktu tiba tidak valid, gunakan format YYYY-MM-DD HH:MM:SS"
			}

			if !waktuTiba.After(RuteBerhenti.Berangkat) {
				return false, "Waktu Tiba harus setelah Waktu Berangkat"
			}

			RuteBerhenti.Tiba = waktuTiba
			return true, ""
		})

		Rute.RuteBerhenti[i] = RuteBerhenti
	}

	Rute.StasiunAwal = Rute.RuteBerhenti[0].StasiunAwal
	Rute.StasiunAkhir = Rute.RuteBerhenti[len(Rute.RuteBerhenti)-1].StasiunAkhir

	Next := utils.Input("Apakah Anda yakin ingin menambahkan Rute ini? (y/n): ", func(value string) (bool, string) {
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
		model.ListRute = append(model.ListRute, Rute)
		utils.PrintMessage("Rute berhasil ditambahkan", "success")
	} else {
		utils.PrintMessage("Rute tidak ditambahkan", "warning")
	}
	RuteController()
}

func DeleteRute() {
	var KodeRute string
	var Index int
	var Have bool

	utils.ClearScreen()
	utils.PrintHead("Delete Rute Kereta Api")

	KodeRute = utils.Input("Masukkan Kode Rute yang ingin dihapus: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Kode Rute tidak boleh kosong"
		}

		Sorting := utils.InsertionSort(model.ListRute, func(a, b model.Rute) bool {
			return a.Kode < b.Kode
		})

		_, Have, Index = utils.BinaryFindOne(Sorting, model.Rute{Kode: value}, func(a model.Rute, b model.Rute) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if !Have {
			return false, fmt.Sprintf("Rute dengan Kode %s tidak ditemukan", value)
		}

		return true, ""
	})

	Confirm := utils.Input("Apakah Anda yakin ingin menghapus Rute ini? (y/n): ", func(value string) (bool, string) {
		if value == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsIn(value, []string{"y", "n"}) {
			return false, "Input tidak valid, silakan coba lagi."
		}
		return true, ""
	})

	utils.ClearScreen()
	if Confirm == "y" {
		model.ListRute = append(model.ListRute[:Index], model.ListRute[Index+1:]...)
		utils.PrintMessage(fmt.Sprintf("Rute dengan Kode %s berhasil dihapus", KodeRute), "success")
	} else {
		utils.PrintMessage("Rute tidak dihapus", "warning")
	}

	RuteController()
}

func TableRuteWithData(data []model.Rute, StasiunAwal model.Stasiun, StasiunAkhir model.Stasiun, Title string) {
	DataMap := make([]map[string]string, 0, len(data))
	for _, rute := range data {
		DataMap = append(DataMap, map[string]string{
			"Kode":          rute.Kode,
			"Nama":          rute.Nama,
			"StasiunAsal":   StasiunAwal.Nama,
			"StasiunTujuan": StasiunAkhir.Nama,
			"Harga":         utils.RupiahFormat(rute.Harga),
			"Gerbong":       strconv.Itoa(rute.Gerbong),
			"Kereta":        rute.Kereta.Nama,
			"RuteBehenti":   strconv.Itoa(len(rute.RuteBerhenti)),
			"Berangkat":     rute.RuteBerhenti[0].Berangkat.Format("2006-01-02 15:04:05"),
			"Tiba":          rute.RuteBerhenti[len(rute.RuteBerhenti)-1].Tiba.Format("2006-01-02 15:04:05"),
		})
	}

	utils.PrintTable(
		[]string{"Kode", "Nama", "Stasiun Asal", "Stasiun Tujuan", "Harga", "Gerbong", "Kereta", "Titik Rute Berhenti", "Berangkat", "Tiba"},
		DataMap,
		[]string{"Kode", "Nama", "StasiunAsal", "StasiunTujuan", "Harga", "Gerbong", "Kereta", "RuteBehenti", "Berangkat", "Tiba"},
		4,
		Title,
	)
}
