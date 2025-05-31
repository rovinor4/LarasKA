package Admin

import (
	"bufio"
	"fmt"
	"laraska/model"
	"laraska/utils"
	"os"
	"strconv"
	"strings"
)

func RuteController() {
	var reader = bufio.NewReader(os.Stdin)

	var RuteListMap []map[string]string

	for _, rute := range model.RuteList {
		RuteListMap = append(RuteListMap, map[string]string{
			"Kode":         rute.Kode,
			"Nama":         rute.Nama,
			"Harga":        utils.RupiahFormat(rute.Harga),
			"StasiunAwal":  rute.StasiunAwal.Nama,
			"StasiunAkhir": rute.StasiunAkhir.Nama,
		})
	}
	utils.PrintTable([]string{"Kode", "Nama", "Harga", "Stasiun Awal", "Stasiun Akhir"}, RuteListMap, []string{"Kode", "Nama", "Harga", "StasiunAwal", "StasiunAkhir"}, 5, "Rute Kereta Api")

	fmt.Println(utils.ColorText("[1] Lihat Detail", 90, 49, false))
	fmt.Println(utils.ColorText("[2] Tambah Rute", 90, 49, false))
	fmt.Println(utils.ColorText("[0] Kembali", 90, 49, false))
	utils.Divider("-")
Step1:
	fmt.Print("Pilih menu: ")
	Select, err := reader.ReadString('\n')
	Select = strings.TrimSpace(Select)

	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step1
	}

	if Select == "" && !utils.IsNumeric(Select) {
		utils.PrintMessage("Pilihan tidak boleh kosong & harus berupa angka", "error")
	}

Step2:
	switch Select {
	case "1":
		fmt.Print("Pilih Kode Rute: ")
		Kode, err := reader.ReadString('\n')
		Kode = strings.TrimSpace(Kode)

		if err != nil {
			utils.PrintMessage("Terjadi kesalahan", "error")
			goto Step2
		}

		if Kode == "" {
			utils.PrintMessage("Kode Rute Tidak Boleh Kosong", "error")
			goto Step2
		}

		Data, Have, _ := utils.FindOne(model.RuteList, model.Rute{Kode: Kode}, func(a, b model.Rute) int {
			if a.Kode < b.Kode {
				return -1
			} else if a.Kode > b.Kode {
				return 1
			}
			return 0
		})

		if !Have {
			utils.PrintMessage("Rute tidak ditemukan", "error")
			goto Step2
		}
		utils.ClearScreen()
		DetailRute(Data)
	case "2":
		utils.ClearScreen()
		TambahRute()
	default:
		utils.PrintMessage("Pilihan tidak valid", "error")
		goto Step1
	}

}

func DetailRute(Rute model.Rute) {

	var Select string
	reader := bufio.NewReader(os.Stdin)

	// PrintHead("Detail Rute Kereta Api")
	fmt.Println(utils.AlignTeksCenter("Detail Rute Kereta Api", 60))

	utils.Divider("-")
	fmt.Println("Kode Rute: ", Rute.Kode)
	fmt.Println("Nama Rute: ", Rute.Nama)
	fmt.Println("Harga: ", utils.RupiahFormat(Rute.Harga))
	fmt.Println("Stasiun Awal: ", Rute.StasiunAwal.Nama)
	fmt.Println("Stasiun Akhir: ", Rute.StasiunAkhir.Nama)
	utils.Divider("-")
	fmt.Println("Rute Berhenti: ")
	for index, stasiun := range Rute.RuteBerhenti {
		if Rute.HargaTetap {
			fmt.Printf("%d. %s - %s - %s (%s-%s)\n", index+1, stasiun.StasiunAwal.Nama, stasiun.StasiunAkhir.Nama, utils.RupiahFormat(Rute.Harga), stasiun.Berangkat.Format("15:04"), stasiun.Tiba.Format("15:04"))
		} else {
			fmt.Printf("%d. %s - %s - %s (%s-%s)\n", index+1, stasiun.StasiunAwal.Nama, stasiun.StasiunAkhir.Nama, utils.RupiahFormat(Rute.Harga*(index+1)), stasiun.Berangkat.Format("15:04"), stasiun.Tiba.Format("15:04"))
		}
	}
	utils.Divider("-")

	fmt.Println(utils.ColorText("[1] Edit Rute", 90, 49, false))
	fmt.Println(utils.ColorText("[2] Hapus Rute", 90, 49, false))
	fmt.Println(utils.ColorText("[0] Kembali", 90, 49, false))

	utils.Divider("-")
Step1:
	fmt.Print("Pilih menu: ")
	Select, err := reader.ReadString('\n')
	Select = strings.TrimSpace(Select)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step1
	}
	if Select == "" && !utils.IsNumeric(Select) {
		utils.PrintMessage("Pilihan tidak boleh kosong & harus berupa angka", "error")
		goto Step1
	}

	// Step2:
	switch Select {
	case "1":
		utils.ClearScreen()
		EditRute(Rute)
	case "2":
		DeleteRute(Rute)
	case "0":
		utils.ClearScreen()
		RuteController()
	default:
		utils.PrintMessage("Pilihan tidak valid", "error")
		goto Step1
	}
}

func TambahRute() {
	reader := bufio.NewReader(os.Stdin)

	utils.PrintHead("Tambah Data Rute")

Step1:
	fmt.Print("Masukan Kode Rute: ")
	Kode, err := reader.ReadString('\n')
	Kode = strings.TrimSpace(Kode)

	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step1
	}

	if Kode == "" {
		utils.PrintMessage("Kode Rute tidak boleh kosong", "error")
		goto Step1
	}

	//check unique
	_, HaveKode, _ := utils.FindOne(model.RuteList, model.Rute{Kode: Kode}, func(a, b model.Rute) int {
		if a.Kode < b.Kode {
			return -1
		} else if a.Kode > b.Kode {
			return 1
		}
		return 0
	})

	if HaveKode {
		utils.PrintMessage("Kode Rute sudah ada", "error")
		goto Step1
	}

Step2:
	fmt.Print("Masukan Nama Rute: ")
	Nama, err := reader.ReadString('\n')
	Nama = strings.TrimSpace(Nama)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step2
	}
	if Nama == "" {
		utils.PrintMessage("Nama Rute tidak boleh kosong", "error")
		goto Step2
	}
Step3:
	fmt.Print("Masukan Harga Rute: ")
	Harga, err := reader.ReadString('\n')
	Harga = strings.TrimSpace(Harga)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step3
	}
	if Harga == "" {
		utils.PrintMessage("Harga Rute tidak boleh kosong", "error")
		goto Step3
	}
	HargaInt, err := strconv.Atoi(Harga)
	if err != nil {
		utils.PrintMessage("Harga Rute harus berupa angka", "error")
		goto Step3
	}
Step4:
	HargaTetapBool := false
	fmt.Print("Masukan Harga Tetap (y/n): ")
	HargaTetap, err := reader.ReadString('\n')
	HargaTetap = strings.TrimSpace(HargaTetap)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step4
	}
	if HargaTetap == "" {
		utils.PrintMessage("Harga Tetap tidak boleh kosong", "error")
		goto Step4
	}
	if HargaTetap != "y" && HargaTetap != "n" {
		utils.PrintMessage("Harga Tetap harus berupa y/n", "error")
		goto Step4
	}
	if HargaTetap == "y" {
		HargaTetapBool = true
	} else {
		HargaTetapBool = false
	}
Step5:
	fmt.Print("Masukan Jumlah Gerbong: ")
	Gerbong, err := reader.ReadString('\n')
	Gerbong = strings.TrimSpace(Gerbong)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step5
	}
	if Gerbong == "" {
		utils.PrintMessage("Jumlah Gerbong tidak boleh kosong", "error")
		goto Step5
	}

	if !utils.IsNumeric(Gerbong) {
		utils.PrintMessage("Jumlah Gerbong harus berupa angka", "error")
		goto Step5
	}

	GerbongInt, err := strconv.Atoi(Gerbong)
	if err != nil {
		utils.PrintMessage("Jumlah Gerbong harus berupa angka", "error")
		goto Step5
	}
	utils.Divider("-")
	for _, kereta := range model.ListKereta {
		fmt.Printf("[%d] %s\n", kereta.Kode, kereta.Nama)
	}
Step6:
	fmt.Print("Masukan Kode Kereta: ")
	KodeKereta, err := reader.ReadString('\n')
	KodeKereta = strings.TrimSpace(KodeKereta)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step6
	}
	if KodeKereta == "" {
		utils.PrintMessage("Kode Kereta tidak boleh kosong", "error")
		goto Step6
	}

	KodeKeretaInt, err := strconv.Atoi(KodeKereta)
	if err != nil {
		utils.PrintMessage("Kode Kereta harus berupa angka", "error")
		goto Step6
	}

	var Kereta model.Kereta
	var Have bool
	Kereta, Have, _ = utils.FindOne(model.ListKereta, model.Kereta{Kode: KodeKeretaInt}, func(a, b model.Kereta) int {
		if a.Kode < b.Kode {
			return -1
		} else if a.Kode > b.Kode {
			return 1
		}
		return 0
	})
	if !Have {
		utils.PrintMessage("Kode Kereta tidak ditemukan", "error")
		goto Step6
	}
	for _, stasiun := range model.ListStasiun {
		fmt.Printf("[%s] %s\n", stasiun.IDStasiun, stasiun.Nama)
	}
Step7:
	fmt.Print("Masukan Kode Stasiun Awal: ")
	KodeStasiunAwal, err := reader.ReadString('\n')
	KodeStasiunAwal = strings.TrimSpace(KodeStasiunAwal)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step7
	}
	if KodeStasiunAwal == "" {
		utils.PrintMessage("Kode Stasiun Awal tidak boleh kosong", "error")
		goto Step7
	}
	var StasiunAwal model.Stasiun
	StasiunAwal, Have, _ = utils.FindOne(model.ListStasiun, model.Stasiun{IDStasiun: KodeStasiunAwal}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})

	if !Have {
		utils.PrintMessage("Kode Stasiun Awal tidak ditemukan", "error")
		goto Step7
	}
Step8:
	fmt.Print("Masukan Kode Stasiun Akhir: ")
	KodeStasiunAkhir, err := reader.ReadString('\n')
	KodeStasiunAkhir = strings.TrimSpace(KodeStasiunAkhir)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step8
	}
	if KodeStasiunAkhir == "" {
		utils.PrintMessage("Kode Stasiun Akhir tidak boleh kosong", "error")
		goto Step8
	}
	var StasiunAkhir model.Stasiun
	StasiunAkhir, Have, _ = utils.FindOne(model.ListStasiun, model.Stasiun{IDStasiun: KodeStasiunAkhir}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})
	if !Have {
		utils.PrintMessage("Kode Stasiun Akhir tidak ditemukan", "error")
		goto Step8
	}
	// add rute on RuteList
	baru := model.Rute{
		Kode:         Kode,
		Nama:         Nama,
		Harga:        HargaInt,
		HargaTetap:   HargaTetapBool,
		Gerbong:      GerbongInt,
		Kereta:       Kereta,
		StasiunAwal:  StasiunAwal,
		StasiunAkhir: StasiunAkhir,
		RuteBerhenti: []model.RuteBerhenti{},
	}
	model.RuteList = append(model.RuteList, baru)
	utils.ClearScreen()
	fmt.Println(utils.ColorText("Rute berhasil ditambahkan.", 30, 42, false))
	RuteController()
}

func EditRute(Rute model.Rute) {

	utils.PrintHead("Edit Data Rute")
	reader := bufio.NewReader(os.Stdin)

Step1:
	fmt.Printf("Masukan Nama Rute (%s): ", Rute.Nama)
	Nama, err := reader.ReadString('\n')
	Nama = strings.TrimSpace(Nama)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step1
	}

	if Nama == "" {
		utils.PrintMessage("Nama Rute tidak boleh kosong", "error")
		goto Step1
	}
Step2:
	fmt.Printf("Masukan Harga Rute (%d): ", Rute.Harga)
	Harga, err := reader.ReadString('\n')
	Harga = strings.TrimSpace(Harga)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step2
	}
	if Harga == "" {
		utils.PrintMessage("Harga Rute tidak boleh kosong", "error")
		goto Step2
	}
	HargaInt, err := strconv.Atoi(Harga)
	if err != nil {
		utils.PrintMessage("Harga Rute harus berupa angka", "error")
		goto Step2
	}
Step3:
	HargaTetapBool := false
	var hargaTetapStr string
	if Rute.HargaTetap {
		hargaTetapStr = "y"
	} else {
		hargaTetapStr = "n"
	}
	fmt.Printf("Masukan Harga Tetap (%s): ", hargaTetapStr)
	HargaTetap, err := reader.ReadString('\n')
	HargaTetap = strings.TrimSpace(HargaTetap)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step3
	}
	if HargaTetap == "" {
		utils.PrintMessage("Harga Tetap tidak boleh kosong", "error")
		goto Step3
	}
	if HargaTetap != "y" && HargaTetap != "n" {
		utils.PrintMessage("Harga Tetap harus berupa y/n", "error")
		goto Step3
	}
	if HargaTetap == "y" {
		HargaTetapBool = true
	} else {
		HargaTetapBool = false
	}
Step4:
	fmt.Printf("Masukan Jumlah Gerbong (%d): ", Rute.Gerbong)
	Gerbong, err := reader.ReadString('\n')
	Gerbong = strings.TrimSpace(Gerbong)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step4
	}
	if Gerbong == "" {
		utils.PrintMessage("Jumlah Gerbong tidak boleh kosong", "error")
		goto Step4
	}
	if !utils.IsNumeric(Gerbong) {
		utils.PrintMessage("Jumlah Gerbong harus berupa angka", "error")
		goto Step4
	}
	GerbongInt, err := strconv.Atoi(Gerbong)
	if err != nil {
		utils.PrintMessage("Jumlah Gerbong harus berupa angka", "error")
		goto Step4
	}
	utils.Divider("-")
	for _, kereta := range model.ListKereta {
		fmt.Printf("[%d] %s\n", kereta.Kode, kereta.Nama)
	}
Step5:
	fmt.Printf("Masukan Kode Kereta (%d): ", Rute.Kereta.Kode)
	KodeKereta, err := reader.ReadString('\n')
	KodeKereta = strings.TrimSpace(KodeKereta)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step5
	}
	if KodeKereta == "" {
		utils.PrintMessage("Kode Kereta tidak boleh kosong", "error")
		goto Step5
	}
	KodeKeretaInt, err := strconv.Atoi(KodeKereta)
	if err != nil {
		utils.PrintMessage("Kode Kereta harus berupa angka", "error")
		goto Step5
	}
	var Kereta model.Kereta
	var Have bool
	Kereta, Have, _ = utils.FindOne(model.ListKereta, model.Kereta{Kode: KodeKeretaInt}, func(a, b model.Kereta) int {
		if a.Kode < b.Kode {
			return -1
		} else if a.Kode > b.Kode {
			return 1
		}
		return 0
	})
	if !Have {
		utils.PrintMessage("Kode Kereta tidak ditemukan", "error")
		goto Step5
	}
	for _, stasiun := range model.ListStasiun {
		fmt.Printf("[%s] %s\n", stasiun.IDStasiun, stasiun.Nama)
	}
Step6:
	fmt.Printf("Masukan Kode Stasiun Awal (%s): ", Rute.StasiunAwal.IDStasiun)
	KodeStasiunAwal, err := reader.ReadString('\n')
	KodeStasiunAwal = strings.TrimSpace(KodeStasiunAwal)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step6
	}
	if KodeStasiunAwal == "" {
		utils.PrintMessage("Kode Stasiun Awal tidak boleh kosong", "error")
		goto Step6
	}
	var StasiunAwal model.Stasiun
	StasiunAwal, Have, _ = utils.FindOne(model.ListStasiun, model.Stasiun{IDStasiun: KodeStasiunAwal}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})
	if !Have {
		utils.PrintMessage("Kode Stasiun Awal tidak ditemukan", "error")
		goto Step6
	}
Step7:
	fmt.Printf("Masukan Kode Stasiun Akhir (%s): ", Rute.StasiunAkhir.IDStasiun)
	KodeStasiunAkhir, err := reader.ReadString('\n')
	KodeStasiunAkhir = strings.TrimSpace(KodeStasiunAkhir)
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan", "error")
		goto Step7
	}
	if KodeStasiunAkhir == "" {
		utils.PrintMessage("Kode Stasiun Akhir tidak boleh kosong", "error")
		goto Step7
	}
	var StasiunAkhir model.Stasiun
	StasiunAkhir, Have, _ = utils.FindOne(model.ListStasiun, model.Stasiun{IDStasiun: KodeStasiunAkhir}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})
	if !Have {
		utils.PrintMessage("Kode Stasiun Akhir tidak ditemukan", "error")
		goto Step7
	}
	// update rute on RuteList
	for i, rute := range model.RuteList {
		if rute.Kode == Rute.Kode {
			model.RuteList[i].Nama = Nama
			model.RuteList[i].Harga = HargaInt
			model.RuteList[i].HargaTetap = HargaTetapBool
			model.RuteList[i].Gerbong = GerbongInt
			model.RuteList[i].Kereta = Kereta
			model.RuteList[i].StasiunAwal = StasiunAwal
			model.RuteList[i].StasiunAkhir = StasiunAkhir
		}
	}
	utils.ClearScreen()
	fmt.Println(utils.ColorText("Rute berhasil diubah.", 30, 42, false))
	RuteController()
}

func DeleteRute(Rute model.Rute) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Apakah anda yakin ingin menghapus rute ini? (y/n): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm == "y" || confirm == "Y" {
		for i, rute := range model.RuteList {
			if rute.Kode == Rute.Kode {
				model.RuteList = append(model.RuteList[:i], model.RuteList[i+1:]...)
				utils.ClearScreen()
				fmt.Println(utils.ColorText("Rute berhasil dihapus.", 30, 42, false))
				RuteController()
			}
		}
	} else {
		utils.ClearScreen()
		utils.PrintMessage("Rute tidak dihapus", "error")
		DetailRute(Rute)
	}
}
