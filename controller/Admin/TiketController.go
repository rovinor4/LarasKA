package Admin

import (
	"fmt"
	"laraska/model"
	"laraska/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func TiketController() {
	var pilihan int

	utils.PrintHead("Menu Tiket Kereta Api")

	fmt.Println("[1] Tambah Tiket")
	fmt.Println("[2] Cek Tiket")
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
		AddTiket()
	case 2:

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

func TiketRute(Rute []model.Rute, StasiunAwal model.Stasiun, StasiunTujuan model.Stasiun) model.Rute {
	var PilihRute model.Rute
	Stop := false
	var SortBy string

	for !Stop {

		utils.ClearScreen()
		var Data []model.Rute = Rute

		if SortBy != "" {
			Data = utils.InsertionSort(Rute, func(a, b model.Rute) bool {
				switch SortBy {
				case "1": // sort by harga asc
					return a.Harga < b.Harga
				case "2": // sort by harga desc
					return a.Harga > b.Harga
				case "3": // sort by waktu keberangkatan asc
					return a.RuteBerhenti[0].Berangkat.Before(b.RuteBerhenti[0].Berangkat)
				case "4": // sort by waktu keberangkatan desc
					return a.RuteBerhenti[0].Berangkat.After(b.RuteBerhenti[0].Berangkat)
				}

				return false
			})
		}

		TableRuteWithData(Data, StasiunAwal, StasiunTujuan, fmt.Sprintf("Rute dari %s ke %s", StasiunAwal.Nama, StasiunTujuan.Nama))

		fmt.Println(utils.ColorText("[1] Urutkan", 90, 49, false))
		fmt.Println(utils.ColorText("[2] Pilih Rute", 90, 49, false))

		InputPilihan := utils.Input("Masukkan Pilihan: ", func(input string) (bool, string) {
			if input == "" {
				return false, "Pilihan tidak boleh kosong"
			}
			if !utils.IsNumeric(input) {
				return false, "Pilihan harus berupa angka"
			}
			if !utils.IsIn(input, []string{"1", "2"}) {
				return false, "Pilihan tidak valid"
			}
			return true, ""
		})
		Pilihan, _ := strconv.Atoi(InputPilihan)

		switch Pilihan {
		case 1:
			utils.Divider("-")
			fmt.Println(utils.ColorText("[1] Sort Harga (Rendah ke Tinggi)", 90, 49, false))
			fmt.Println(utils.ColorText("[2] Sort Harga (Tinggi ke Rendah)", 90, 49, false))
			fmt.Println(utils.ColorText("[3] Sort Waktu Keberangkatan (Paling Awal)", 90, 49, false))
			fmt.Println(utils.ColorText("[4] Sort Waktu Keberangkatan (Paling Akhir)", 90, 49, false))

			SortBy = utils.Input("Masukkan Pilihan: ", func(input string) (bool, string) {
				if input == "" {
					return false, "Pilihan tidak boleh kosong"
				}
				if !utils.IsNumeric(input) {
					return false, "Pilihan harus berupa angka"
				}
				if !utils.IsIn(input, []string{"1", "2", "3", "4"}) {
					return false, "Pilihan tidak valid"
				}
				return true, ""
			})
		case 2:
			utils.Input("Masukan ID Rute yang ingin dipilih: ", func(input string) (bool, string) {

				if input == "" {
					return false, "ID Rute tidak boleh kosong"
				}

				sortedRute := utils.InsertionSort(Rute, func(a, b model.Rute) bool {
					return a.Kode < b.Kode
				})

				var found bool
				PilihRute, found, _ = utils.BinaryFindOne(sortedRute, model.Rute{Kode: input}, func(a, b model.Rute) int {
					if a.Kode < b.Kode {
						return -1
					} else if a.Kode > b.Kode {
						return 1
					}
					return 0
				})

				if !found {
					return false, "Rute tidak ditemukan, silakan coba lagi"
				}

				Stop = true
				return true, ""
			})
		}

	}

	return PilihRute
}

func AddTiket() {
	var PilihRute model.Rute
	StasiunAwal := TiketStasiun("Stasiun Awal", model.Stasiun{})
	StasiunTujuan := TiketStasiun("Stasiun Tujuan", StasiunAwal)

	utils.ClearScreen()
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

	InputJumlahPenumpang := utils.Input("Masukkan Jumlah Penumpang : ", func(input string) (bool, string) {
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

	var Rute []model.Rute

	TanggalKeberangkatanParse, _ := time.Parse("02-01-2006", TanggalKeberangkatan)

	for _, rute := range model.ListRute {

		RuteAwal := false
		RuteTujuan := false
		JadwalRute := false

		for _, Berhenti := range rute.RuteBerhenti {
			if !RuteAwal && Berhenti.StasiunAwal.IDStasiun == StasiunAwal.IDStasiun {
				RuteAwal = true
			}

			if !RuteTujuan && Berhenti.StasiunAkhir.IDStasiun == StasiunTujuan.IDStasiun {
				RuteTujuan = true
			}

			if !JadwalRute && Berhenti.Berangkat.After(TanggalKeberangkatanParse) {
				JadwalRute = true
			}
		}

		if RuteAwal && RuteTujuan && JadwalRute {
			Rute = append(Rute, rute)
		}
	}

	if len(Rute) == 0 {
		utils.PrintMessage("Rute tidak ditemukan antara stasiun awal dan tujuan", "error")
		AddTiket()
		return
	}

	PilihRute = TiketRute(Rute, StasiunAwal, StasiunTujuan)

	utils.ClearScreen()
	utils.PrintHead("Tambah Tiket | Konfirmasi")
	Berangkat := ""
	Tiba := ""
	for _, ruteBerhenti := range PilihRute.RuteBerhenti {
		if ruteBerhenti.StasiunAwal.IDStasiun == StasiunAwal.IDStasiun {
			Berangkat = ruteBerhenti.Berangkat.Format("02-01-2006 15:04")
		}

		if ruteBerhenti.StasiunAkhir.IDStasiun == StasiunTujuan.IDStasiun {
			Tiba = ruteBerhenti.Tiba.Format("02-01-2006 15:04")
		}
	}

	utils.PrintBoxLeft(60, []string{
		fmt.Sprintf("Stasiun Awal: %s - %s", StasiunAwal.Nama, StasiunAwal.Kota),
		fmt.Sprintf("Stasiun Tujuan: %s - %s", StasiunTujuan.Nama, StasiunTujuan.Kota),
		fmt.Sprintf("Waktu Berangkat: %s", Berangkat),
		fmt.Sprintf("Waktu Tiba: %s", Tiba),
		fmt.Sprintf("Jumlah Penumpang: %d", JumlahPenumpang),
		fmt.Sprintf("Rute yang dipilih: %s - %s", PilihRute.Kode, PilihRute.Kereta.Nama),
		fmt.Sprintf("Total Harga: %s", utils.RupiahFormat(PilihRute.Harga*JumlahPenumpang)),
	})

	Konfirmasi := utils.Input("Apakah data sudah benar? (y/n): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}

		if !utils.IsIn(input, []string{"y", "n"}) {
			return false, "Input tidak valid, silakan masukkan 'y' atau 'n'"
		}

		return true, ""
	})

	if Konfirmasi == "n" {
		utils.ClearScreen()
		utils.PrintMessage("Proses pembatalan tiket", "error")
		TiketController()
		return
	}

	utils.PrintHead("Tambah Tiket | Data Penumpang")
	var Penumpang []model.Penumpang
	for i := 0; i < JumlahPenumpang; i++ {

		fmt.Printf("Data Penumpang %d\n", i+1)
		Nama := utils.Input("Masukkan Nama Penumpang: ", func(input string) (bool, string) {
			if input == "" {
				return false, "Nama tidak boleh kosong"
			}
			return true, ""
		})

		NIK := utils.Input("Masukkan NIK Penumpang: ", func(input string) (bool, string) {
			if input == "" {
				return false, "NIK tidak boleh kosong"
			}
			if len(input) != 16 || !utils.IsNumeric(input) {
				return false, "NIK harus berupa 16 digit angka"
			}
			return true, ""
		})

		Gerbong := 0
		utils.Input("Masukkan Gerbong Penumpang: ", func(input string) (bool, string) {
			if input == "" {
				return false, "Gerbong tidak boleh kosong"
			}
			if !utils.IsNumeric(input) {
				return false, "Gerbong harus berupa angka"
			}

			Gerbong, _ = strconv.Atoi(input)
			if Gerbong <= 0 || Gerbong > PilihRute.Gerbong {
				return false, fmt.Sprintf("Gerbong harus antara 1 dan %d", PilihRute.Gerbong)
			}
			return true, ""
		})

		TempatDuduk := utils.Input("Masukkan Tempat Duduk Penumpang: ", func(input string) (bool, string) {
			if input == "" {
				return false, "Tempat duduk tidak boleh kosong"
			}

			pola := `^[A-E](1[0-6]|[1-9])$`
			regex := regexp.MustCompile(pola)

			if !regex.MatchString(input) {
				return false, "Tempat duduk harus dalam format A1, B2, C3, dst. (A-E untuk kolom dan 1-16 untuk baris)"
			}

			return true, ""
		})

		Penumpang = append(Penumpang, model.Penumpang{
			Kode:        fmt.Sprintf("PN%s-%s", PilihRute.Kode, utils.GenerateRandomCode(5)),
			Nama:        Nama,
			NIK:         NIK,
			Gerbong:     Gerbong,
			TempatDuduk: TempatDuduk,
		})
	}

	utils.ClearScreen()

	utils.PrintHead("Tambah Tiket | Konfirmasi Data Penumpang")

	utils.PrintBoxLeft(60, []string{
		fmt.Sprintf("Stasiun Awal: %s - %s", StasiunAwal.Nama, StasiunAwal.Kota),
		fmt.Sprintf("Stasiun Tujuan: %s - %s", StasiunTujuan.Nama, StasiunTujuan.Kota),
		fmt.Sprintf("Waktu Berangkat: %s", Berangkat),
		fmt.Sprintf("Waktu Tiba: %s", Tiba),
		fmt.Sprintf("Jumlah Penumpang: %d", JumlahPenumpang),
		fmt.Sprintf("Rute yang dipilih: %s - %s", PilihRute.Kode, PilihRute.Kereta.Nama),
		fmt.Sprintf("Total Harga: %s", utils.RupiahFormat(PilihRute.Harga*JumlahPenumpang)),
	})

	for i, p := range Penumpang {
		utils.PrintBoxLeft(60, []string{
			fmt.Sprintf("Penumpang Ke-%d", i+1),
			fmt.Sprintf("Kode Penumpang: %s", p.Kode),
			fmt.Sprintf("Nama: %s", p.Nama),
			fmt.Sprintf("NIK: %s", p.NIK),
			fmt.Sprintf("Gerbong: %d", p.Gerbong),
			fmt.Sprintf("Tempat Duduk: %s", p.TempatDuduk),
		})
		utils.Divider("-")
	}

	Confirm := utils.Input("Apakah data penumpang sudah benar? (y/n): ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsIn(input, []string{"y", "n"}) {
			return false, "Input tidak valid, silakan masukkan 'y' atau 'n'"
		}
		return true, ""
	})

	utils.ClearScreen()
	if Confirm == "y" {
		Tiket := model.Tiket{
			Kode:         fmt.Sprintf("TK%s-%s", PilihRute.Kode, utils.GenerateRandomCode(5)),
			Rute:         PilihRute,
			Price:        PilihRute.Harga * JumlahPenumpang,
			User:         model.User{},
			Penumpang:    Penumpang,
			StasiunAwal:  StasiunAwal,
			StasiunAkhir: StasiunTujuan,
		}
		model.ListTiket = append(model.ListTiket, Tiket)
		utils.PrintMessage("Tiket berhasil ditambahkan", "success")
	} else {
		utils.PrintMessage("Proses pembatalan tiket", "error")
	}
	TiketController()
}

func CekTiket() {
	utils.ClearScreen()
	utils.PrintHead("Cek Tiket")

	var Pilihan int
	var Tiket model.Tiket
	var Penumpang model.Penumpang

	fmt.Println("[1] Cek Tiket Berdasarkan Kode Tiket")
	fmt.Println("[2] Cek Tiket Berdasarkan Kode Penumpang")
	fmt.Println("[0] Kembali ke Menu Tiket")
	utils.Divider("-")

	InputPilihan := utils.Input("Masukkan Pilihan: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Pilihan tidak boleh kosong"
		}

		if !utils.IsNumeric(input) {
			return false, "Pilihan harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "0"}) {
			return false, "Pilihan tidak valid"
		}

		return true, ""
	})

	Pilihan, _ = strconv.Atoi(InputPilihan)

	utils.Divider("-")

	Title := "Kode Tiket"
	if Pilihan == 2 {
		Title = "Kode Penumpang"
	}

	utils.Input(fmt.Sprintf("Masukan %s: ", Title), func(input string) (bool, string) {
		if input == "" {
			return false, "Kode tiket tidak boleh kosong"
		}

		Stop := false
		Found := false
		for i := 0; i < len(model.ListTiket) && !Stop; i++ {
			TK := model.ListTiket[i]
			for j := 0; j < len(TK.Penumpang) && !Stop; j++ {
				if Pilihan == 1 && TK.Kode == input {
					Tiket = TK
					Stop = true
					Found = true
				} else if Pilihan == 2 && TK.Penumpang[j].Kode == input {
					Tiket = TK
					Penumpang = TK.Penumpang[j]
					Stop = true
					Found = true
				}
			}
		}

		if !Found {
			return false, "Tiket tidak ditemukan, silakan coba lagi"
		}

		return true, ""
	})

	utils.ClearScreen()
	utils.PrintHead("Cek Tiket | Detail Tiket")
	utils.PrintBoxLeft(60, []string{
		fmt.Sprintf("Kode Tiket: %s", Tiket.Kode),
		fmt.Sprintf("Rute: %s - %s", Tiket.Rute.Kode, Tiket.Rute.Kereta.Nama),
		fmt.Sprintf("Stasiun Awal: %s - %s", Tiket.StasiunAwal.Nama, Tiket.StasiunAwal.Kota),
		fmt.Sprintf("Stasiun Tujuan: %s - %s", Tiket.StasiunAkhir.Nama, Tiket.StasiunAkhir.Kota),
		fmt.Sprintf("Waktu Berangkat: %s", Tiket.Rute.RuteBerhenti[0].Berangkat.Format("02-01-2006 15:04")),
		fmt.Sprintf("Waktu Tiba: %s", Tiket.Rute.RuteBerhenti[len(Tiket.Rute.RuteBerhenti)-1].Tiba.Format("02-01-2006 15:04")),
		fmt.Sprintf("Total Harga: %s", utils.RupiahFormat(Tiket.Price)),
	})

	if Pilihan == 1 {
		fmt.Println("Daftar Penumpang:")
		for _, p := range Tiket.Penumpang {
			utils.PrintBoxLeft(60, []string{
				fmt.Sprintf("Kode Penumpang: %s", p.Kode),
				fmt.Sprintf("Nama: %s", p.Nama),
				fmt.Sprintf("NIK: %s", p.NIK),
				fmt.Sprintf("Gerbong: %d", p.Gerbong),
				fmt.Sprintf("Tempat Duduk: %s", p.TempatDuduk),
			})
		}
	} else {
		utils.PrintBoxLeft(60, []string{
			fmt.Sprintf("Kode Penumpang: %s", Penumpang.Kode),
			fmt.Sprintf("Nama: %s", Penumpang.Nama),
			fmt.Sprintf("NIK: %s", Penumpang.NIK),
			fmt.Sprintf("Gerbong: %d", Penumpang.Gerbong),
			fmt.Sprintf("Tempat Duduk: %s", Penumpang.TempatDuduk),
		})
	}

	fmt.Println(utils.ColorText("[0] Kembali ke Menu Tiket", 90, 49, false))

	utils.Divider("-")

	InputKembali := utils.Input("Masukkan pilihan: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Pilihan tidak boleh kosong"
		}

		if !utils.IsNumeric(input) {
			return false, "Pilihan harus berupa angka"
		}

		if !utils.IsIn(input, []string{"0"}) {
			return false, "Pilihan tidak valid"
		}

		return true, ""
	})

	if InputKembali == "0" {
		utils.ClearScreen()
		TiketController()
	}
}
