package controller

import (
	"bufio"
	"fmt"
	"laraska/model"
	"os"
	"strconv"
	"strings"
	"time"
)

func MenuRute() {
	var choice string
	var menuList = []string{
		"Tampilkan Rute",
		"Tambah Rute",
		"Edit Rute",
		"Hapus Rute",
		"Kembali ke Menu Awal",
	}

	PrintJudul("Menu Rute")

	for index, menu := range menuList {
		fmt.Printf("[%d] %s\n", index+1, menu)
	}
	Pembatas("-")

	fmt.Print("Pilih menu: ")
	_, err := fmt.Scan(&choice)
	if err != nil || !isNumeric(choice) {
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		MenuRute()
		return
	}

	switch choice {
	case "1":
		ClearScreen()
		TampilkanRute()
	case "2":
		ClearScreen()
		TambahRute()
	case "3":
		ClearScreen()
		EditRute()
	case "4":
		ClearScreen()
		HapusRute()
	case "5":
		ClearScreen()
	default:
		PrintError("Pilihan tidak valid, silakan coba lagi.")
		MenuRute()
	}
}

func tableRute(Title string) {
	var mapped []map[string]string
	for _, dt := range model.ListRute {
		mapped = append(mapped, map[string]string{
			"Kode":          dt.Kode,
			"Harga":         strconv.Itoa(dt.Harga),
			"Kapasitas":     strconv.Itoa(dt.Kapasitas),
			"Gerbong":       strconv.Itoa(dt.Gerbong),
			"Keberangkatan": dt.Keberangkatan.Format("2006-01-02 15:04"),
			"Kedatangan":    dt.Kedatangan.Format("2006-01-02 15:04"),
			"Kereta":        dt.Kereta.Nama,
			"StasiunAwal":   dt.StasiunAwal.Nama,
			"StasiunTujuan": dt.StasiunTujuan.Nama,
		})
	}

	PrintTable(
		[]string{"Kode", "Harga", "Kapasitas", "Gerbong", "Keberangkatan", "Kedatangan", "Kereta", "StasiunAwal", "StasiunTujuan"},
		mapped,
		[]string{"Kode", "Harga", "Kapasitas", "Gerbong", "Keberangkatan", "Kedatangan", "Kereta", "StasiunAwal", "StasiunTujuan"},
		4,
		Title,
	)
}

func TampilkanRute() {

	var searchOn, sortOn, sortType string

	stop := false
	reader := bufio.NewReader(os.Stdin)

	for !stop {
		var menu string
		var mapped []map[string]string
		Data := model.ListRute

		switch {
		case sortOn == "Kode" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Kode < b.Kode })
		case sortOn == "Kode" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Kode > b.Kode })
		case sortOn == "Harga" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Harga < b.Harga })
		case sortOn == "Harga" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Harga > b.Harga })
		case sortOn == "Kapasitas" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Kapasitas < b.Kapasitas })
		case sortOn == "Kapasitas" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Kapasitas > b.Kapasitas })
		case sortOn == "Gerbong" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Gerbong < b.Gerbong })
		case sortOn == "Gerbong" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Gerbong > b.Gerbong })
		case sortOn == "Keberangkatan" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Keberangkatan.Before(b.Keberangkatan) })
		case sortOn == "Keberangkatan" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Keberangkatan.After(b.Keberangkatan) })
		case sortOn == "Kedatangan" && sortType == "asc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Kedatangan.Before(b.Kedatangan) })
		case sortOn == "Kedatangan" && sortType == "desc":
			Data = InsertionSort(Data, func(a, b model.Rute) bool { return a.Kedatangan.After(b.Kedatangan) })
		}

		for _, dt := range Data {
			if searchOn == "" || strings.Contains(strings.ToLower(fmt.Sprintf("%s %s %s %s %s %s",
				dt.Kode,
				dt.Kereta.Nama,
				dt.StasiunAwal.Nama,
				dt.StasiunTujuan.Nama,
				dt.Kedatangan.Format("2006-01-02 15:04"),
				dt.Keberangkatan.Format("2006-01-02 15:04"),
			)), strings.ToLower(searchOn)) {
				mapped = append(mapped, map[string]string{
					"Kode":          dt.Kode,
					"Harga":         strconv.Itoa(dt.Harga),
					"Kapasitas":     strconv.Itoa(dt.Kapasitas),
					"Gerbong":       strconv.Itoa(dt.Gerbong),
					"Keberangkatan": dt.Keberangkatan.Format("2006-01-02 15:04"),
					"Kedatangan":    dt.Kedatangan.Format("2006-01-02 15:04"),
					"Kereta":        dt.Kereta.Nama,
					"StasiunAwal":   dt.StasiunAwal.Nama,
					"StasiunTujuan": dt.StasiunTujuan.Nama,
				})
			}
		}

		Title := "Data Kereta"
		if searchOn != "" {
			Title = fmt.Sprintf("Hasil Pencarian Data Kereta : %s", searchOn)
		}

		PrintTable(
			[]string{"Kode", "Harga", "Kapasitas", "Gerbong", "Keberangkatan", "Kedatangan", "Kereta", "StasiunAwal", "StasiunTujuan"},
			mapped,
			[]string{"Kode", "Harga", "Kapasitas", "Gerbong", "Keberangkatan", "Kedatangan", "Kereta", "StasiunAwal", "StasiunTujuan"},
			4,
			Title,
		)

		fmt.Println(ColorText("[1] Pencarian", 90, 49, false))
		fmt.Println(ColorText("[2] Tampilkan Seluruh Data", 90, 49, false))
		fmt.Println(ColorText("[3] Sorting Data", 90, 49, false))
		fmt.Println(ColorText("[4] Kembali Ke Menu Stasiun", 90, 49, false))

		Pembatas("-")

		fmt.Print("Masukan nomor menu : ")
		_, err := fmt.Scan(&menu)
		if err != nil || !isNumeric(menu) {
			ClearScreen()
			PrintError("Pilihan tidak valid, silakan coba lagi.")
		}

		switch menu {
		case "1":
			fmt.Print("Masukan keyword pencarian : ")
			search, _ := reader.ReadString('\n')
			searchOn = strings.TrimSpace(search)
			ClearScreen()
		case "2":
			searchOn = ""
			sortOn = ""
			sortType = ""
			ClearScreen()
		case "3":
			fmt.Print("Pilih kolom untuk sort : ")
			fmt.Scan(&sortOn)
			fmt.Print("Pilih jenis sort (asc/desc) : ")
			fmt.Scan(&sortType)
			ClearScreen()
		case "4":
			stop = true
			ClearScreen()
			MenuKereta()
		default:
			ClearScreen()
			PrintError("Pilihan tidak valid, silakan coba lagi.")
		}

	}

}

func TambahRute() {
	PrintJudul("Tambah Rute Perjalanan")

	reader := bufio.NewReader(os.Stdin)

StepHarga:
	fmt.Print("Masukkan Harga: ")
	hargaInput, _ := reader.ReadString('\n')
	hargaInput = strings.TrimSpace(hargaInput)
	harga, err := strconv.Atoi(hargaInput)
	if err != nil {
		PrintError("Harga harus berupa angka.")
		goto StepHarga
	}

StepKapasitas:
	fmt.Print("Masukkan Kapasitas: ")
	kapasitasInput, _ := reader.ReadString('\n')
	kapasitasInput = strings.TrimSpace(kapasitasInput)
	kapasitas, err := strconv.Atoi(kapasitasInput)
	if err != nil {
		PrintError("Kapasitas harus berupa angka.")
		goto StepKapasitas
	}

StepGerbong:
	fmt.Print("Masukkan Jumlah Gerbong: ")
	gerbongInput, _ := reader.ReadString('\n')
	gerbongInput = strings.TrimSpace(gerbongInput)
	gerbong, err := strconv.Atoi(gerbongInput)
	if err != nil {
		PrintError("Jumlah gerbong harus berupa angka.")
		goto StepGerbong
	}

StepKeberangkatan:
	fmt.Print("Masukkan Waktu Keberangkatan (YYYY-MM-DD HH:MM): ")
	keberangkatanInput, _ := reader.ReadString('\n')
	keberangkatanInput = strings.TrimSpace(keberangkatanInput)
	keberangkatan, err := time.Parse("2006-01-02 15:04", keberangkatanInput)
	if err != nil {
		PrintError("Format tanggal tidak valid.")
		goto StepKeberangkatan
	}

StepKedatangan:
	fmt.Print("Masukkan Waktu Kedatangan (YYYY-MM-DD HH:MM): ")
	kedatanganInput, _ := reader.ReadString('\n')
	kedatanganInput = strings.TrimSpace(kedatanganInput)
	kedatangan, err := time.Parse("2006-01-02 15:04", kedatanganInput)
	if err != nil {
		PrintError("Format tanggal tidak valid.")
		goto StepKedatangan
	}

	fmt.Println("Pilih Kereta:")
	for i, k := range model.ListKereta {
		fmt.Printf("[%d] %s (%s)\n", i, k.Nama, k.Kelas)
	}
StepKereta:
	fmt.Print("Masukkan nomor kereta: ")
	keretaInput, _ := reader.ReadString('\n')
	keretaInput = strings.TrimSpace(keretaInput)
	idxKereta, err := strconv.Atoi(keretaInput)
	if err != nil || idxKereta < 0 || idxKereta >= len(model.ListKereta) {
		PrintError("Kereta tidak valid.")
		goto StepKereta
	}

	fmt.Println("Pilih Stasiun Awal:")
	for i, s := range model.ListStasiun {
		fmt.Printf("[%d] %s (%s)\n", i, s.Nama, s.Kota)
	}
StepStAwal:
	fmt.Print("Masukkan nomor stasiun awal: ")
	stAwalInput, _ := reader.ReadString('\n')
	stAwalInput = strings.TrimSpace(stAwalInput)
	idxStAwal, err := strconv.Atoi(stAwalInput)
	if err != nil || idxStAwal < 0 || idxStAwal >= len(model.ListStasiun) {
		PrintError("Stasiun awal tidak valid.")
		goto StepStAwal
	}

	fmt.Println("Pilih Stasiun Tujuan:")
	for i, s := range model.ListStasiun {
		fmt.Printf("[%d] %s (%s)\n", i, s.Nama, s.Kota)
	}
StepStTujuan:
	fmt.Print("Masukkan nomor stasiun tujuan: ")
	stTujuanInput, _ := reader.ReadString('\n')
	stTujuanInput = strings.TrimSpace(stTujuanInput)
	idxStTujuan, err := strconv.Atoi(stTujuanInput)
	if err != nil || idxStTujuan < 0 || idxStTujuan >= len(model.ListStasiun) || idxStTujuan == idxStAwal {
		PrintError("Stasiun tujuan tidak valid atau sama dengan stasiun awal.")
		goto StepStTujuan
	}

	data := model.Rute{
		Harga:         harga,
		Kapasitas:     kapasitas,
		Gerbong:       gerbong,
		Keberangkatan: keberangkatan,
		Kedatangan:    kedatangan,
		Kereta:        model.ListKereta[idxKereta],
		StasiunAwal:   model.ListStasiun[idxStAwal],
		StasiunTujuan: model.ListStasiun[idxStTujuan],
	}

	model.ListRute = append(model.ListRute, data)
	fmt.Println(ColorText("Data Rute berhasil ditambahkan.", 30, 42, false))
	MenuRute()
}

func EditRute() {
	var kodeRute string
	tableRute("Edit Rute Perjalanan")
pilihan:
	fmt.Print("Pilih kode rute yang ingin di edit : ")
	_, err := fmt.Scan(&kodeRute)
	if err != nil {
		PrintError("Input tidak valid")
		goto pilihan
	}

	index := FindIndexByKode(model.ListRute, kodeRute)
	if index == -1 {
		PrintError("Kode rute tidak tersedia")
		goto pilihan
	}

	Pembatas("-")
	old := model.ListRute[index]
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Masukkan Harga [%d]: ", old.Harga)
	hargaInput, _ := reader.ReadString('\n')
	hargaInput = strings.TrimSpace(hargaInput)
	harga := old.Harga
	if hargaInput != "" {
		if val, err := strconv.Atoi(hargaInput); err == nil {
			harga = val
		} else {
			PrintError("Harga tidak valid, gunakan nilai sebelumnya.")
		}
	}

	fmt.Printf("Masukkan Kapasitas [%d]: ", old.Kapasitas)
	kapasitasInput, _ := reader.ReadString('\n')
	kapasitasInput = strings.TrimSpace(kapasitasInput)
	kapasitas := old.Kapasitas
	if kapasitasInput != "" {
		if val, err := strconv.Atoi(kapasitasInput); err == nil {
			kapasitas = val
		} else {
			PrintError("Kapasitas tidak valid, gunakan nilai sebelumnya.")
		}
	}

	fmt.Printf("Masukkan Jumlah Gerbong [%d]: ", old.Gerbong)
	gerbongInput, _ := reader.ReadString('\n')
	gerbongInput = strings.TrimSpace(gerbongInput)
	gerbong := old.Gerbong
	if gerbongInput != "" {
		if val, err := strconv.Atoi(gerbongInput); err == nil {
			gerbong = val
		} else {
			PrintError("Gerbong tidak valid, gunakan nilai sebelumnya.")
		}
	}

	fmt.Printf("Masukkan Waktu Keberangkatan [%s]: ", old.Keberangkatan.Format("2006-01-02 15:04"))
	keberangkatanInput, _ := reader.ReadString('\n')
	keberangkatanInput = strings.TrimSpace(keberangkatanInput)
	keberangkatan := old.Keberangkatan
	if keberangkatanInput != "" {
		if val, err := time.Parse("2006-01-02 15:04", keberangkatanInput); err == nil {
			keberangkatan = val
		} else {
			PrintError("Format tanggal tidak valid, gunakan nilai sebelumnya.")
		}
	}

	fmt.Printf("Masukkan Waktu Kedatangan [%s]: ", old.Kedatangan.Format("2006-01-02 15:04"))
	kedatanganInput, _ := reader.ReadString('\n')
	kedatanganInput = strings.TrimSpace(kedatanganInput)
	kedatangan := old.Kedatangan
	if kedatanganInput != "" {
		if val, err := time.Parse("2006-01-02 15:04", kedatanganInput); err == nil {
			kedatangan = val
		} else {
			PrintError("Format tanggal tidak valid, gunakan nilai sebelumnya.")
		}
	}

	fmt.Printf("Masukkan nomor kereta [%s]:\n", old.Kereta.Nama)
	for i, k := range model.ListKereta {
		fmt.Printf("[%d] %s (%s)\n", i, k.Nama, k.Kelas)
	}
	fmt.Print("Pilih kereta (enter untuk skip): ")
	keretaInput, _ := reader.ReadString('\n')
	keretaInput = strings.TrimSpace(keretaInput)
	kereta := old.Kereta
	if keretaInput != "" {
		if idx, err := strconv.Atoi(keretaInput); err == nil && idx >= 0 && idx < len(model.ListKereta) {
			kereta = model.ListKereta[idx]
		} else {
			PrintError("Kereta tidak valid, gunakan nilai sebelumnya.")
		}
	}

	fmt.Printf("Masukkan nomor stasiun awal [%s]:\n", old.StasiunAwal.Nama)
	for i, s := range model.ListStasiun {
		fmt.Printf("[%d] %s (%s)\n", i, s.Nama, s.Kota)
	}
	fmt.Print("Pilih stasiun awal (enter untuk skip): ")
	stAwalInput, _ := reader.ReadString('\n')
	stAwalInput = strings.TrimSpace(stAwalInput)
	stAwal := old.StasiunAwal
	if stAwalInput != "" {
		if idx, err := strconv.Atoi(stAwalInput); err == nil && idx >= 0 && idx < len(model.ListStasiun) {
			stAwal = model.ListStasiun[idx]
		} else {
			PrintError("Stasiun awal tidak valid, gunakan nilai sebelumnya.")
		}
	}

	fmt.Printf("Masukkan nomor stasiun tujuan [%s]:\n", old.StasiunTujuan.Nama)
	for i, s := range model.ListStasiun {
		fmt.Printf("[%d] %s (%s)\n", i, s.Nama, s.Kota)
	}
	fmt.Print("Pilih stasiun tujuan (enter untuk skip): ")
	stTujuanInput, _ := reader.ReadString('\n')
	stTujuanInput = strings.TrimSpace(stTujuanInput)
	stTujuan := old.StasiunTujuan
	if stTujuanInput != "" {
		if idx, err := strconv.Atoi(stTujuanInput); err == nil && idx >= 0 && idx < len(model.ListStasiun) && model.ListStasiun[idx].Nama != stAwal.Nama {
			stTujuan = model.ListStasiun[idx]
		} else {
			PrintError("Stasiun tujuan tidak valid atau sama dengan stasiun awal, gunakan nilai sebelumnya.")
		}
	}

	model.ListRute[index] = model.Rute{
		Kode:          old.Kode,
		Harga:         harga,
		Kapasitas:     kapasitas,
		Gerbong:       gerbong,
		Keberangkatan: keberangkatan,
		Kedatangan:    kedatangan,
		Kereta:        kereta,
		StasiunAwal:   stAwal,
		StasiunTujuan: stTujuan,
	}
	fmt.Println(ColorText("Data Rute berhasil diubah.", 30, 42, false))
	MenuRute()

}

func HapusRute() {
	var kodeRute string
	tableRute("Hapus Rute Perjalanan")
pilihan:
	fmt.Print("Pilih kode rute yang ingin dihapus : ")
	_, err := fmt.Scan(&kodeRute)
	if err != nil {
		PrintError("Input tidak valid")
		goto pilihan
	}

	index := FindIndexByKode(model.ListRute, kodeRute)
	if index == -1 {
		PrintError("Kode rute tidak tersedia")
		goto pilihan
	}

	model.ListRute = append(model.ListRute[:index], model.ListRute[index+1:]...)
	ClearScreen()
	fmt.Println(ColorText("Data Rute berhasil dihapus.", 30, 42, false))
	MenuRute()
}


func FindIndexByKode(data []model.Rute, kode string) int {
	for i, v := range data {
		if v.Kode == kode {
			return i
		}
	}
	return -1
}
