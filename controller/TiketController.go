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

func TambahTiketKereta() {
	var NoAwal, NoTujuan string
	var StasiunAwal, StasiunTujuan model.Stasiun
	// var RuteOn model.Rute
	reader := bufio.NewReader(os.Stdin)

	var mappedStasiun []map[string]string
	for _, dt := range model.ListStasiun {
		mappedStasiun = append(mappedStasiun, map[string]string{
			"IDStasiun": dt.IDStasiun,
			"Nama":      dt.Nama,
			"Kota":      dt.Kota,
		})
	}

	PrintTable(
		[]string{"Kode", "Nama", "Kota"},
		mappedStasiun,
		[]string{"IDStasiun", "Nama", "Kota"},
		4,
		"Data Stasiun",
	)

StepStasiun:
	fmt.Print("Masukan Kode Stasiun Awal : ")
	fmt.Scan(&NoAwal)
	fmt.Print("Masukan Kode Stasiun Tujuan : ")
	fmt.Scan(&NoTujuan)

	var ok1, ok2 bool
	for _, s := range model.ListStasiun {
		if s.IDStasiun == NoAwal {
			StasiunAwal, ok1 = s, true
		} else if s.IDStasiun == NoTujuan {
			StasiunTujuan, ok2 = s, true
		}
		if ok1 && ok2 {
			break
		}
	}
	if !ok1 || !ok2 {
		PrintError("Kode Stasiun tidak valid")
		goto StepStasiun
	}

StepTanggal:
	fmt.Print("Masukkan Tanggal Keberangkatan (YYYY-MM-DD): ")
	dateInput, _ := reader.ReadString('\n')
	dateInput = strings.TrimSpace(dateInput)
	parsedDate, err := time.Parse("2006-01-02", dateInput)
	if err != nil {
		PrintError("Format tidak valid. Gunakan YYYY-MM-DD.")
		goto StepTanggal
	}

	var mapped []map[string]string
	for _, dt := range model.ListRute {
		if dt.StasiunAwal.IDStasiun == StasiunAwal.IDStasiun &&
			dt.StasiunTujuan.IDStasiun == StasiunTujuan.IDStasiun &&
			dt.Keberangkatan.Format("2006-01-02") == parsedDate.Format("2006-01-02") {
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

	ClearScreen()

	PrintTable(
		[]string{"Kode", "Harga", "Kapasitas", "Gerbong", "Keberangkatan", "Kedatangan", "Kereta", "StasiunAwal", "StasiunTujuan"},
		mapped,
		[]string{"Kode", "Harga", "Kapasitas", "Gerbong", "Keberangkatan", "Kedatangan", "Kereta", "StasiunAwal", "StasiunTujuan"},
		4,
		"Pilih Rute Yang Tersedia",
	)

PilihRute:
	var pilihRoute string
	fmt.Print("Pilih Kode Rute : ")
	fmt.Scan(&pilihRoute)

	hasRoute := false
	for _, dt := range mapped {
		if dt["Kode"] == pilihRoute {
			hasRoute = true
		}
	}

	if !hasRoute {
		PrintError("Kode Rute tidak valid")
		goto PilihRute
	}

	fmt.Printf("Masukan Stasiun Awal : %s \n", StasiunAwal.Nama)
	fmt.Printf("Masukan Stasiun Tujuan : %s \n", StasiunTujuan.Nama)
	fmt.Printf("Masukkan Tanggal Keberangkatan (YYYY-MM-DD): %s \n", parsedDate.Format("2006-01-02"))


}
