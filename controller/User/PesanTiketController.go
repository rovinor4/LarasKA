package user

import (
	"bufio"
	"fmt"
	"laraska/controller"
	"laraska/model"
	"laraska/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func PesanTiket() {

	reader := bufio.NewReader(os.Stdin)
	var Rute = model.Rute{}
	var tiket = model.Tiket{}
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

FormStasiunAwal:
	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		"Pemesanan Tiket Kereta",
	})
	fmt.Print("Masukkan kode stasiun awal (e.g. MR): ")
	stasiunAwal, errSA := reader.ReadString('\n')
	stasiunAwal = strings.TrimSpace(stasiunAwal)
	if stasiunAwal == "" || errSA != nil {
		fmt.Println("kosong")
		fmt.Println("errSA", errSA)
		goto FormStasiunAwal
	}

	dataStasiunAwal, isStasiunAwalFound := utils.BinaryFindOne(sortedListStasiun, model.Stasiun{IDStasiun: stasiunAwal}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})

	if !isStasiunAwalFound {
		fmt.Printf("Stasiun dengan kode \"%s\" tidak ditemukan\n", stasiunAwal)
		goto FormStasiunAwal
	}

	Rute.StasiunAwal = dataStasiunAwal

FormStasiunAkhir:
	fmt.Print("Masukkan kode stasiun tujuan (e.g. SB): ")
	stasiunAkhir, errST := reader.ReadString('\n')
	stasiunAkhir = strings.TrimSpace(stasiunAkhir)
	if stasiunAwal == "" || errST != nil {
		fmt.Println("errST: ", errST)
		goto FormStasiunAkhir
	}

	dataStasiunAkhir, isStasiunAkhirFound := utils.BinaryFindOne(sortedListStasiun, model.Stasiun{IDStasiun: stasiunAkhir}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})

	if !isStasiunAkhirFound {
		fmt.Printf("Stasiun dengan kode \"%s\" tidak ditemukan\n", stasiunAkhir)
		goto FormStasiunAkhir
	}

	Rute.StasiunAkhir = dataStasiunAkhir

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
		fmt.Printf("rute '%s' -> '%s' saat ini tidak tersedia\n", Rute.StasiunAwal.IDStasiun, Rute.StasiunAkhir.IDStasiun)
		goto FormStasiunAwal
	}

	for _, rute := range ruteTersedia {
		mappedRute = append(mappedRute, map[string]string{
			"Kode":          rute.Kode,
			"Nama":          rute.Nama,
			"Kereta":        rute.Kereta.Nama,
			"Stasiun Awal":  fmt.Sprintf("%s - %s", rute.StasiunAwal.Kota, rute.StasiunAwal.Nama),
			"Stasiun Akhir": fmt.Sprintf("%s - %s", rute.StasiunAkhir.Kota, rute.StasiunAkhir.Nama),
			"Harga":         strconv.Itoa(rute.Harga),
		})
	}

	utils.ClearScreen()
	utils.PrintTable(
		[]string{"Kode", "Nama", "Kereta", "Stasiun Awal", "Stasiun Akhir", "Harga"},
		mappedRute,
		[]string{"Kode", "Nama", "Kereta", "Stasiun Awal", "Stasiun Akhir", "Harga"},
		4,
		"Rute Tersedia",
	)

FormPilihRute:
	var found = false
	fmt.Print("Pilih Kode Rute: ")
	inputKodeRute, errKode := reader.ReadString('\n')
	inputKodeRute = strings.TrimSpace(inputKodeRute)
	if inputKodeRute == "" || errKode != nil {
		utils.PrintMessage("Kode rute tidak boleh kosong", "error")
		goto FormPilihRute
	}

	for _, rute := range mappedRute {
		if rute["Kode"] == inputKodeRute {
			found = true
		}
	}

	if !found {
		errFmt := fmt.Sprintf("Rute dengan kode \"%s\" tidak ditemukan\n", inputKodeRute)
		utils.PrintMessage(errFmt, "error")
		goto FormPilihRute
	}

	ruteResult, isRuteFound := utils.BinaryFindOne(sortedRuteByKode, model.Rute{Kode: inputKodeRute}, func(a, b model.Rute) int {
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

Confirmation:
	fmt.Print("Apakah anda ingin melanjutkan (y/n): ")
	Continue, _ := reader.ReadString('\n')
	Continue = strings.TrimSpace(Continue)
	if strings.ToLower(Continue) == "y" {
		goto FormPilihJadwal
	} else if strings.ToLower(Continue) == "n" {
		mappedRute = []map[string]string{}
		utils.ClearScreen()
		goto FormStasiunAwal
	} else {
		goto Confirmation
	}

FormPilihJadwal:
	utils.ClearScreen()
	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		"Detail Pemesanan",
	})
	fmt.Print("Masukkan tanggal (DD/MM/YYYY): ")
	Jadwal, errJadwal := reader.ReadString('\n')
	Jadwal = strings.TrimSpace(Jadwal)
	if Jadwal == "" || errJadwal != nil {
		utils.PrintMessage("Jadwal harus diisi", "error")
		goto FormPilihJadwal
	}

	if !regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/[0-9]{4}$`).MatchString(Jadwal) {
		utils.PrintMessage("Format jadwal salah", "error")
		goto FormPilihJadwal
	}

	parsedJadwal, errParse := time.Parse("02/01/2006", Jadwal)
	if errParse != nil {
		utils.PrintMessage("Format jadwal salah (e.g DD/MM/YYY)", "error")
		goto FormPilihJadwal
	}

	if parsedJadwal.Before(time.Now()) {
		utils.PrintMessage("Tanggal yang anda inputkan sudah kadaluarsa", "error")
		goto FormPilihJadwal
	}
FormJumlahPenumpang:
	var penumpang []model.Penumpang
	fmt.Print("Masukan jumlah penumpang: ")
	jmlPenumpangStr, errJmlPenumpangStr := reader.ReadString('\n')
	jmlPenumpangStr = strings.TrimSpace(jmlPenumpangStr)
	if jmlPenumpangStr == "" || errJmlPenumpangStr != nil {
		fmt.Println("Jumlah penumpang tidak boleh kosong")
		goto FormJumlahPenumpang
	}

	jmlPenumpang, errJmlPenumpang := strconv.Atoi(jmlPenumpangStr)
	if jmlPenumpang < 0 || errJmlPenumpang != nil {
		fmt.Println("Jumlah penumpang harus bilangan positif")
		goto FormJumlahPenumpang
	}

	utils.Divider("-")
	fmt.Println("Nama: ", controller.AuthData.User.NamaLengkap)
	fmt.Println("NIK: ", controller.AuthData.User.NIK)
	utils.Divider("-")

	fmt.Print("Tambah user ini sebagai penumpang? (y/n): ")
	userAsPenumpang, _ := reader.ReadString('\n')
	userAsPenumpang = strings.TrimSpace(userAsPenumpang)

	if strings.ToLower(userAsPenumpang) == "y" {
		penumpang = append(penumpang, model.Penumpang{
			// TODO: Kode,
			Nama: controller.AuthData.User.NamaLengkap,
			NIK:  controller.AuthData.User.NIK,
		})
	} else if strings.ToLower(userAsPenumpang) == "n" {
		formDetailPenumpang(reader, jmlPenumpang, &penumpang)
	} else {
		utils.PrintMessage("Masukkan y/n", "error")
		goto FormJumlahPenumpang
	}

	formDetailPenumpang(reader, jmlPenumpang, &penumpang)

	tiket = model.Tiket{
		// TODO: Kode
		User:      model.User{NamaLengkap: penumpang[0].Nama, NIK: penumpang[0].NIK},
		Rute:      Rute,
		Price:     jmlPenumpang * Rute.Harga,
		Penumpang: penumpang,
		CreatedAt: time.Now(),
	}
	oldListTiket := model.ListTiket
	model.ListTiket = append(model.ListTiket, tiket)
	if len(model.ListTiket) > len(oldListTiket) {
		utils.ClearScreen()
		utils.PrintMessage(fmt.Sprint("Tiket Berhasil dipesan atas nama ", strings.TrimSpace(tiket.User.NamaLengkap)), "success")
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
			Nama: nama,
			NIK:  nik,
		})
	}
}

func ShowHistoryTiket() {
	reader := bufio.NewReader(os.Stdin)
	var mappedTiket []map[string]string
	var isUserIncluded bool

	ascendingTiketByTimeCreated := utils.SelectionSort(model.ListTiket, func(a, b model.Tiket) bool {
		return a.CreatedAt.Before(b.CreatedAt)
	})

	if len(model.ListTiket) == 0 {
		utils.ClearScreen()
		utils.PrintMessage("Belum ada tiket yang dipesan!", "error")
		MenuAwalUser()
	}

	for i, tiket := range ascendingTiketByTimeCreated {

		if strings.TrimSpace(tiket.Penumpang[0].Nama) == strings.TrimSpace(controller.AuthData.User.NamaLengkap) {
			isUserIncluded = true
		} else {
			isUserIncluded = false
		}
		jmlPenumpang := ""

		if isUserIncluded {
			jmlPenumpang += fmt.Sprintf("%s (Included)", strconv.Itoa(len(tiket.Penumpang)))
		} else {
			jmlPenumpang += fmt.Sprint(strconv.Itoa(len(tiket.Penumpang)))
		}

		mappedTiket = append(mappedTiket, map[string]string{
			"No.":               strconv.Itoa(i),
			"Kode":              "X", // TODO: Kode
			"Tanggal Pembuatan": tiket.CreatedAt.Format("02/01/2006 15:04"),
			"Atas Nama":         tiket.User.NamaLengkap,
			"Dari/Ke":           fmt.Sprintf("%s - %s", tiket.Rute.StasiunAwal.Nama, tiket.Rute.StasiunAkhir.Nama),
			"Jumlah Penumpang":  jmlPenumpang,
			"Total Harga":       fmt.Sprint("Rp. ", strconv.Itoa(tiket.Price)),
		})
	}

	utils.PrintTable(
		[]string{"Kode", "Tanggal Pembuatan", "Atas Nama", "Dari/Ke", "Jumlah Penumpang", "Total Harga"},
		mappedTiket,
		[]string{"Kode", "Tanggal Pembuatan", "Atas Nama", "Dari/Ke", "Jumlah Penumpang", "Total Harga"},
		2,
		"History Pemesanan Tiket",
	)

Menu: // TODO: Menu
	fmt.Print("Kembali ke menu awal (Y/n): ")
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" || err != nil {
		goto Menu
	} else if strings.ToLower(input) == "y" {
		utils.ClearScreen()
		MenuAwalUser()
	} else {
		goto Menu
	}
}
