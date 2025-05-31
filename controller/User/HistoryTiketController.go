package user

import (
	"bufio"
	"fmt"
	"laraska/model"
	"laraska/utils"
	"os"
	"strconv"
)

func ShowHistoryTiket() {
	reader := bufio.NewReader(os.Stdin)
	var mappedTiket []map[string]string

	ascendingTiketByTimeCreated := utils.SelectionSort(model.ListTiket, func(a, b model.Tiket) bool {
		return a.CreatedAt.Before(b.CreatedAt)
	})

	if len(model.ListTiket) == 0 {
		utils.ClearScreen()
		utils.PrintMessage("Belum ada tiket yang dipesan!", "error")
		MenuAwalUser()
	}

	for i, tiket := range ascendingTiketByTimeCreated {

		mappedTiket = append(mappedTiket, map[string]string{
			"no":       strconv.Itoa(i + 1),
			"kode":     tiket.Kode,
			"atasNama": tiket.User.NamaLengkap,
			"rute":     fmt.Sprintf("%s - %s", tiket.Rute.StasiunAwal.Nama, tiket.Rute.StasiunAkhir.Nama),
			"jadwal":   tiket.Rute.RuteBerhenti[0].Berangkat.Format("02/01/2006 15:04"),
		})
	}

	utils.PrintTable(
		[]string{"No.", "Kode", "Atas Nama", "Dari/Ke", "Jadwal"},
		mappedTiket,
		[]string{"no", "kode", "atasNama", "rute", "jadwal"},
		2,
		"History Pemesanan Tiket",
	)

	fmt.Println("[1] Detail tiket")
	fmt.Println("[2] Kembali")

	input := utils.Input("Pilih Menu: ", func(value string) (bool, string) {
		if value == "" {
			return false, ""
		}

		if !utils.IsIn(value, []string{"1", "2"}) {
			return false, "Pilihan tidak tersedia"
		}

		return true, ""

	})

	switch input {
	case "1":
		detailTiket(*reader, ascendingTiketByTimeCreated)
	case "2":
		utils.ClearScreen()
		MenuAwalUser()
	default:
		utils.ClearScreen()
		if input != "" {
			s := fmt.Sprintf("Tidak ada pilihan \"%s\"", input)
			utils.PrintMessage(s, "error")
		}
		ShowHistoryTiket()
	}
}

// TODO: Tiket layout
func detailTiket(reader bufio.Reader, ascendingTiketByTimeCreated []model.Tiket) {
	inputNomor := utils.Input("Pilih nomor tiket: ", func(value string) (bool, string) {
		if value == "" {
			return false, ""
		}

		if !utils.IsIn(value, []string{"1", "2"}) {
			return false, "Pilihan tidak tersedia"
		}

		return true, ""

	})

	nomor, err := strconv.Atoi(inputNomor)
	if err != nil {
		errFmt := fmt.Sprintf("Invalid input \"%s\"", inputNomor)
		utils.PrintMessage(errFmt, "error")
		detailTiket(reader, ascendingTiketByTimeCreated)
	} else if nomor < 1 || nomor > len(ascendingTiketByTimeCreated) {
		errFmt := fmt.Sprintf("Error: Out of range \"%s\"", strconv.Itoa(nomor))
		utils.PrintMessage(errFmt, "error")
		detailTiket(reader, ascendingTiketByTimeCreated)
	}

	selectedTiket := ascendingTiketByTimeCreated[nomor-1]
	ruteBerhenti := selectedTiket.Rute.RuteBerhenti

	bannerFmt := fmt.Sprintf("%s - %s | %s", selectedTiket.Rute.StasiunAwal.Kota, selectedTiket.Rute.StasiunAkhir.Kota, selectedTiket.Rute.Nama)
	utils.PrintHead(bannerFmt)
	utils.PrintBoxWithText(60, []string{
		"\n",
		fmt.Sprintf("%s - %s", selectedTiket.Rute.StasiunAwal.Nama, selectedTiket.Rute.StasiunAkhir.Nama),
		fmt.Sprintf("Atas Nama: %s", selectedTiket.User.NamaLengkap),
		// fmt.Sprintf("Nama Penumpang: %s", selectedTiket.Penumpang[0].Nama),
		fmt.Sprintf("Jadwal Keberangkatan : %s", selectedTiket.Rute.RuteBerhenti[0].Berangkat.Format("2 January 2006 15:04")),
		fmt.Sprintf("Kode Tiket: %s", selectedTiket.Kode),
		fmt.Sprintf("Kereta: %s (%s)", selectedTiket.Rute.Kereta.Nama, selectedTiket.Rute.Kereta.Kelas),
		fmt.Sprintf("Gerbong: %d", selectedTiket.Gerbong),
		fmt.Sprintf("Tempat Duduk: %s", selectedTiket.TempatDuduk),
		fmt.Sprintf("Perkiraan Tiba: %s", ruteBerhenti[len(ruteBerhenti)-1].Tiba.Format("2 January 2006 15:04")),
		fmt.Sprintf("Dibuat Pada: %s", selectedTiket.CreatedAt.Format("2 January 2006")),
		"\n",
	})

	fmt.Print("Press ENTER to continue")
	reader.ReadByte()
	fmt.Println()
	utils.ClearScreen()
	ShowHistoryTiket()
}
