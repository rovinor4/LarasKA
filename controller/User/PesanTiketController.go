package User

import (
	"bufio"
	"fmt"
	"laraska/controller"
	"laraska/model"
	"laraska/utils"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func PesanTiket(authUser model.User) {

	reader := bufio.NewReader(os.Stdin)
	var Rute = model.Rute{}
	var penumpangs []model.Penumpang
	var tikets = []model.Tiket{}

	var mappedRute []map[string]string

	handleStasiunInput(authUser, "Stasiun Awal", model.ListStasiun, &Rute)
	utils.ClearScreen()
	handleStasiunInput(authUser, "Stasiun Tujuan", model.ListStasiun, &Rute)

	ruteTersedia := utils.FindMany(model.ListRute, model.Rute{
		StasiunAwal:  Rute.StasiunAwal,
		StasiunAkhir: Rute.StasiunAkhir,
		// RuteBerhenti: []model.RuteBerhenti{{Berangkat: Rute.RuteBerhenti[0].Berangkat}},
	}, func(a, b model.Rute) int {
		if a.StasiunAwal.IDStasiun == b.StasiunAwal.IDStasiun &&
			a.StasiunAkhir.IDStasiun == b.StasiunAkhir.IDStasiun {
			return 0
		} else {
			return 1
		}
	})

	if len(ruteTersedia) == 0 {
		utils.ClearScreen()
		errFmt := fmt.Sprintf("Rute '%s' -> '%s' saat ini tidak tersedia\n", Rute.StasiunAwal.IDStasiun, Rute.StasiunAkhir.IDStasiun)
		utils.PrintMessage(errFmt, "error")
		PesanTiket(authUser)
	}

	for no, rute := range ruteTersedia {
		mappedRute = append(mappedRute, map[string]string{
			"No":            strconv.Itoa(no + 1),
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
		[]string{"No", "Kode", "Kereta", "Stasiun Awal", "Stasiun Akhir", "Jadwal", "Harga"},
		mappedRute,
		[]string{"No", "Kode", "Kereta", "Stasiun Awal", "Stasiun Akhir", "Jadwal", "Harga"},
		2,
		"Rute Tersedia",
	)

	utils.Input("Pilih nomor rute: ", func(value string) (bool, string) {
		if value == "" {
			errfmt := fmt.Sprintf("Kode \"%s\" tidak ditemuakan", value)
			return false, errfmt
		}

		noRute, errNoRute := strconv.Atoi(value)
		if errNoRute != nil {
			return false, "Input harus berupa angka"
		} else if noRute < 1 || noRute > len(ruteTersedia) {
			return false, "Nomor yang anda masukkan tidak sesuai"
		}

		for _, rute := range mappedRute {
			if rute["No"] == value {
				Rute = ruteTersedia[noRute-1]
				return true, ""
			}
		}

		return false, ""
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
			PesanTiket(authUser)
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
	fmt.Println("Nama: ", controller.AuthData.User.NamaLengkap)
	fmt.Println("NIK: ", controller.AuthData.User.NIK)
	utils.Divider("-")
	utils.Input("Tambah user ini sebagai penumpang (y/n): ", func(value string) (bool, string) {
		if value == "" {
			return false, "Pilihan tidak tersedia"
		}

		if strings.ToLower(value) == "y" {
			penumpangs = append(penumpangs, model.Penumpang{
				Kode: GenerateKodePenumpang(),
				Nama: controller.AuthData.User.NamaLengkap,
				NIK:  controller.AuthData.User.NIK,
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
	formCreateTiket(reader, Rute, &tikets, &penumpangs)

	model.ListTiket = append(model.ListTiket, tikets...)
	if len(model.ListTiket) > len(oldListTiket) {
		utils.ClearScreen()
		utils.PrintMessage("Tiket Berhasil dipesan", "success")
		MenuAwalUser(authUser)
	} else {
		utils.ClearScreen()
		utils.PrintMessage("Terjadi kesalahan saat memesan tiket!", "error")
		MenuAwalUser(authUser)
	}
}

func handleStasiunInput(authUser model.User, stasiun string, listStasiun []model.Stasiun, rute *model.Rute) {
	var mappedStasiun []map[string]string
	var stasiunLower = strings.ToLower(strings.ReplaceAll(stasiun, " ", ""))

	for _, stat := range listStasiun {
		mappedStasiun = append(mappedStasiun, map[string]string{
			"idStasiun": stat.IDStasiun,
			"stasiun":   stat.Nama,
			"kota":      stat.Kota,
		})
	}
	utils.PrintTable(
		[]string{"ID Stasiun", "Stasiun", "Kota"},
		mappedStasiun,
		[]string{"idStasiun", "stasiun", "kota"},
		4,
		"Pemesanan Tiket Kereta | List Stasiun",
	)

	fmt.Println("[1] Search")
	fmt.Println("[2] Pilih Stasiun")
	fmt.Println("[3] Kembali")
	utils.Divider("-")

	input := utils.Input("Pilih Menu: ", func(value string) (bool, string) {
		if value == "" {
			return false, ""
		}

		if !utils.IsIn(value, []string{"1", "2", "3"}) {
			return false, "Pilihan tidak tersedia"
		}

		if !utils.IsNumeric(value) {
			return false, "Pilihan harus berupa angka"
		}

		return true, ""
	})

	inputAsInt, _ := strconv.Atoi(input)
	switch inputAsInt {
	case 1:
		handleSearch(&listStasiun)
		handleStasiunInput(authUser, stasiun, listStasiun, rute)
		return
	case 2:
	case 3:
		utils.ClearScreen()
		MenuAwalUser(authUser)
	}

	sortedListStasiun := utils.SelectionSort(listStasiun, func(a, b model.Stasiun) bool {
		return a.IDStasiun < b.IDStasiun
	})

	inputPrompt := fmt.Sprintf("Masukkan kode %s: ", stasiun)

	if stasiunLower == "stasiunawal" {
		utils.Input(inputPrompt, func(value string) (bool, string) {
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
				rute.StasiunAwal = dataStasiunAwal
				return true, ""
			}

		})
	}

	if stasiunLower == "stasiuntujuan" {
		utils.Input(inputPrompt, func(value string) (bool, string) {
			if value == "" {
				return false, "Input tidak boleh kosong"
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
				rute.StasiunAkhir = dataStasiunAkhir
				return true, ""
			}
		})
	}

}
func handleSearch(listStasiun *[]model.Stasiun) {

	keyword := utils.Input("Masukkan Keyword (IDStasiun/Kota/Nama): ", func(value string) (bool, string) {
		return true, ""
	})
	keyword = strings.ToLower(keyword)

	if keyword == "" {
		utils.ClearScreen()
		return
	}

	var newListStasiun []model.Stasiun
	for _, stasiun := range model.ListStasiun {
		if strings.Contains(strings.ToLower(stasiun.IDStasiun), keyword) ||
			strings.Contains(strings.ToLower(stasiun.Kota), keyword) ||
			strings.Contains(strings.ToLower(stasiun.Nama), keyword) {
			newListStasiun = append(newListStasiun, stasiun)
		}
	}
	*listStasiun = newListStasiun
	utils.ClearScreen()
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
			Kode: GenerateKodePenumpang(), // TODO: Simplified kode generator
			Nama: nama,
			NIK:  nik,
		})
	}
}

func formCreateTiket(reader *bufio.Reader, rute model.Rute, tikets *[]model.Tiket, penumpangs *[]model.Penumpang) {
	Tikets := slices.Clone(model.ListTiket)
	Tikets = append(Tikets, *tikets...)

	for p := 0; p < len(*penumpangs); {
		fmt.Printf("\nBooking kursi untuk %s (%s)\n", (*penumpangs)[p].Nama, (*penumpangs)[p].Kode)
		utils.Divider("-")

		gerbong := 0
		isGerbongValid := false
		for !isGerbongValid {
			fmt.Print("Masukkan nomor gerbong: ")
			inputGerbong, errGerbong := reader.ReadString('\n')
			inputGerbong = strings.TrimSpace(inputGerbong)

			gerbongAsInt, errGerbongAsInt := strconv.Atoi(inputGerbong)
			if inputGerbong == "" || errGerbong != nil {
				utils.PrintMessage("Error: Gerbong tidak boleh kosong", "error")
			} else if gerbongAsInt > rute.Gerbong {
				errFmt := fmt.Sprintf("Error: %s hanya menyediakan %d gerbong", rute.Kereta.Nama, rute.Gerbong)
				utils.PrintMessage(errFmt, "error")
			} else if errGerbongAsInt != nil || gerbongAsInt <= 0 {
				utils.PrintMessage("Error: Harap masukkan angka positif", "error")
			} else {
				gerbong = gerbongAsInt
				isGerbongValid = true
			}
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

				for t := 0; t < len(Tikets) && isKursiAvailable; t++ {
					for penumpang := 0; penumpang < len(Tikets[t].Penumpang) && isKursiAvailable; {
						if Tikets[t].Penumpang[penumpang].Gerbong == gerbong &&
							Tikets[t].Penumpang[penumpang].TempatDuduk == tempatDuduk {
							isKursiAvailable = false
							utils.PrintMessage("Maaf kursi sudah dipesan", "error")
						}
						penumpang++
					}
				}

				if isKursiAvailable {
					tiket := model.Tiket{
						Kode:  GenerateKodeTiket(), // TODO: Simplified kode generator
						Rute:  rute,
						Price: rute.Harga,
						User:  model.User{NamaLengkap: (*penumpangs)[0].Nama, NIK: (*penumpangs)[0].NIK},
						Penumpang: []model.Penumpang{
							{
								Kode:        (*penumpangs)[p].Kode,
								Nama:        (*penumpangs)[p].Nama,
								NIK:         (*penumpangs)[p].NIK,
								Gerbong:     gerbong,
								TempatDuduk: tempatDuduk,
							},
						},
						StasiunAwal:  rute.StasiunAwal,
						StasiunAkhir: rute.StasiunAkhir,
					}
					*tikets = append(*tikets, tiket)
					Tikets = append(Tikets, tiket)
					isKursiValid = true
					p++
				}
			}
		}
	}
}
