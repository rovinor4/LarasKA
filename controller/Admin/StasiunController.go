package admin

import (
	"bufio"
	"fmt"
	"laraska/model"
	"laraska/utils"
	"os"
	"strings"
)

func StasiunController() {

	render := bufio.NewReader(os.Stdin)

	utils.PrintHead("Menu Stasiun")
	fmt.Println("[1] Daftar Stasiun")
	fmt.Println("[2] Tambah Stasiun")
	fmt.Println("[3] Edit Stasiun")
	fmt.Println("[4] Hapus Stasiun")
	fmt.Println("[5] Kembali Ke Menu Utama")

	utils.Divider("-")
	fmt.Print("Pilih menu: ")

Step1:
	menu, err := render.ReadString('\n')
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan saat membaca input", "error")
		return
	}
	menu = strings.TrimSpace(menu)
	if menu == "" {
		utils.PrintMessage("Input tidak boleh kosong", "error")
		goto Step1
	}

	switch menu {
	case "1":
		showStasiun()
	case "2":
		addStasiun()
	case "3":
		updateStasiun()
	case "4":
		deleteStasiun()
	case "5":
		MenuAwalAdmin()
	default:
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
		goto Step1
	}
}

func showStasiun() {
	utils.ClearScreen()
	render := bufio.NewReader(os.Stdin)
	var Search, SortColom string
	var SortAsc bool

Step1:
	var mapped []map[string]string
	Data := model.ListStasiun

	switch {
	case SortColom == "ID" && SortAsc:
		Data = utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool { return a.IDStasiun < b.IDStasiun })
	case SortColom == "ID" && !SortAsc:
		Data = utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool { return a.IDStasiun > b.IDStasiun })
	case SortColom == "Nama" && SortAsc:
		Data = utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool { return a.Nama < b.Nama })
	case SortColom == "Nama" && !SortAsc:
		Data = utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool { return a.Nama > b.Nama })
	case SortColom == "Kota" && SortAsc:
		Data = utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool { return a.Kota < b.Kota })
	case SortColom == "Kota" && !SortAsc:
		Data = utils.InsertionSort(model.ListStasiun, func(a, b model.Stasiun) bool { return a.Kota > b.Kota })
	}

	for _, dt := range Data {
		//searching
		if Search == "" ||
			strings.Contains(strings.ToLower(dt.IDStasiun), strings.ToLower(Search)) ||
			strings.Contains(strings.ToLower(dt.Nama), strings.ToLower(Search)) ||
			strings.Contains(strings.ToLower(dt.Kota), strings.ToLower(Search)) {
			mapped = append(mapped, map[string]string{
				"IDStasiun": dt.IDStasiun,
				"Nama":      dt.Nama,
				"Kota":      dt.Kota,
			})
		}
	}

	utils.PrintTable(
		[]string{"ID", "Nama", "Kota"},
		mapped,
		[]string{"IDStasiun", "Nama", "Kota"},
		4,
		"Daftar Stasiun",
	)

	fmt.Println(utils.ColorText("[1] Pencarian", 90, 49, false))
	fmt.Println(utils.ColorText("[2] Sortir", 90, 49, false))
	fmt.Println(utils.ColorText("[3] Tampilkan Seluruh Data", 90, 49, false))
	fmt.Println(utils.ColorText("[4] Kembali Ke Menu Kereta", 90, 49, false))
	utils.Divider("-")
Step2:
	fmt.Print("Pilih menu: ")
	menu, err := render.ReadString('\n')
	menu = strings.TrimSpace(menu)
	if err != nil || menu == "" {
		utils.PrintMessage("Input tidak boleh kosong", "error")
		goto Step2
	}

Step3:
	switch menu {
	case "1":
		fmt.Print("Masukkan kata kunci pencarian: ")
		Search, err = render.ReadString('\n')
		Search = strings.TrimSpace(Search)
		if err != nil || Search == "" {
			utils.PrintMessage("Input tidak boleh kosong", "error")
			goto Step3
		}
		utils.ClearScreen()
		goto Step1
	case "2":

		fmt.Print("Masukkan kolom yang ingin disortir (ID/Nama/Kota): ")
		SortColom, err = render.ReadString('\n')
		SortColom = strings.TrimSpace(SortColom)
		if err != nil || SortColom == "" {
			utils.PrintMessage("Input tidak boleh kosong", "error")
			goto Step3
		}
		if SortColom != "ID" && SortColom != "Nama" && SortColom != "Kota" {
			utils.PrintMessage("Kolom tidak valid, silakan coba lagi.", "error")
			goto Step3
		}
		fmt.Print("Apakah ingin diurutkan secara ascending? (y/n): ")
		ascInput, err := render.ReadString('\n')
		ascInput = strings.TrimSpace(ascInput)
		if err != nil || (ascInput != "y" && ascInput != "n") {
			utils.PrintMessage("Input tidak valid, silakan coba lagi.", "error")
			goto Step3
		}

		if ascInput == "y" {
			SortAsc = true
		} else {
			SortAsc = false
		}
		utils.ClearScreen()
		goto Step1
	case "3":
		Search = ""
		SortColom = ""
		SortAsc = false
		utils.ClearScreen()
		goto Step1
	case "4":
		utils.ClearScreen()
		StasiunController()
	default:
		utils.PrintMessage("Pilihan tidak valid, silakan coba lagi.", "error")
		goto Step2
	}
}

func form(Stasiun model.Stasiun, ForUpdate bool) model.Stasiun {
	render := bufio.NewReader(os.Stdin)

	format := "Masukkan %s %s: "
	if ForUpdate {
		format = "Masukan %s (%s): "
	}

FormIdStasiun:
	fmt.Printf(format, "ID Stasiun", Stasiun.IDStasiun)
	IDStasiun, err := render.ReadString('\n')
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan saat membaca input", "error")
		goto FormIdStasiun
	}

	IDStasiun = strings.TrimSpace(IDStasiun)
	if IDStasiun == "" && !ForUpdate {
		utils.PrintMessage("ID Stasiun tidak boleh kosong", "error")
		goto FormIdStasiun
	}

	// urutkan ID Stasiun
	urut := utils.SelectionSort(model.ListStasiun, func(a, b model.Stasiun) bool {
		return a.IDStasiun < b.IDStasiun
	})

	// check binary search
	_, have := utils.BinaryFindOne(urut, model.Stasiun{IDStasiun: IDStasiun}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})

	if have && !ForUpdate {
		utils.PrintMessage("ID Stasiun sudah ada, silakan coba lagi.", "error")
		goto FormIdStasiun
	}

	if !ForUpdate || (ForUpdate && IDStasiun != "") {
		Stasiun.IDStasiun = IDStasiun
	}

FormNamaStasiun:
	fmt.Printf(format, "Nama Stasiun", Stasiun.Nama)
	NamaStasiun, err := render.ReadString('\n')
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan saat membaca input", "error")
		goto FormNamaStasiun
	}
	NamaStasiun = strings.TrimSpace(NamaStasiun)
	if NamaStasiun == "" && !ForUpdate {
		utils.PrintMessage("Nama Stasiun tidak boleh kosong", "error")
		goto FormNamaStasiun
	}
	Stasiun.Nama = NamaStasiun
FormKotaStasiun:
	fmt.Printf(format, "Kota", Stasiun.Kota)
	KotaStasiun, err := render.ReadString('\n')
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan saat membaca input", "error")
		goto FormKotaStasiun
	}
	KotaStasiun = strings.TrimSpace(KotaStasiun)
	if KotaStasiun == "" && !ForUpdate {
		utils.PrintMessage("Kota tidak boleh kosong", "error")
		goto FormKotaStasiun
	}
	Stasiun.Kota = KotaStasiun
	return Stasiun

}

func addStasiun() {
	utils.PrintHead("Tambah Stasiun")
	var stasiun model.Stasiun
	stasiun = form(stasiun, false)

	// tambahkan stasiun ke list
	model.ListStasiun = append(model.ListStasiun, stasiun)
	utils.ClearScreen()
	utils.PrintMessage("Stasiun berhasil ditambahkan.", "success")
	StasiunController()
}

func updateStasiun() {
	reader := bufio.NewReader(os.Stdin)
	utils.PrintHead("Edit Stasiun")
Step1:
	fmt.Print("Masukkan id stasiun yang ingin diedit: ")
	IdStasiun, err := reader.ReadString('\n')
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan saat membaca input", "error")
		goto Step1
	}
	if IdStasiun == "" || !utils.IsNumeric(IdStasiun) {
		utils.PrintMessage("ID Stasiun tidak valid", "error")
		goto Step1
	}

	IdStasiun = strings.TrimSpace(IdStasiun)
	DataFinde, have := utils.FindOne(model.ListStasiun, model.Stasiun{IDStasiun: IdStasiun}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})

	if !have {
		utils.PrintMessage("ID Stasiun tidak ditemukan", "error")
		goto Step1
	}

	Form := form(DataFinde, true)

	for i, st := range model.ListStasiun {
		if st.IDStasiun == Form.IDStasiun {
			model.ListStasiun[i] = Form
			utils.ClearScreen()
			utils.PrintMessage("Stasiun berhasil diupdate.", "success")
			StasiunController()
			break
		}
	}
}

func deleteStasiun() {
	reader := bufio.NewReader(os.Stdin)
	utils.PrintHead("Hapus Stasiun")
Step1:
	fmt.Println("Masukkan id stasiun yang ingin dihapus: ")
	IdStasiun, err := reader.ReadString('\n')
	if err != nil {
		utils.PrintMessage("Terjadi kesalahan saat membaca input", "error")
		goto Step1
	}
	if IdStasiun == "" || !utils.IsNumeric(IdStasiun) {
		utils.PrintMessage("ID Stasiun tidak valid", "error")
		goto Step1
	}
	IdStasiun = strings.TrimSpace(IdStasiun)
	DataFinde, have := utils.FindOne(model.ListStasiun, model.Stasiun{IDStasiun: IdStasiun}, func(a, b model.Stasiun) int {
		if a.IDStasiun < b.IDStasiun {
			return -1
		} else if a.IDStasiun > b.IDStasiun {
			return 1
		}
		return 0
	})
	if !have {
		utils.PrintMessage("ID Stasiun tidak ditemukan", "error")
		goto Step1
	}

	for i, st := range model.ListStasiun {
		if st.IDStasiun == DataFinde.IDStasiun {
			model.ListStasiun = append(model.ListStasiun[:i], model.ListStasiun[i+1:]...)
			utils.ClearScreen()
			utils.PrintMessage("Stasiun berhasil dihapus.", "success")
			StasiunController()
			break
		}
	}

}
