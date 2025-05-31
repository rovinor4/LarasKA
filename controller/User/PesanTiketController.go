package user

import (
	"bufio"
	"fmt"
	"laraska/model"
	"laraska/utils"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

func PesanTiket() {

	reader := bufio.NewReader(os.Stdin)
	var Rute = model.Rute{}
	var penumpangs []model.Penumpang
	var tikets = []model.Tiket{}

	var mappedRute []map[string]string

	sortedListStasiun := utils.SelectionSort(model.ListStasiun, func(a, b model.Stasiun) bool {
		return a.IDStasiun < b.IDStasiun
	})

	sortedRuteByIDStasiun := utils.SelectionSort(model.RuteList, func(a, b model.Rute) bool {
		if a.StasiunAwal.IDStasiun != b.StasiunAkhir.IDStasiun {
			return a.StasiunAwal.IDStasiun < b.StasiunAwal.IDStasiun
		}
		return a.StasiunAkhir.IDStasiun < b.StasiunAkhir.IDStasiun
	})

	sortedRuteByKode := utils.SelectionSort(model.RuteList, func(a, b model.Rute) bool {
		return a.Kode < b.Kode
	})

	utils.PrintHead("Pemesanan Tiket Kereta")

	utils.Input("Masukkan kode stasiun awal (e.g. MR): ", func(value string) (bool, string) {
		if value == "" {
			return false, "Input tidak boleh kosong"
		}

		dataStasiunAwal, isStasiunAwalFound, _ := utils.BinaryFindOne(sortedListStasiun, model.Stasiun{IDStasiun: value}, func(a, b model.Stasiun) int {
			if a.IDStasiun < b.IDStasiun {
				return -1
			} else if a.IDStasiun > b.IDStasiun {
				return 1
			}
			return 0
		})

		if !isStasiunAwalFound {
			errFmt := fmt.Sprintf("Stasiun dengan kode \"%s\" tidak ditemukan\n", value)
			return false, errFmt
		} else {
			Rute.StasiunAwal = dataStasiunAwal
			return true, ""
		}

	})

	utils.Input("Masukkan kode stasiun tujuan (e.g. SB): ", func(value string) (bool, string) {
		if value == "" {
			return false, "Kode stasiun akhir tidak boleh kosong"
		}

		dataStasiunAkhir, isStasiunAkhirFound, _ := utils.BinaryFindOne(sortedListStasiun, model.Stasiun{IDStasiun: value}, func(a, b model.Stasiun) int {
			if a.IDStasiun < b.IDStasiun {
				return -1
			} else if a.IDStasiun > b.IDStasiun {
				return 1
			}
			return 0
		})

		if !isStasiunAkhirFound {
			errFmt := fmt.Sprintf("Stasiun dengan kode \"%s\" tidak ditemukan\n", value)
			return false, errFmt
		} else {
			Rute.StasiunAkhir = dataStasiunAkhir
			return true, ""
		}

	})

	ruteTersedia := utils.BinaryFindMany(sortedRuteByIDStasiun, model.Rute{
		StasiunAwal:  Rute.StasiunAwal,
		StasiunAkhir: Rute.StasiunAkhir,
	}, func(a, b model.Rute) int {
		if a.StasiunAwal.IDStasiun < b.StasiunAwal.IDStasiun {
			return -1
		} else if a.StasiunAwal.IDStasiun > b.StasiunAwal.IDStasiun {
			return 1
		}

		if a.StasiunAkhir.IDStasiun < b.StasiunAkhir.IDStasiun {
			return -1
		} else if a.StasiunAkhir.IDStasiun > b.StasiunAkhir.IDStasiun {
			return 1
		}
		return 0
	})

	if len(ruteTersedia) == 0 {
		utils.ClearScreen()
		errFmt := fmt.Sprintf("Rute '%s' -> '%s' saat ini tidak tersedia\n", Rute.StasiunAwal.IDStasiun, Rute.StasiunAkhir.IDStasiun)
		utils.PrintMessage(errFmt, "error")
		PesanTiket()
	}

	for _, rute := range ruteTersedia {
		mappedRute = append(mappedRute, map[string]string{
			"Kode":          rute.Kode,
			"Kereta":        rute.Kereta.Nama,
			"Stasiun Awal":  fmt.Sprintf("%s - %s", rute.StasiunAwal.Kota, rute.StasiunAwal.Nama),
			"Stasiun Akhir": fmt.Sprintf("%s - %s", rute.StasiunAkhir.Kota, rute.StasiunAkhir.Nama),
			"Jadwal":        rute.RuteBerhenti[0].Berangkat.Format("02/01/2006 15:04"),
			"Harga":         strconv.Itoa(rute.Harga),
		})
	}

	utils.ClearScreen()
	utils.PrintTable(
		[]string{"Kode", "Kereta", "Stasiun Awal", "Stasiun Akhir", "Jadwal", "Harga"},
		mappedRute,
		[]string{"Kode", "Kereta", "Stasiun Awal", "Stasiun Akhir", "Jadwal", "Harga"},
		2,
		"Rute Tersedia",
	)

	utils.Input("Pilih Kode Rute: ", func(value string) (bool, string) {
		var found = false
		if value == "" {
			errfmt := fmt.Sprintf("Kode \"%s\" tidak ditemuakan", value)
			return false, errfmt
		}

		for _, rute := range mappedRute {
			if rute["Kode"] == value {
				found = true
			}
		}

		if !found {
			errFmt := fmt.Sprintf("Rute dengan kode \"%s\" tidak ditemukan\n", value)
			return false, errFmt
		}

		ruteResult, isRuteFound, _ := utils.BinaryFindOne(sortedRuteByKode, model.Rute{Kode: value}, func(a, b model.Rute) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if isRuteFound {
			Rute = ruteResult
		}

		return true, ""
	})

	utils.Input("Apakah anda ingin melanjutkan (y/n): ", func(value string) (bool, string) {
		if value == "" {
			return false, ""
		}
		if strings.ToLower(value) == "y" {
			return true, ""
		} else if strings.ToLower(value) == "n" {
			mappedRute = []map[string]string{}
			utils.ClearScreen()
			PesanTiket()
		} else {
			return false, "Pilihan tidak tersedia"
		}

		return false, ""

	})

	jmlPenumpangStr := utils.Input("Masukkan jumlah Penumpang: ", func(value string) (bool, string) {
		if value == "" {
			return false, "Jumlah penumpang tidak boleh kosong"
		}

		jmlPenumpang, errJmlPenumpang := strconv.Atoi(value)
		if jmlPenumpang < 0 || errJmlPenumpang != nil {
			return false, "Jumlah penumpang harus bilangan positif"
		}

		return true, ""

	})

	utils.Divider("-")
	fmt.Println("Nama: ", model.AuthData.User.NamaLengkap)
	fmt.Println("NIK: ", model.AuthData.User.NIK)
	utils.Divider("-")
	utils.Input("Tambah user ini sebagai penumpang (y/n): ", func(value string) (bool, string) {
		if value == "" {
			return false, "Pilihan tidak tersedia"
		}

		if strings.ToLower(value) == "y" {
			penumpangs = append(penumpangs, model.Penumpang{
				Kode: GenerateKodePenumpang(),
				Nama: model.AuthData.User.NamaLengkap,
				NIK:  model.AuthData.User.NIK,
			})
			return true, ""
		} else if strings.ToLower(value) == "n" {
			return true, ""
		} else {
			return false, "Pilihan tidak tersedia"
		}

	})

	jmlPenumpang, _ := strconv.Atoi(jmlPenumpangStr)
	formDetailPenumpang(reader, jmlPenumpang, &penumpangs)

	oldListTiket := model.ListTiket
	formBookingKursi(reader, Rute, &tikets, &penumpangs)

	model.ListTiket = append(model.ListTiket, tikets...)
	if len(model.ListTiket) > len(oldListTiket) {
		utils.ClearScreen()
		// utils.PrintMessage(fmt.Sprint("Tiket Berhasil dipesan atas nama ", strings.TrimSpace(tiket.Penumpang[0].Nama)), "success")
		utils.PrintMessage("Tiket Berhasil dipesan", "success")
		MenuAwalUser()
	} else {
		utils.ClearScreen()
		utils.PrintMessage("Terjadi kesalahan saat memesan tiket!", "error")
		MenuAwalUser()
	}
}

func formDetailPenumpang(reader *bufio.Reader, jmlPenumpang int, penumpangs *[]model.Penumpang) {
	for len(*penumpangs) < jmlPenumpang {
		fmt.Println()
		fmt.Printf("Detail Penumpang #%d\n", len(*penumpangs)+1)
		nama := ""
		isNamaValid := false
		for !isNamaValid {
			fmt.Print("Nama: ")
			input, err := reader.ReadString('\n')
			nama = strings.TrimSpace(input)
			isNamaValid = (nama != "" && err == nil)

			if isNamaValid {
				isNamaDuplikat := false
				i := 0
				for i < len(*penumpangs) && !isNamaDuplikat {
					isNamaDuplikat = ((*penumpangs)[i].Nama == nama)
					i++
				}

				if isNamaDuplikat {
					fmt.Println("Error: Nama sudah ada")
					isNamaValid = false
				}
			}

			if !isNamaValid && nama == "" {
				fmt.Println("Error: Nama tidak boleh kosong")
			}
		}

		nik := ""
		isNIKValid := false
		for !isNIKValid {
			fmt.Print("NIK: ")
			input, err := reader.ReadString('\n')
			nik = strings.TrimSpace(input)

			isNIKValid = true

			if nik == "" || err != nil {
				utils.PrintMessage("Error: NIK tidak boleh kosong", "error")
				isNIKValid = false
			}

			if isNIKValid && !regexp.MustCompile(`^\d{16}$`).MatchString(nik) {
				utils.PrintMessage("Error: Format NIK salah", "error")
				isNIKValid = false
			}

			if isNIKValid {
				isDuplikat := false
				for i := 0; i < len(*penumpangs) && !isDuplikat; i++ {
					if (*penumpangs)[i].NIK == nik {
						isDuplikat = true
					}
				}
				if isDuplikat {
					isNIKValid = false
					utils.PrintMessage("Error: NIK sudah terdaftar", "error")
				}
			}
		}

		*penumpangs = append(*penumpangs, model.Penumpang{
			Kode: GenerateKodePenumpang(),
			Nama: nama,
			NIK:  nik,
		})
	}
}
func formBookingKursi(reader *bufio.Reader, rute model.Rute, tikets *[]model.Tiket, penumpangs *[]model.Penumpang) {
	// gabung model.ListTiket dengan tiket yang akan dibuat
	// Tickets = model.ListTiket
	Tikets := slices.Clone(model.ListTiket)
	Tikets = append(Tikets, *tikets...)

	for p := 0; p < len(*penumpangs); {
		fmt.Printf("\nBooking kursi untuk %s (%s)\n", (*penumpangs)[p].Nama, (*penumpangs)[p].Kode)
		utils.Divider("-")

		gerbong := 0
		isGerbongValid := false
		for !isGerbongValid {
			promptGerbong := fmt.Sprintf("Masukkan nomor gerbong (1-%d): ", rute.Gerbong)
			utils.Input(promptGerbong, func(value string) (bool, string) {
				if value == "" {
					return false, "Error: Gerbong tidak boleh kosong"
				}

				gerbongAsInt, errGerbongAsInt := strconv.Atoi(value)
				if gerbongAsInt > rute.Gerbong {
					errFmt := fmt.Sprintf("Error: %s hanya menyediakan %d gerbong", rute.Kereta.Nama, rute.Gerbong)
					return false, errFmt
				} else if errGerbongAsInt != nil || gerbongAsInt <= 0 {
					return false, "Harap masukkan angka positif"
				} else {
					gerbong = gerbongAsInt
					isGerbongValid = true
					return true, ""
				}

			})
		}

		tempatDuduk := ""
		isKursiValid := false
		for !isKursiValid {
			fmt.Print("Masukkan nomor tempat duduk: ")
			inputKursi, errKursi := reader.ReadString('\n')
			tempatDuduk = strings.TrimSpace(inputKursi)

			if tempatDuduk == "" || errKursi != nil {
				utils.PrintMessage("Mohon memesan tempat duduk untuk kenyamanan anda", "error")
			} else {
				isKursiAvailable := true

				for _, tiket := range Tikets {
					if tiket.Rute.Kereta.Nama == rute.Kereta.Nama &&
						tiket.Gerbong == gerbong &&
						tiket.TempatDuduk == tempatDuduk {
						isKursiAvailable = false
						utils.PrintMessage("Maaf kursi sudah dipesan", "error")
					}
				}

				if isKursiAvailable {
					tiket := model.Tiket{
						// Penumpang:   []model.Penumpang{(*penumpangs)[p]},
						Kode:        GenerateKodeTiket(),
						User:        model.User{NamaLengkap: (*penumpangs)[p].Nama, NIK: (*penumpangs)[p].NIK},
						Rute:        rute,
						Price:       rute.Harga,
						Gerbong:     gerbong,
						TempatDuduk: tempatDuduk,
						CreatedAt:   time.Now(),
					}
					*tikets = append(*tikets, tiket)
					fmt.Println(*tikets)
					// update list semua tiket termasuk tiket yang baru dipesen
					Tikets = append(Tikets, tiket)
					isKursiValid = true
					p++
				}
			}
		}
	}
}
