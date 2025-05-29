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
	// var mappedStasiunAwal []map[string]string
	fmt.Print("Masukkan kode stasiun awal (case sensitive e.g. MR): ")
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
	fmt.Print("Masukkan kode stasiun tujuan (case sensitive e.g. SB): ")
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

	var found = false
	utils.ClearScreen()
	utils.PrintTable(
		[]string{"Kode", "Nama", "Kereta", "Stasiun Awal", "Stasiun Akhir", "Harga"},
		mappedRute,
		[]string{"Kode", "Nama", "Kereta", "Stasiun Awal", "Stasiun Akhir", "Harga"},
		4,
		"Rute Tersedia",
	)

Confirmation:
	fmt.Print("Apakah anda ingin melanjnutkan (Y/n): ")
	Continue, _ := reader.ReadString('\n')
	Continue = strings.TrimSpace(Continue)
	if strings.ToLower(Continue) == "y" {
		goto FormPilihRute
	} else if strings.ToLower(Continue) == "n" {
		mappedRute = []map[string]string{}
		utils.ClearScreen()
		goto FormStasiunAwal
	} else {
		goto Confirmation
	}

FormPilihRute:
	fmt.Print("Pilih Kode Rute: ")
	inputKodeRute, errKode := reader.ReadString('\n')
	inputKodeRute = strings.TrimSpace(inputKodeRute)
	if inputKodeRute == "" || errKode != nil {
		fmt.Printf("Kode \"%s\" tidak ditemukan\n", inputKodeRute)
		goto FormPilihRute
	}

	for _, rute := range mappedRute {
		if rute["Kode"] == inputKodeRute {
			found = true
		}
	}

	if !found {
		fmt.Printf("Kode \"%s\" tidak ditemukan", inputKodeRute)
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

FormPilihJadwal:
	fmt.Print("Masukkan tanggal (DD/MM/YYYY): ")
	Jadwal, errJadwal := reader.ReadString('\n')
	Jadwal = strings.TrimSpace(Jadwal)
	if Jadwal == "" || errJadwal != nil {
		fmt.Println("Jadwal harus diisi")
		goto FormPilihJadwal
	}

	if !regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/[0-9]{4}$`).MatchString(Jadwal) {
		fmt.Println("Format jadwal salah")
		goto FormPilihJadwal
	}

	parsedJadwal, errParse := time.Parse("02/01/2006", Jadwal)
	if errParse != nil {
		fmt.Println("Format tanggal salah: ", errParse)
		goto FormPilihJadwal
	}

	if parsedJadwal.Before(time.Now()) {
		fmt.Println("Tanggal yang anda inputkan sudah kadaluarsa")
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
	userAsPenumpang, errUserAsPenumpang := reader.ReadString('\n')
	userAsPenumpang = strings.TrimSpace(userAsPenumpang)

	if strings.ToLower(userAsPenumpang) == "y" {
		jmlPenumpang -= 1

		penumpang = append(penumpang, model.Penumpang{
			// TODO: Kode,
			Nama: controller.AuthData.User.NamaLengkap,
			NIK:  controller.AuthData.User.NIK,
		})

		penumpangs := formPenumpang(reader, jmlPenumpang)
		penumpang = append(penumpang, penumpangs...)
		goto Complete
	} else if userAsPenumpang == "" || errUserAsPenumpang != nil {
		goto FormJumlahPenumpang
	} else if strings.ToLower(userAsPenumpang) == "n" {
		penumpangs := formPenumpang(reader, jmlPenumpang)
		penumpang = append(penumpangs, penumpangs...)
		goto Complete
	}

Complete:
	tiket = model.Tiket{
		// TODO: Kode
		User:      model.User{NamaLengkap: penumpang[0].Nama, NIK: penumpang[0].NIK}, // User??
		Rute:      Rute,
		Price:     jmlPenumpang * Rute.Harga,
		Penumpang: penumpang,
		CreatedAt: time.Now(),
	}
	model.ListTiket = append(model.ListTiket, tiket)
	if len(model.ListTiket) > 0 {
		fmt.Println("Tiket berhasil dipesan!")

		for _, t := range model.ListTiket {
			fmt.Println("List Tiket")
			fmt.Println("Nama: ", t.User.NamaLengkap)
			fmt.Println("CreatedAt: ", t.CreatedAt)
			fmt.Printf("Rute: %s -> %s\n", t.Rute.StasiunAwal, t.Rute.StasiunAkhir)
			fmt.Println("Price: ", t.Price)
			fmt.Println("Penumpang: ", t.Penumpang)

		}
		utils.ClearScreen()
		fmt.Println("Tiket Berhasil dipesan atas nama ", tiket.User.NamaLengkap)
		MenuAwalUser()
	} else {
		fmt.Println("Gagal memesan tiket")

	}
}

func formPenumpang(reader *bufio.Reader, jmlPenumpang int) []model.Penumpang {
	var penumpang []model.Penumpang
	for i := 0; i < jmlPenumpang; i++ {
	DataPenumpang:
		fmt.Println("Data Penumpang")
		fmt.Print("Nama: ")
		nama, errNama := reader.ReadString('\n')
		nama = strings.TrimSpace(nama)
		if nama == "" || errNama != nil {
			fmt.Println("Nama penumpang tidak boleh kosong")
			goto DataPenumpang
		}

		fmt.Print("NIK: ")
		nik, errNik := reader.ReadString('\n')
		nik = strings.TrimSpace(nik)
		if nik == "" || errNik != nil {
			fmt.Println("Nama penumpang tidak boleh kosong")
			goto DataPenumpang
		}

		var nikRegex = regexp.MustCompile(`^\d{16}$`)
		if !nikRegex.MatchString(nik) {
			utils.PrintMessage("Format NIK salah", "error")
			goto DataPenumpang
		}

		penumpang = append(penumpang, model.Penumpang{
			Nama: nama,
			NIK:  nik,
		})
	}

	return penumpang
}

// TODO: ShowHistoryTiket()
func ShowHistoryTiket() {
	reader := bufio.NewReader(os.Stdin)
	if len(model.ListTiket) > 0 {
		utils.Divider("-")
		for _, tiket := range model.ListTiket {
			fmt.Println("Nama: ", tiket.User.NamaLengkap)
			fmt.Printf("Dari %s - %s ke %s - %s\n", tiket.Rute.StasiunAwal.Kota, tiket.Rute.StasiunAwal.Nama, tiket.Rute.StasiunAkhir.Kota, tiket.Rute.StasiunAwal.Nama)
			fmt.Println("Penumpang: ", tiket.Penumpang)
			fmt.Println("Dibuat pada: ", tiket.CreatedAt.Format("02/01/2006"))
			utils.Divider("-")
		}

	Menu:
		fmt.Print("Kembali ke menu awal (Y/n)")
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
	} else {
		utils.ClearScreen()
		fmt.Println("Belum ada tiket yang dipesan!")
		MenuAwalUser()
	}
}
